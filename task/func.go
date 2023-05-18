package task

import (
	"context"

	"github.com/xuxife/dag"
)

func VoidTask(name string) *BaseTask {
	return Func(name, func(_ context.Context) error { return nil })
}

func Func(name string, f func(context.Context) error) *BaseTask {
	return &BaseTask{
		BaseVertex: *dag.NewVertex(name),
		F:          f,
	}
}

type BaseTask struct {
	dag.BaseVertex
	F         func(context.Context) error
	InputFunc func([]Task) // InputFunc overwrites .Input
}

func (t *BaseTask) Run(ctx context.Context) error {
	return t.F(ctx)
}

func (t *BaseTask) Input(ts ...Task) {
	if t.InputFunc != nil {
		t.InputFunc(ts)
	}
}

func (t *BaseTask) Output() []any {
	return nil
}

//go:generate sh -c "go run ./script/main.go -num_in 5 -num_out 5 | gofmt > genfunc.go"

func BaseInputFunc(f func(...any)) func([]Task) {
	return func(ts []Task) {
		switch len(ts) {
		case 0:
		case 1:
			f(ts[0].Output()...)
		default:
			panic("default Input func only accept 1 task, use customized input function via setting BaseTask.InputFunc")
		}
	}
}
