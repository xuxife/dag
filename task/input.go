package task

func Input(name string, vs ...any) *input {
	i := &input{}
	i.values = vs
	i.Base = *VoidTask(name)
	return i
}

var _ Task = &input{}

type input struct {
	Base
	values []any
}

func (i input) Output() []any {
	return i.values
}
