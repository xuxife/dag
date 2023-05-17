package task

import "context"

type func0_1[O1 any] struct {
	BaseTask
	out1 O1
}

func Func0_1[O1 any](f func(context.Context) (O1, error)) *func0_1[O1] {
	t := &func0_1[O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx)
		return err
	})
	return t
}

func (t *func0_1[O1]) Input(in ...any) {
}

func (t *func0_1[O1]) Output() []any {
	return []any{t.out1}
}

func (t *func0_1[O1]) TInput() {
}

func (t *func0_1[O1]) TOutput() O1 {
	return t.out1
}

type func0_2[O1, O2 any] struct {
	BaseTask
	out1 O1
	out2 O2
}

func Func0_2[O1, O2 any](f func(context.Context) (O1, O2, error)) *func0_2[O1, O2] {
	t := &func0_2[O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx)
		return err
	})
	return t
}

func (t *func0_2[O1, O2]) Input(in ...any) {
}

func (t *func0_2[O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func0_2[O1, O2]) TInput() {
}

func (t *func0_2[O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func0_3[O1, O2, O3 any] struct {
	BaseTask
	out1 O1
	out2 O2
	out3 O3
}

func Func0_3[O1, O2, O3 any](f func(context.Context) (O1, O2, O3, error)) *func0_3[O1, O2, O3] {
	t := &func0_3[O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx)
		return err
	})
	return t
}

func (t *func0_3[O1, O2, O3]) Input(in ...any) {
}

func (t *func0_3[O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func0_3[O1, O2, O3]) TInput() {
}

func (t *func0_3[O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func0_4[O1, O2, O3, O4 any] struct {
	BaseTask
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func0_4[O1, O2, O3, O4 any](f func(context.Context) (O1, O2, O3, O4, error)) *func0_4[O1, O2, O3, O4] {
	t := &func0_4[O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx)
		return err
	})
	return t
}

func (t *func0_4[O1, O2, O3, O4]) Input(in ...any) {
}

func (t *func0_4[O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func0_4[O1, O2, O3, O4]) TInput() {
}

func (t *func0_4[O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func0_5[O1, O2, O3, O4, O5 any] struct {
	BaseTask
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func0_5[O1, O2, O3, O4, O5 any](f func(context.Context) (O1, O2, O3, O4, O5, error)) *func0_5[O1, O2, O3, O4, O5] {
	t := &func0_5[O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx)
		return err
	})
	return t
}

func (t *func0_5[O1, O2, O3, O4, O5]) Input(in ...any) {
}

func (t *func0_5[O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func0_5[O1, O2, O3, O4, O5]) TInput() {
}

func (t *func0_5[O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}

type func1_0[I1 any] struct {
	BaseTask
	in1 I1
}

func Func1_0[I1 any](f func(context.Context, I1) error) *func1_0[I1] {
	t := &func1_0[I1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_0[I1]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_0[I1]) Output() []any {
	return []any{}
}

func (t *func1_0[I1]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_0[I1]) TOutput() {
	return
}

type func1_1[I1, O1 any] struct {
	BaseTask
	in1  I1
	out1 O1
}

func Func1_1[I1, O1 any](f func(context.Context, I1) (O1, error)) *func1_1[I1, O1] {
	t := &func1_1[I1, O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_1[I1, O1]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_1[I1, O1]) Output() []any {
	return []any{t.out1}
}

func (t *func1_1[I1, O1]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_1[I1, O1]) TOutput() O1 {
	return t.out1
}

type func1_2[I1, O1, O2 any] struct {
	BaseTask
	in1  I1
	out1 O1
	out2 O2
}

func Func1_2[I1, O1, O2 any](f func(context.Context, I1) (O1, O2, error)) *func1_2[I1, O1, O2] {
	t := &func1_2[I1, O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_2[I1, O1, O2]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_2[I1, O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func1_2[I1, O1, O2]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_2[I1, O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func1_3[I1, O1, O2, O3 any] struct {
	BaseTask
	in1  I1
	out1 O1
	out2 O2
	out3 O3
}

func Func1_3[I1, O1, O2, O3 any](f func(context.Context, I1) (O1, O2, O3, error)) *func1_3[I1, O1, O2, O3] {
	t := &func1_3[I1, O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_3[I1, O1, O2, O3]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_3[I1, O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func1_3[I1, O1, O2, O3]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_3[I1, O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func1_4[I1, O1, O2, O3, O4 any] struct {
	BaseTask
	in1  I1
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func1_4[I1, O1, O2, O3, O4 any](f func(context.Context, I1) (O1, O2, O3, O4, error)) *func1_4[I1, O1, O2, O3, O4] {
	t := &func1_4[I1, O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_4[I1, O1, O2, O3, O4]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_4[I1, O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func1_4[I1, O1, O2, O3, O4]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_4[I1, O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func1_5[I1, O1, O2, O3, O4, O5 any] struct {
	BaseTask
	in1  I1
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func1_5[I1, O1, O2, O3, O4, O5 any](f func(context.Context, I1) (O1, O2, O3, O4, O5, error)) *func1_5[I1, O1, O2, O3, O4, O5] {
	t := &func1_5[I1, O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx, t.in1)
		return err
	})
	return t
}

func (t *func1_5[I1, O1, O2, O3, O4, O5]) Input(in ...any) {
	t.in1 = in[1].(I1)
}

func (t *func1_5[I1, O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func1_5[I1, O1, O2, O3, O4, O5]) TInput(in1 I1) {
	t.in1 = in1
}

func (t *func1_5[I1, O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}

type func2_0[I1, I2 any] struct {
	BaseTask
	in1 I1
	in2 I2
}

func Func2_0[I1, I2 any](f func(context.Context, I1, I2) error) *func2_0[I1, I2] {
	t := &func2_0[I1, I2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_0[I1, I2]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_0[I1, I2]) Output() []any {
	return []any{}
}

func (t *func2_0[I1, I2]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_0[I1, I2]) TOutput() {
	return
}

type func2_1[I1, I2, O1 any] struct {
	BaseTask
	in1  I1
	in2  I2
	out1 O1
}

func Func2_1[I1, I2, O1 any](f func(context.Context, I1, I2) (O1, error)) *func2_1[I1, I2, O1] {
	t := &func2_1[I1, I2, O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_1[I1, I2, O1]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_1[I1, I2, O1]) Output() []any {
	return []any{t.out1}
}

func (t *func2_1[I1, I2, O1]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_1[I1, I2, O1]) TOutput() O1 {
	return t.out1
}

type func2_2[I1, I2, O1, O2 any] struct {
	BaseTask
	in1  I1
	in2  I2
	out1 O1
	out2 O2
}

func Func2_2[I1, I2, O1, O2 any](f func(context.Context, I1, I2) (O1, O2, error)) *func2_2[I1, I2, O1, O2] {
	t := &func2_2[I1, I2, O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_2[I1, I2, O1, O2]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_2[I1, I2, O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func2_2[I1, I2, O1, O2]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_2[I1, I2, O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func2_3[I1, I2, O1, O2, O3 any] struct {
	BaseTask
	in1  I1
	in2  I2
	out1 O1
	out2 O2
	out3 O3
}

func Func2_3[I1, I2, O1, O2, O3 any](f func(context.Context, I1, I2) (O1, O2, O3, error)) *func2_3[I1, I2, O1, O2, O3] {
	t := &func2_3[I1, I2, O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_3[I1, I2, O1, O2, O3]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_3[I1, I2, O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func2_3[I1, I2, O1, O2, O3]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_3[I1, I2, O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func2_4[I1, I2, O1, O2, O3, O4 any] struct {
	BaseTask
	in1  I1
	in2  I2
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func2_4[I1, I2, O1, O2, O3, O4 any](f func(context.Context, I1, I2) (O1, O2, O3, O4, error)) *func2_4[I1, I2, O1, O2, O3, O4] {
	t := &func2_4[I1, I2, O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_4[I1, I2, O1, O2, O3, O4]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_4[I1, I2, O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func2_4[I1, I2, O1, O2, O3, O4]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_4[I1, I2, O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func2_5[I1, I2, O1, O2, O3, O4, O5 any] struct {
	BaseTask
	in1  I1
	in2  I2
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func2_5[I1, I2, O1, O2, O3, O4, O5 any](f func(context.Context, I1, I2) (O1, O2, O3, O4, O5, error)) *func2_5[I1, I2, O1, O2, O3, O4, O5] {
	t := &func2_5[I1, I2, O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx, t.in1, t.in2)
		return err
	})
	return t
}

func (t *func2_5[I1, I2, O1, O2, O3, O4, O5]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
}

func (t *func2_5[I1, I2, O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func2_5[I1, I2, O1, O2, O3, O4, O5]) TInput(in1 I1, in2 I2) {
	t.in1, t.in2 = in1, in2
}

func (t *func2_5[I1, I2, O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}

type func3_0[I1, I2, I3 any] struct {
	BaseTask
	in1 I1
	in2 I2
	in3 I3
}

func Func3_0[I1, I2, I3 any](f func(context.Context, I1, I2, I3) error) *func3_0[I1, I2, I3] {
	t := &func3_0[I1, I2, I3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_0[I1, I2, I3]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_0[I1, I2, I3]) Output() []any {
	return []any{}
}

func (t *func3_0[I1, I2, I3]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_0[I1, I2, I3]) TOutput() {
	return
}

type func3_1[I1, I2, I3, O1 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	out1 O1
}

func Func3_1[I1, I2, I3, O1 any](f func(context.Context, I1, I2, I3) (O1, error)) *func3_1[I1, I2, I3, O1] {
	t := &func3_1[I1, I2, I3, O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_1[I1, I2, I3, O1]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_1[I1, I2, I3, O1]) Output() []any {
	return []any{t.out1}
}

func (t *func3_1[I1, I2, I3, O1]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_1[I1, I2, I3, O1]) TOutput() O1 {
	return t.out1
}

type func3_2[I1, I2, I3, O1, O2 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	out1 O1
	out2 O2
}

func Func3_2[I1, I2, I3, O1, O2 any](f func(context.Context, I1, I2, I3) (O1, O2, error)) *func3_2[I1, I2, I3, O1, O2] {
	t := &func3_2[I1, I2, I3, O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_2[I1, I2, I3, O1, O2]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_2[I1, I2, I3, O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func3_2[I1, I2, I3, O1, O2]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_2[I1, I2, I3, O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func3_3[I1, I2, I3, O1, O2, O3 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	out1 O1
	out2 O2
	out3 O3
}

func Func3_3[I1, I2, I3, O1, O2, O3 any](f func(context.Context, I1, I2, I3) (O1, O2, O3, error)) *func3_3[I1, I2, I3, O1, O2, O3] {
	t := &func3_3[I1, I2, I3, O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_3[I1, I2, I3, O1, O2, O3]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_3[I1, I2, I3, O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func3_3[I1, I2, I3, O1, O2, O3]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_3[I1, I2, I3, O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func3_4[I1, I2, I3, O1, O2, O3, O4 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func3_4[I1, I2, I3, O1, O2, O3, O4 any](f func(context.Context, I1, I2, I3) (O1, O2, O3, O4, error)) *func3_4[I1, I2, I3, O1, O2, O3, O4] {
	t := &func3_4[I1, I2, I3, O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_4[I1, I2, I3, O1, O2, O3, O4]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_4[I1, I2, I3, O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func3_4[I1, I2, I3, O1, O2, O3, O4]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_4[I1, I2, I3, O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func3_5[I1, I2, I3, O1, O2, O3, O4, O5 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func3_5[I1, I2, I3, O1, O2, O3, O4, O5 any](f func(context.Context, I1, I2, I3) (O1, O2, O3, O4, O5, error)) *func3_5[I1, I2, I3, O1, O2, O3, O4, O5] {
	t := &func3_5[I1, I2, I3, O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx, t.in1, t.in2, t.in3)
		return err
	})
	return t
}

func (t *func3_5[I1, I2, I3, O1, O2, O3, O4, O5]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
}

func (t *func3_5[I1, I2, I3, O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func3_5[I1, I2, I3, O1, O2, O3, O4, O5]) TInput(in1 I1, in2 I2, in3 I3) {
	t.in1, t.in2, t.in3 = in1, in2, in3
}

func (t *func3_5[I1, I2, I3, O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}

type func4_0[I1, I2, I3, I4 any] struct {
	BaseTask
	in1 I1
	in2 I2
	in3 I3
	in4 I4
}

func Func4_0[I1, I2, I3, I4 any](f func(context.Context, I1, I2, I3, I4) error) *func4_0[I1, I2, I3, I4] {
	t := &func4_0[I1, I2, I3, I4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_0[I1, I2, I3, I4]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_0[I1, I2, I3, I4]) Output() []any {
	return []any{}
}

func (t *func4_0[I1, I2, I3, I4]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_0[I1, I2, I3, I4]) TOutput() {
	return
}

type func4_1[I1, I2, I3, I4, O1 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	out1 O1
}

func Func4_1[I1, I2, I3, I4, O1 any](f func(context.Context, I1, I2, I3, I4) (O1, error)) *func4_1[I1, I2, I3, I4, O1] {
	t := &func4_1[I1, I2, I3, I4, O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_1[I1, I2, I3, I4, O1]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_1[I1, I2, I3, I4, O1]) Output() []any {
	return []any{t.out1}
}

func (t *func4_1[I1, I2, I3, I4, O1]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_1[I1, I2, I3, I4, O1]) TOutput() O1 {
	return t.out1
}

type func4_2[I1, I2, I3, I4, O1, O2 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	out1 O1
	out2 O2
}

func Func4_2[I1, I2, I3, I4, O1, O2 any](f func(context.Context, I1, I2, I3, I4) (O1, O2, error)) *func4_2[I1, I2, I3, I4, O1, O2] {
	t := &func4_2[I1, I2, I3, I4, O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_2[I1, I2, I3, I4, O1, O2]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_2[I1, I2, I3, I4, O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func4_2[I1, I2, I3, I4, O1, O2]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_2[I1, I2, I3, I4, O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func4_3[I1, I2, I3, I4, O1, O2, O3 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	out1 O1
	out2 O2
	out3 O3
}

func Func4_3[I1, I2, I3, I4, O1, O2, O3 any](f func(context.Context, I1, I2, I3, I4) (O1, O2, O3, error)) *func4_3[I1, I2, I3, I4, O1, O2, O3] {
	t := &func4_3[I1, I2, I3, I4, O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_3[I1, I2, I3, I4, O1, O2, O3]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_3[I1, I2, I3, I4, O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func4_3[I1, I2, I3, I4, O1, O2, O3]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_3[I1, I2, I3, I4, O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func4_4[I1, I2, I3, I4, O1, O2, O3, O4 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func4_4[I1, I2, I3, I4, O1, O2, O3, O4 any](f func(context.Context, I1, I2, I3, I4) (O1, O2, O3, O4, error)) *func4_4[I1, I2, I3, I4, O1, O2, O3, O4] {
	t := &func4_4[I1, I2, I3, I4, O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_4[I1, I2, I3, I4, O1, O2, O3, O4]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_4[I1, I2, I3, I4, O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func4_4[I1, I2, I3, I4, O1, O2, O3, O4]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_4[I1, I2, I3, I4, O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5 any](f func(context.Context, I1, I2, I3, I4) (O1, O2, O3, O4, O5, error)) *func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5] {
	t := &func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx, t.in1, t.in2, t.in3, t.in4)
		return err
	})
	return t
}

func (t *func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
}

func (t *func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5]) TInput(in1 I1, in2 I2, in3 I3, in4 I4) {
	t.in1, t.in2, t.in3, t.in4 = in1, in2, in3, in4
}

func (t *func4_5[I1, I2, I3, I4, O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}

type func5_0[I1, I2, I3, I4, I5 any] struct {
	BaseTask
	in1 I1
	in2 I2
	in3 I3
	in4 I4
	in5 I5
}

func Func5_0[I1, I2, I3, I4, I5 any](f func(context.Context, I1, I2, I3, I4, I5) error) *func5_0[I1, I2, I3, I4, I5] {
	t := &func5_0[I1, I2, I3, I4, I5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_0[I1, I2, I3, I4, I5]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_0[I1, I2, I3, I4, I5]) Output() []any {
	return []any{}
}

func (t *func5_0[I1, I2, I3, I4, I5]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_0[I1, I2, I3, I4, I5]) TOutput() {
	return
}

type func5_1[I1, I2, I3, I4, I5, O1 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	in5  I5
	out1 O1
}

func Func5_1[I1, I2, I3, I4, I5, O1 any](f func(context.Context, I1, I2, I3, I4, I5) (O1, error)) *func5_1[I1, I2, I3, I4, I5, O1] {
	t := &func5_1[I1, I2, I3, I4, I5, O1]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_1[I1, I2, I3, I4, I5, O1]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_1[I1, I2, I3, I4, I5, O1]) Output() []any {
	return []any{t.out1}
}

func (t *func5_1[I1, I2, I3, I4, I5, O1]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_1[I1, I2, I3, I4, I5, O1]) TOutput() O1 {
	return t.out1
}

type func5_2[I1, I2, I3, I4, I5, O1, O2 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	in5  I5
	out1 O1
	out2 O2
}

func Func5_2[I1, I2, I3, I4, I5, O1, O2 any](f func(context.Context, I1, I2, I3, I4, I5) (O1, O2, error)) *func5_2[I1, I2, I3, I4, I5, O1, O2] {
	t := &func5_2[I1, I2, I3, I4, I5, O1, O2]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_2[I1, I2, I3, I4, I5, O1, O2]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_2[I1, I2, I3, I4, I5, O1, O2]) Output() []any {
	return []any{t.out1, t.out2}
}

func (t *func5_2[I1, I2, I3, I4, I5, O1, O2]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_2[I1, I2, I3, I4, I5, O1, O2]) TOutput() (O1, O2) {
	return t.out1, t.out2
}

type func5_3[I1, I2, I3, I4, I5, O1, O2, O3 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	in5  I5
	out1 O1
	out2 O2
	out3 O3
}

func Func5_3[I1, I2, I3, I4, I5, O1, O2, O3 any](f func(context.Context, I1, I2, I3, I4, I5) (O1, O2, O3, error)) *func5_3[I1, I2, I3, I4, I5, O1, O2, O3] {
	t := &func5_3[I1, I2, I3, I4, I5, O1, O2, O3]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_3[I1, I2, I3, I4, I5, O1, O2, O3]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_3[I1, I2, I3, I4, I5, O1, O2, O3]) Output() []any {
	return []any{t.out1, t.out2, t.out3}
}

func (t *func5_3[I1, I2, I3, I4, I5, O1, O2, O3]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_3[I1, I2, I3, I4, I5, O1, O2, O3]) TOutput() (O1, O2, O3) {
	return t.out1, t.out2, t.out3
}

type func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	in5  I5
	out1 O1
	out2 O2
	out3 O3
	out4 O4
}

func Func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4 any](f func(context.Context, I1, I2, I3, I4, I5) (O1, O2, O3, O4, error)) *func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4] {
	t := &func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4}
}

func (t *func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_4[I1, I2, I3, I4, I5, O1, O2, O3, O4]) TOutput() (O1, O2, O3, O4) {
	return t.out1, t.out2, t.out3, t.out4
}

type func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5 any] struct {
	BaseTask
	in1  I1
	in2  I2
	in3  I3
	in4  I4
	in5  I5
	out1 O1
	out2 O2
	out3 O3
	out4 O4
	out5 O5
}

func Func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5 any](f func(context.Context, I1, I2, I3, I4, I5) (O1, O2, O3, O4, O5, error)) *func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5] {
	t := &func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5]{}
	t.BaseTask = *newFunc(func(ctx context.Context) error {
		var err error
		t.out1, t.out2, t.out3, t.out4, t.out5, err = f(ctx, t.in1, t.in2, t.in3, t.in4, t.in5)
		return err
	})
	return t
}

func (t *func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5]) Input(in ...any) {
	t.in1 = in[1].(I1)
	t.in2 = in[2].(I2)
	t.in3 = in[3].(I3)
	t.in4 = in[4].(I4)
	t.in5 = in[5].(I5)
}

func (t *func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5]) Output() []any {
	return []any{t.out1, t.out2, t.out3, t.out4, t.out5}
}

func (t *func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5]) TInput(in1 I1, in2 I2, in3 I3, in4 I4, in5 I5) {
	t.in1, t.in2, t.in3, t.in4, t.in5 = in1, in2, in3, in4, in5
}

func (t *func5_5[I1, I2, I3, I4, I5, O1, O2, O3, O4, O5]) TOutput() (O1, O2, O3, O4, O5) {
	return t.out1, t.out2, t.out3, t.out4, t.out5
}
