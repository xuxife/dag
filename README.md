# DAG - Yet Another Directed-Acyclic-Graph Module with Task Executor

DAG is a Golang module for getting the topological order in a Directed Acyclic Graph.
The module includes a built-in task executor that supports multiple workflow patterns (spawns, pipelines, etc.).

# Features

- Link vertices verbosely to create DAG
- Topological sort
- Task executor for running tasks with dependencies
- Input and Output between linked tasks
- Type safe task handlers

# Example Code

- Task Executor: [task/plan_test.go](task/plan_test.go)
- Operator: [operator/operator_test.go](operator/operator_test.go)

# Usage

## Topological Sort a DAG

```go
d, err := dag.New(
    dag.From("A").To("B"),
    dag.From("B").To("C"),
    dag.From("D").To("C"),
)
sorted, err := d.Sort()
for _, vertices := range sorted {
    fmt.Println(vertices)
    // {"A", "D"}
    // {"B"}
    // {"C"}
}
```

## Task Executor

## Add Depedency

```go
foo := task.Input("input_foo", "hello world")
upper := task.Func1_1("upper", func(ctx context.Context, s string) (string, error) {
    return strings.ToUpper(s), nil
})
// Below 3 lines are equivalent
plan := task.Add(foo).Then(upper)
plan = task.Add(upper).DependsOn(foo)
plan = task.Pipelines(foo, upper)
```

## Run and Get Output

```go
err := plan.Run(context.Context())
upper.Output() // ["HELLO WORLD"]
```

## Bring Your Own Task

```go
type TaskAdd struct {
    task.Base
    addons []int
    result int
}

func (a *TaskAdd) Run(_ context.Context) error {
    a.result = 0
    for _, addon := range a.addons {
        a.result += addon
    }
    return nil
}

func (a *TaskAdd) Input(ts ...Task) {
    a.addons = []int{}
    for _, t := range ts {
        for _, o := range t.Output() {
            if num, ok := o.(int); ok {
                a.addons = append(a.addons, num)
            }
        }
    }
}

func (a *TaskAdd) Output() []any {
    return []any{a.result}
}

func TestAdd(t *testing.T) {
    a := &TaskAdd{}
    b := task.Input("input_b", 1)
    c := task.Input("input_c", 2)
    plan := task.Add(a).DependsOn(b, c)
    err := plan.Run(context.Background())
    assert.NoError(t, err)
    assert.EqualValues(t, 3, a.result)
}
```

## Use Built-In Func

Provide a generic type FuncX_Y (X input, Y output)

```go
input := task.Input("input_1", "hello", "world")
concate := task.Func2_1(func(_ context.Context, a, b string) (string, error) {
    return fmt.Sprintf("%s %s", a, b), nil
})

plan := task.Pipelines(input, concate)
err := plan.Run(context.Background())
assert.NoError(t, err)
assert.Equal(t, "hello world", concate.Output()[0].(string))
// or use TOutput() to get type safe output
assert.Equal(t, "hello world", concate.TOutput())
```