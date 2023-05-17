package task

import (
	"context"

	"github.com/xuxife/dag"
)

func VoidTask() Task {
	return Func(func(_ context.Context) error { return nil })
}

func Func(f func(context.Context) error) Task {
	return newFunc(f)
}

func newFunc(f func(context.Context) error) *BaseTask {
	return &BaseTask{BaseVertex: *dag.NewVertex(), F: f}
}

//go:generate sh -c "go run ./script/main.go -num_in 5 -num_out 5 | gofmt > genfunc.go"
