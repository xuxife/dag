package task_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/xuxife/dag/task"
)

func TestPlan(t *testing.T) {
	ctx := context.Background()
	a := task.Func1_1(func(ctx context.Context, d int) (int, error) {
		return d + 10, nil
	})
	b := task.Func1_1(func(ctx context.Context, s string) (string, error) {
		return fmt.Sprintf("b: %s", s), nil
	})
	c := task.Func(func(ctx context.Context) error {
		return errors.New("")
	})
	p, err := task.New(
		// task.Add(a),
		// task.Add(a).DependsOn(b),
		task.Add(a).Then(c),
		task.Add(a).Then(task.Mux(func(d int) string {
			return fmt.Sprintf("mux: %d", d)
		})).Then(b),
	// task.Pipelines(a, b),
	)
	if err != nil {
		t.Fatal(err)
	}
	a.Input(100)
	ch, err := p.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	s := p.Statuses()
	t.Log(s)
	for r := range ch {
		t.Log(r)
	}
	s = p.Statuses()
	t.Log(s)
	t.Log(b.Output()...)
}
