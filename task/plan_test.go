package task_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/require"
	r "github.com/stretchr/testify/require"
	"github.com/xuxife/dag/task"
)

// # Task
//
// Let's start defining some tasks.
//
// A task should implement interface Task,
// which includes 3 methods: Run, Input, Output.
// Always embed the BaseTask to get the default implementation.
//
//   - Run(context.Context) error
//     the main logic of the task
//
//   - Input(...Task)
//     set the input of the task, will be called before Run.
//     the arguments are pre-order tasks.
//     i.e. if taskA depends on taskB, then taskA would receive: taskA.Input([]Task{taskB}),
//     then taskA could get output from taskB via taskB.Output(),
//     then taskA could do some pre-process before setting itself's input.
//
//   - Output() []any
//     get the output from the task.

// taskAdd accepts all int inputs, and output the sum of them.
type taskAdd struct {
	task.BaseTask
	args []int
	out  int
}

func (a *taskAdd) Run(ctx context.Context) error {
	a.out = 0
	for _, arg := range a.args {
		a.out += arg
	}
	return nil
}

func (a *taskAdd) Input(ts ...task.Task) {
	for _, t := range ts {
		for _, out := range t.Output() {
			// accepts all int output from pre-order tasks
			if number, ok := out.(int); ok {
				a.args = append(a.args, number)
			}
		}
	}
}

func (a *taskAdd) Output() []any {
	return []any{a.out}
}

func TestExample(t *testing.T) {
	ctx := context.Background()
	// ## Create Task
	// a normal pattern to interacts with Task is:
	add := &taskAdd{}
	add.Input(task.Input("a", 1), task.Input("b", 2))
	require.NoError(t, add.Run(ctx))
	require.EqualValues(t, 3, add.Output()[0])

	// ## Func
	//
	// task package also provides a family of built-in tasks: Func
	// A base function accepts a context.Context and returns an error.
	helloWorld := task.Func("helloWorld", func(ctx context.Context) error {
		fmt.Println("Hello World!")
		return nil
	})
	err := helloWorld.Run(ctx)
	require.NoError(t, err)
	// Other candidate in Func family are generics types,
	// name as FuncX_Y with X input and Y output.
	upper := task.Func1_1("upper", func(_ context.Context, s string) (string, error) {
		return strings.ToUpper(s), nil
	})
	concate := task.Func2_1("concate", func(ctx context.Context, a, b string) (string, error) {
		return fmt.Sprintf("%s %s", a, b), nil
	})
	// set Input and get Output
	upper.Input(task.Input("groot", "i am groot")) // see below for more about Input
	err = upper.Run(ctx)
	require.NoError(t, err)
	output := upper.Output()
	require.Len(t, output, 1)
	require.Equal(t, "I AM GROOT", output[0])
	// Func family also provides a type safe way to customize Input and get Output
	require.Equal(t, "I AM GROOT", upper.TOutput())
	concate.TInput("yes", "madam")

	// ## Input
	//
	// task.Input is a task only generates output.
	var foo task.Task = task.Input("foo", "hello")
	// it's equivalent to
	foo = task.Func0_1("foo", func(_ context.Context) (string, error) {
		return "hello", nil
	})
	require.NoError(t, foo.Run(ctx))
	require.Equal(t, "hello", foo.Output()[0])

	// # Plan
	//
	// Plan organizes tasks into a DAG, and execute them in topological order.
	//
	// ## Create Plan
	//
	// Plan can be created by connecting tasks:
	// - task.Add
	// - task.Pipelines
	// - task.New
	//
	// ### task.Add
	// task.Add adds tasks independently
	p := task.Add(
		helloWorld, upper, concate,
	)
	// --> helloWorld
	// --> upper 		// <~~ TInput("i am groot")
	// --> concate		// <~~ TInput("yes", "madam")
	p.Add(helloWorld) // add is idempotent

	// ### task.Pipelines
	// task.Pipelines adds tasks in a chain.
	task.Pipelines(
		foo, upper,
	)
	// foo (Input: "hello") ---> upper

	// ### task.New
	// task.New groups a set of Plans into single Plan,
	// it would merge the task DAGs, and check whether there is cycle.
	_, err = task.New(
		task.Add(helloWorld).Then(foo),
		task.Add(foo).Then(upper),
		task.Add(upper).Then(helloWorld),
	)
	//     helloWorld
	//     /        \
	//   foo ------ upper
	require.ErrorContains(t, err, "cycle")

	// ## Connect Tasks
	//
	// Tasks can be connected by:
	// - (*Plan).Add		// same as task.Add
	// - (*Plan).Pipelines	// same as task.Pipelines
	// - (*Plan).Then
	// - (*Plan).DependsOn
	//
	// ### .Then
	// .Then adds dependency between tasks.
	task.Add(foo).Then(upper)
	// foo ---> upper

	// ### .DependsOn
	// .DependsOn adds dependency between tasks, in opposite direction of .Then
	task.Add(upper).DependsOn(foo)
	// foo ---> upper

	// ## Execute Plan
	//
	// Plan executes tasks in topological order,
	// meaning that tasks without incoming edges will be executed first.
	// When a task is succeeded, the successors of the task without other dependencies will be executed.
	//
	// Plan can be executed by:
	// - (*Plan).Start		// start all tasks async
	// - (*Plan).Wait		// wait for all tasks to finish
	// - (*Plan).Run 		// = Start + Wait, start all tasks, and wait to finish
	p = task.Pipelines(foo, upper)
	require.NoError(t, p.Run(ctx))
	require.Equal(t, "HELLO", upper.TOutput())

	// ## Input and Output
	//
	// Plan connects tasks, and pass output from pre-order tasks to post-order tasks.
	bar := task.Input("bar", "world")
	p, err = task.New(
		task.Pipelines(foo, upper, concate),
		task.Add(bar).Then(concate),
	)
	// foo (Input: "hello") ---> upper ---> concate
	// bar (Input: "world") -------------------^
	// notice concate accepts two inputs
	// we need to customize concate.InputFunc to clarify the behavior
	concate.InputFunc = func(ts []task.Task) {
		var upperOut, barOut string
		for _, t := range ts {
			switch {
			case t == upper:
				upperOut = upper.TOutput()
			case t == bar:
				barOut = bar.Output()[0].(string)
			}
		}
		concate.TInput(upperOut, barOut)
	}
	require.NoError(t, err)
	require.NoError(t, p.Run(ctx))
	require.Equal(t, "HELLO world", concate.TOutput())
}

