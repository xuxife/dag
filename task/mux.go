package task

import (
	"context"
	"reflect"
)

func Mux(f any) Task {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		panic("mux: f must be a function")
	}
	fv := reflect.ValueOf(f)
	m := &mux{
		in:  make([]reflect.Value, ft.NumIn()),
		out: make([]reflect.Value, ft.NumOut()),
	}
	for i := range m.in {
		m.in[i] = reflect.Zero(ft.In(i))
	}
	for i := range m.out {
		m.out[i] = reflect.Zero(ft.Out(i))
	}
	m.BaseTask = *newFunc(func(_ context.Context) error {
		m.out = fv.Call(m.in)
		return nil
	})
	return m
}

type mux struct {
	BaseTask

	in  []reflect.Value
	out []reflect.Value
}

func (m *mux) Input(in ...any) {
	for i, v := range in {
		m.in[i] = reflect.ValueOf(v)
	}
}

func (m *mux) Output() []any {
	out := make([]any, len(m.out))
	for i, v := range m.out {
		out[i] = v.Interface()
	}
	return out
}
