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
//		a := task.Func(func(ctx context.Context) error {
//			// task A
//			return nil
//		})
//		b := task.Func(func(ctx context.Context) error {
//			// task B
//			return nil
//		})
//		p, err := task.New(
//		   	task.Add(a).Then(b),
//		)
//		// handle err
//		taskChan, err := p.Start(context.Background())
//		// handle err
//	 	for taskResult := range taskChan {
//	 		// handle taskResult
//	 	}
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

	started     atomic.Int32 // the number of running tasks
	cancel      func()
	cancelOnce  sync.Once
	status      sync.Map // map[Task]TaskStatus
	input       sync.Map // map[Task][]Task
	result      sync.Map // map[Task]TaskResult
	oneTaskDone chan struct{}
	finished    chan struct{}
}

// Add the given tasks to current Plan,
// tasks are independent.
// Add ignores previous tasks.
//
//	 /-> A
//	-|-> B
//	 \-> C
func (p *Plan) Add(ts ...Task) *Plan {
	p.mustNotStarted()
	p.DAG.AddVertex(t2v(ts)...)
	p.attention = ts
	return p
}

// task to vertex
func t2v(ts []Task) []dag.Vertex {
	vs := []dag.Vertex{}
	for _, t := range ts {
		vs = append(vs, t)
	}
	return vs
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

	d := p.DAG.Duplicate()
	if _, err := d.Sort(); err != nil {
		return err
	}

	for _, t := range d.GetVertices() {
		p.setStatus(t.(Task), StatusPending)
	}

	p.oneTaskDone = make(chan struct{}, len(d.GetVertices()))
	p.finished = make(chan struct{})

	p.cancelOnce = sync.Once{}
	ctx, p.cancel = context.WithCancel(ctx)

	startTaskRank0 := func() bool {
		batch, err := d.Rank(0)
		if err != nil {
			panic(err) // should never happen
		}
		if len(batch) < 1 {
			return true // exit condition
		}
		for _, t := range batch {
			p.startTask(ctx, d, t.(Task))
		}
		return false
	}

	startTaskRank0()
	go func() {
		defer close(p.finished)
		defer p.Cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case <-p.oneTaskDone:
				noMoreTasks := startTaskRank0()
				if noMoreTasks {
					return
				}
			}
		}
	}()

	return nil
}

func (p *Plan) finishTask(d *dag.DAG, t Task) {
	for _, e := range d.GetEdgesFrom(t) {
		// enqueue the output from current task to
		// tasks depending on current task
		tt := e.To().(Task)
		actual, _ := p.input.LoadOrStore(tt, []Task{})
		ts := actual.([]Task)
		ts = append(ts, t)
		p.input.Store(tt, ts)
	}
	d.DeleteVertex(t)
}

func (p *Plan) startTask(ctx context.Context, d *dag.DAG, t Task) {
	if p.Status(t) != StatusPending {
		return
	}
	p.started.Add(1)
	go func(ctx context.Context, task Task) {
		var result TaskResult
		defer func() {
			p.started.Add(-1)
			p.result.Store(task, result)
			if result.Err == nil {
				p.finishTask(d, task)
				p.oneTaskDone <- struct{}{}
			} else {
				p.Cancel()
			}
		}()
		// load input for the task from finished previous tasks
		v, _ := p.input.LoadOrStore(t, []Task{})
		t.Input(v.([]Task)...)
		p.input.Delete(t)
		// kick off the task
		if err := task.Run(ctx); err != nil {
			p.setStatus(task, StatusFailed)
			result = TaskResult{Task: task, Err: err, Status: StatusFailed}
			return
		}
		p.setStatus(task, StatusSuccess)
		result = TaskResult{Task: task, Err: nil, Status: StatusSuccess}
	}(ctx, t)
}

func (p *Plan) IsRunning() bool {
	return p.started.Load() != 0
}

func (p *Plan) mustNotStarted() {
	if p.IsRunning() {
		panic("plan has already started")
	}
}

func (p *Plan) setStatus(t Task, s TaskStatus) {
	p.status.Store(t, s)
}

// Status gets the status of the given task.
func (p *Plan) Status(t Task) TaskStatus {
	s, ok := p.status.Load(t)
	if !ok {
		return StatusUnknown
	}
	return s.(TaskStatus)
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

func (p *Plan) Wait() error {
	var err error
	<-p.finished
	p.result.Range(func(key, value any) bool {
		if r := value.(TaskResult); r.Err != nil {
			err = multierr.Append(err, r)
		}
		return true
	})
	return err
}

func (p *Plan) Result(t Task) TaskResult {
	v, ok := p.result.Load(t)
	if !ok {
		return TaskResult{
			Task:   t,
			Status: StatusUnknown,
		}
	}
	return v.(TaskResult)
}