func TestPlan(t *testing.T) {
	ctx := context.Background()
	a := task.Func1_1("a", func(ctx context.Context, d int) (string, error) {
		return fmt.Sprintf("a: %d", d), nil
	})
	b := task.Func1_1("b", func(ctx context.Context, s string) (string, error) {
		return fmt.Sprintf("b: %s", s), nil
	})
	c := task.Func("c", func(ctx context.Context) error {
		return errors.New("c: error")
	})
	t.Run("task.Add tasks independently", func(t *testing.T) {
		// task.Add tasks don't have any dependency
		// aka. they will be started at the same time
		// --> a
		// --> b
		// --> c
		p, err := task.New(
			task.Add(a, b, c),
			task.Add(a), // add is idempotent
		)
		r.NoError(t, err)
		// task input can be set by task.Input
		p.Add(task.Input("input_a", 123)).Then(a)
		// TInput is type safe, but limit to task.Func family
		// and TInput is transparent to Plan, it directs to Task
		b.TInput("hello")
		// Then the graph becomes
		//  input_a --> a
		//          --> b  // <~~ TInput("hello")
		//          --> c

		r.NoError(t, p.Start(ctx))
		// tasks status can be checked
		// by p.Status or p.Statuses
		t.Log(p.Status(a))
		for task, status := range p.Statuses() {
			t.Log(task, status)
		}
		// use p.Wait() to wait for all tasks to finish
		r.Error(t, p.Wait()) // because c will return error
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
		r.NoError(t, p.Start(ctx))
		r.Error(t, p.Wait()) // because c will return error
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
		r.NoError(t, p.Start(ctx))
		r.Error(t, p.Wait()) // because c will return error
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
	t.Run("multiple input should customize InputFunc", func(t *testing.T) {
		d := task.Func0_2("d", func(ctx context.Context) (int, string, error) {
			return 123, "d: hello", nil
		})
		// a --\
		// d -----> b
		p, err := task.New(
			task.Add(a, d).Then(b),
		)
		r.NoError(t, err)
		a.TInput(323)
		defer gostub.New().Stub(&b.InputFunc, func(ts []task.Task) {
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
		}).Reset()
		r.NoError(t, p.Start(ctx))
		r.NoError(t, p.Wait())
		r.Equal(t, "b: a: 323, d: hello", b.TOutput())
	})
	t.Run("multiple dependencies would copy output to each task", func(t *testing.T) {
		d := task.Func1_1("d", func(ctx context.Context, s string) (string, error) {
			return fmt.Sprintf("d: %s", s), nil
		})
		//     /--> b
		// a -----> d
		p, err := task.New(
			task.Add(a).Then(b, d),
		)
		r.NoError(t, err)
		a.TInput(222)
		r.NoError(t, p.Start(ctx))
		r.NoError(t, p.Wait())
		r.Equal(t, "b: a: 222", b.TOutput())
		r.Equal(t, "d: a: 222", d.TOutput())
	})
}
