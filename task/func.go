package task

import (
	"context"
)

func VoidTask(name string) *Base {
	return Func(name, func(_ context.Context) error { return nil })
}

func Func(name string, f func(context.Context) error) *Base {
	return &Base{
		Name: name,
		F:    f,
	}
}

type Base struct {
	Name      string
	F         func(context.Context) error
	InputFunc func([]Task) // InputFunc overwrites .Input
}

func (t *Base) Run(ctx context.Context) error {
	return t.F(ctx)
}

func (t *Base) Input(ts ...Task) {
	if t.InputFunc != nil {
		t.InputFunc(ts)
	}
}

func (t *Base) Output() []any {
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
			panic("default Input func only accept 1 task, use customized input function via setting Base.InputFunc")
		}
	}
}
