package task_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	r "github.com/stretchr/testify/require"
	"github.com/xuxife/dag/task"
)

func TestPlan(t *testing.T) {
	ctx := context.Background()
	a := task.Func1_1(func(ctx context.Context, d int) (string, error) {
		return fmt.Sprintf("a: %d", d), nil
	})
	b := task.Func1_1(func(ctx context.Context, s string) (string, error) {
		return fmt.Sprintf("b: %s", s), nil
	})
	c := task.Func(func(ctx context.Context) error {
		return errors.New("")
	})

	t.Run("task.Add tasks independently", func(t *testing.T) {
		// task.Add tasks don't have any dependency
		// aka. they will be started at the same time
		//   /--> a
		// - |--> b
		//   \--> c
		p, err := task.New(
			task.Add(a, b, c),
			task.Add(a), // add is idempotent
		)
		r.NoError(t, err)
		// task input can be set by task.Input
		a.Input(task.Input(1))
		// TInput is type safe, but limit to task.Func family
		b.TInput("hello")

		run, err := p.Start(ctx)
		r.NoError(t, err)
		// tasks status can be checked
		// by p.Status or p.Statuses
		t.Log(p.Status(a))
		for task, status := range p.Statuses() {
			t.Log(task, status)
		}
		// run is channel of task.Result
		// task.Result can be process by iterating the channel
		for res := range run {
			t.Log(res.Err)
			t.Log(res.Status)
			t.Log(res.Task)
		}
		// or just use run.Wait
		err = run.Wait()
		r.Error(t, err)
		// output can be retrieved by task.Output
		t.Log(a.Output()...)
		// TOutput is type safe, but limit to task.Func family
		bOutput := b.TOutput()
		r.Equal(t, "b: hello", bOutput)
	})
	t.Run(".Then and .DependsOn will add dependency", func(t *testing.T) {
		// a ---> b ---> c
		p, err := task.New(
			task.Add(a).Then(b),
			task.Add(c).DependsOn(b),
		)
		r.NoError(t, err)
		a.TInput(123)
		run, err := p.Start(ctx)
		r.NoError(t, err)
		r.Error(t, run.Wait()) // because c will return error
		// but a and b will still be executed
		r.Equal(t, "a: 123", a.TOutput())
		r.Equal(t, "b: a: 123", b.TOutput())
	})
	t.Run("task.Pipelines will add dependency between pipelines", func(t *testing.T) {
		// a ---> b ---> c
		p, err := task.New(
			task.Pipelines(a, b, c),
		)
		r.NoError(t, err)
		a.TInput(123)
		run, err := p.Start(ctx)
		r.NoError(t, err)
		r.Error(t, run.Wait()) // because c will return error
		// but a and b will still be executed
		r.Equal(t, "a: 123", a.TOutput())
		r.Equal(t, "b: a: 123", b.TOutput())
	})
	t.Run("cycle dependency will be detected", func(t *testing.T) {
		//   a
		//  / \
		// b - c
		_, err := task.New(
			task.Pipelines(a, b, c, a),
		)
		r.ErrorContains(t, err, "cycle")
	})
	t.Run("multiple input should use .UseInput to customize input", func(t *testing.T) {
		d := task.Func0_2(func(ctx context.Context) (int, string, error) {
			return 123, "d: hello", nil
		})
		// a --\
		// d --|--> b
		p, err := task.New(
			task.Add(a, d).Then(b),
		)
		r.NoError(t, err)
		a.TInput(321)
		b.UseInput(func(ts []task.Task) {
			var aOut, dOut string
			for _, ta := range ts {
				switch {
				case ta == a:
					aOut = a.TOutput() // you can just get output from a, since here, a and d are finished
				case ta == d:
					var dOutInt int
					dOutInt, dOut = d.TOutput()
					r.Equal(t, 123, dOutInt)
				}
			}
			b.TInput(fmt.Sprintf("%s, %s", aOut, dOut))
		})
		run, err := p.Start(ctx)
		r.NoError(t, err)
		r.NoError(t, run.Wait())
		r.Equal(t, "b: a: 321, d: hello", b.TOutput())
	})
}
