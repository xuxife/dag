package task

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/xuxife/dag"
	"go.uber.org/multierr"
)

// New creates a new Plan,
//
// Example:
//
//	a := task.Func(func(ctx context.Context) error {
//		// task A
//		return nil
//	})
//	b := task.Func(func(ctx context.Context) error {
//		// task B
//		return nil
//	})
//	plan, err := task.New(
//	   	task.Add(a).Then(b),
//	)
//	// handle err
//	err = plan.Start(context.Background())
//	// handle err
//	err = plan.Wait()
//	// handle err
func New(ps ...*Plan) (*Plan, error) {
	p := newPlan()
	dags := []*dag.DAG{}
	for _, pp := range ps {
		pp.mustNotStarted()
		dags = append(dags, &pp.DAG)
	}
	p.DAG.Merge(dags...)
	return p, p.Check()
}

func newPlan() *Plan {
	d, _ := dag.New()
	return &Plan{
		DAG: *d,
	}
}

// Add creates a new Plan with the given tasks,
// tasks are independent.
//
//	 /-> A
//	-|-> B
//	 \-> C
func Add(ts ...Task) *Plan {
	return newPlan().Add(ts...)
}

// Pipelines creates a new Plan with the given tasks,
// the tasks are depended on the previous one.
//
//	--> A --> B --> C
func Pipelines(ts ...Task) *Plan {
	return newPlan().Pipelines(ts...)
}

type Plan struct {
	dag.DAG
	attention []Task

	runningTask atomic.Int32 // the number of running tasks

	mainDone chan struct{} // whether main loop is finished
	taskDone sync.Map      // map[Task]chan struct{}	signal for task finished
	status   sync.Map      // map[Task]TaskResult		result/status of task
	input    sync.Map      // map[Task]*SafeItems[Task]	input for task from dependencies

	cancel     func()
	cancelOnce sync.Once

	taskSucceeded chan Task // signal for proceeding main loop
}

// Add the given tasks to current Plan,
// tasks are independent.
// Add ignores previous tasks.
//
//	-> A
//	-> B
//	-> C
func (p *Plan) Add(ts ...Task) *Plan {
	p.mustNotStarted()
	p.DAG.AddVertex(t2v(ts)...)
	p.attention = ts
	return p
}

// Then adds follow-up tasks to current Plan,
// the tasks are depended on the previous one.
//
//	   /-> A
//	. -|-> B   // . is previous tasks
//	   \-> C
func (p *Plan) Then(ts ...Task) *Plan {
	p.mustNotStarted()
	p.DAG.AddEdge(dag.From(t2v(p.attention)...).To(t2v(ts)...)...)
	p.attention = ts
	return p
}

// Pipelines add the given tasks to current Plan,
// the tasks are depended on the previous one.
//
//	. --> A --> B --> C   // . is previous tasks
func (p *Plan) Pipelines(ts ...Task) *Plan {
	p.mustNotStarted()
	for i := 0; i < len(ts)-1; i++ {
		p.Add(ts[i]).Then(ts[i+1])
	}
	return p
}

// DependsOn adds the given tasks to current Plan,
// the tasks become the dependencies of the previous tasks.
//
//	A --\
//	B --|-> .   // . is previous tasks
//	C --/
func (p *Plan) DependsOn(ts ...Task) *Plan {
	p.mustNotStarted()
	p.DAG.AddEdge(dag.From(t2v(ts)...).To(t2v(p.attention)...)...)
	p.attention = ts
	return p
}

// Check whether the Plan DAG having circle or other issues.
func (p *Plan) Check() error {
	if _, err := p.DAG.Sort(); err != nil {
		return err
	}
	return nil
}

// Start kicks off the Plan, tasks are running in topological order.
func (p *Plan) Start(ctx context.Context) error {
	p.mustNotStarted()

	d := p.DAG.Clone()
	if _, err := d.Sort(); err != nil {
		return err
	}

	p.mainDone = make(chan struct{})
	p.cancelOnce = sync.Once{}
	ctx, p.cancel = context.WithCancel(ctx)
	p.taskDone, p.status, p.input = sync.Map{}, sync.Map{}, sync.Map{}
	p.taskSucceeded = make(chan Task, len(d.GetVertices()))

	for _, t := range d.GetVertices() {
		tt := t.(Task)
		p.status.Store(tt, TaskResult{Task: tt, Status: StatusPending})
	}

	// start tasks without dependencies
	startTasksNoDep := func() bool {
		batch := d.GetVerticesNoIn()
		if len(batch) < 1 {
			return true // exit condition, no more tasks
		}
		for _, t := range batch {
			p.startTask(ctx, d, t.(Task))
		}
		return false
	}

	startTasksNoDep()
	go func() {
		defer close(p.mainDone)
		defer p.Cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case task := <-p.taskSucceeded:
				// remove the succeeded task from the DAG
				d.DeleteVertex(task)
				if startTasksNoDep() {
					return
				}
			}
		}
	}()

	return nil
}

