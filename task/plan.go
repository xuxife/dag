package task

import (
	"context"
	"fmt"
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
		dags = append(dags, &pp.d)
	}
	p.d.Merge(dags...)
	return p, p.Check()
}

func newPlan() *Plan {
	d, _ := dag.New()
	return &Plan{
		d: *d,
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
	d         dag.DAG
	attention []Task

	started atomic.Int32
	cancel  func()
	status  sync.Map
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
	p.d.AddVertex(t2v(ts)...)
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
	p.d.AddEdge(dag.From(t2v(p.attention)...).To(t2v(ts)...)...)
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
	p.d.AddEdge(dag.From(t2v(ts)...).To(t2v(p.attention)...)...)
	p.attention = ts
	return p
}

// Check whether the Plan DAG having circle or other issues.
func (p *Plan) Check() error {
	var err error
	if _, err = p.d.Sort(); err != nil {
		return err
	}
	for _, e := range p.d.GetEdges() {
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = multierr.Append(err, &TaskInOutMismatchError{
						OutputTask: e.From().(Task),
						InputTask:  e.To().(Task),
						Err:        r,
					})
				}
			}()
			e.To().(Task).Input(e.From().(Task).Output()...)
		}()
	}
	return err
}

type TaskInOutMismatchError struct {
	OutputTask Task
	InputTask  Task
	Err        any
}

func (e *TaskInOutMismatchError) Error() string {
	return fmt.Sprintf("InOutMismatch %s -> %s: %v", e.OutputTask, e.InputTask, e.Err)
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
	for _, t := range p.d.GetVertices() {
		m[t.(Task)] = p.Status(t.(Task))
	}
	return m
}

// Cancel cancels the Plan running tasks.
func (p *Plan) Cancel() {
	if p.cancel == nil {
		return
	}
	p.cancel()
	p.cancel = nil
	for _, t := range p.d.GetVertices() {
		if s := p.Status(t.(Task)); s == StatusPending || s == StatusRunning {
			p.setStatus(t.(Task), StatusCancelled)
		}
	}
}

// Start kicks off the Plan, tasks are running in topological order.
func (p *Plan) Start(ctx context.Context) (<-chan TaskResult, error) {
	p.mustNotStarted()

	if _, err := p.d.Sort(); err != nil {
		return nil, err
	}

	resultChan := make(chan TaskResult, len(p.d.GetVertices()))
	internalChan := make(chan TaskResult)

	d := p.d.Duplicate()
	for _, t := range d.GetVertices() {
		p.setStatus(t.(Task), StatusPending)
	}

	ctx, p.cancel = context.WithCancel(ctx)

	go func(d *dag.DAG) {
		defer close(internalChan)
		defer close(resultChan)
		for {
			batch, err := d.Rank(0)
			if err != nil {
				panic(err) // should never happen
			}
			if len(batch) < 1 {
				break
			}
			for _, t := range batch {
				p.startTask(ctx, t.(Task), internalChan)
			}
			result := <-internalChan
			resultChan <- result
			if result.Err != nil {
				p.Cancel()
				for p.IsRunning() {
					resultChan <- <-internalChan
				}
				break
			} else {
				output := result.Task.Output()
				for _, e := range d.GetEdgesFrom(result.Task) {
					e.To().(Task).Input(output...)
				}
				d.DeleteVertex(result.Task)
			}
		}
	}(d)

	return resultChan, nil
}

func (p *Plan) startTask(ctx context.Context, t Task, resultChan chan<- TaskResult) {
	if p.Status(t) != StatusPending {
		return
	}
	p.started.Add(1)
	go func(ctx context.Context, task Task) {
		defer p.started.Add(-1)
		if err := task.Run(ctx); err != nil {
			p.setStatus(task, StatusFailed)
			resultChan <- TaskResult{Task: task, Err: err, Status: StatusFailed}
			return
		}
		p.setStatus(task, StatusSuccess)
		resultChan <- TaskResult{Task: task, Err: nil, Status: StatusSuccess}
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
