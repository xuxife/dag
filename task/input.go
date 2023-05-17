package task

func Input(name string, vs ...any) *input {
	i := &input{}
	i.values = vs
	i.BaseTask = *VoidTask(name)
	return i
}

var _ Task = &input{}

type input struct {
	BaseTask
	values []any
}

func (i input) Output() []any {
	return i.values
}