func (p *Plan) finishTask(d *dag.DAG, from Task) {
	for _, t := range d.GetVerticesFrom(from) {
		// enqueue the output from current task to
		// tasks depending on current task
		to := t.(Task)
		toInput, _ := p.input.LoadOrStore(to, NewSafeItem[Task]())
		toInput.(*SafeItems[Task]).Append(from)
	}
}

func (p *Plan) startTask(ctx context.Context, d *dag.DAG, t Task) {
	if p.Status(t) != StatusPending {
		return
	}
	p.runningTask.Add(1)
	p.taskDone.Store(t, make(chan struct{}))
	p.status.Store(t, TaskResult{Task: t, Status: StatusRunning})
	go func(ctx context.Context, task Task) {
		var result TaskResult
		defer func() {
			p.runningTask.Add(-1)
			p.status.Store(task, result)
			if result.Err == nil {
				p.finishTask(d, task)
				p.taskSucceeded <- task
			} else {
				p.Cancel()
			}
			// send singal to inform task finished
			taskDone, _ := p.taskDone.Load(task)
			close(taskDone.(chan struct{}))
		}()
		p.loadInput(task)
		// kick off the task
		if err := task.Run(ctx); err != nil {
			result = TaskResult{Task: task, Err: err, Status: StatusFailed}
			return
		}
		result = TaskResult{Task: task, Err: nil, Status: StatusSuccess}
	}(ctx, t)
}

// load input for the task from finished previous tasks
func (p *Plan) loadInput(t Task) {
	// dependency tasks must exit
	for _, from := range p.DAG.GetVerticesTo(t) {
		dep := from.(Task)
		depFinished, _ := p.taskDone.Load(dep)
		<-depFinished.(chan struct{})
	}
	tInput, ok := p.input.LoadAndDelete(t)
	if !ok {
		return
	}
	t.Input(tInput.(*SafeItems[Task]).Get()...)
}

func (p *Plan) IsRunning() bool {
	if p.runningTask.Load() != 0 {
		return true
	}
	if p.mainDone == nil {
		return false
	}
	select {
	case <-p.mainDone: // closed
		return true
	default:
		return false
	}
}

func (p *Plan) mustNotStarted() {
	if p.IsRunning() {
		panic("plan has already started")
	}
}

// Status gets the status of the given task.
func (p *Plan) Status(t Task) TaskStatus {
	return p.Result(t).Status
}

// Statuses gets the statuses of all tasks.
func (p *Plan) Statuses() map[Task]TaskStatus {
	m := map[Task]TaskStatus{}
	for _, t := range p.DAG.GetVertices() {
		m[t.(Task)] = p.Status(t.(Task))
	}
	return m
}

// Cancel cancels the Plan running tasks.
func (p *Plan) Cancel() {
	p.cancelOnce.Do(func() {
		p.cancel()
		p.cancel = nil
		for _, t := range p.DAG.GetVertices() {
			p.status.CompareAndSwap(t, StatusPending, StatusCancelled)
		}
	})
}

// Wait waits for all tasks to finish.
func (p *Plan) Wait() error {
	var err error
	<-p.mainDone
	p.taskDone.Range(func(task, done any) bool {
		<-done.(chan struct{})
		return true
	})
	p.status.Range(func(stask, result any) bool {
		if r := result.(TaskResult); r.Status != StatusSuccess {
			err = multierr.Append(err, r)
		}
		return true
	})
	return err
}

// Result gets the result of the given task.
func (p *Plan) Result(t Task) TaskResult {
	v, ok := p.status.Load(t)
	if !ok {
		return TaskResult{
			Task:   t,
			Status: StatusUnknown,
		}
	}
	return v.(TaskResult)
}

// Results gets the results of all tasks.
func (p *Plan) Results() map[Task]TaskResult {
	results := map[Task]TaskResult{}
	p.status.Range(func(task, result any) bool {
		results[task.(Task)] = result.(TaskResult)
		return true
	})
	return results
}

// Run starts all tasks and wait for them to finish.
func (p *Plan) Run(ctx context.Context) error {
	if err := p.Start(ctx); err != nil {
		return err
	}
	return p.Wait()
}

// task to vertex
func t2v(ts []Task) []dag.Vertex {
	vs := []dag.Vertex{}
	for _, t := range ts {
		vs = append(vs, t)
	}
	return vs
}
