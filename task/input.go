package task

func Input(vs ...any) *input {
	i := &input{}
	i.values = vs
	i.BaseTask = *VoidTask()
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
