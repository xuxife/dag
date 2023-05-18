package operator_test

import (
	"context"
	"time"
)

type K8sClient interface {
	Get(ctx context.Context, key any, obj any, opts ...any) error
	List(ctx context.Context, list any, opts ...any) error
	Create(ctx context.Context, obj any, opts ...any) error
	Delete(ctx context.Context, obj any, opts ...any) error
	Update(ctx context.Context, obj any, opts ...any) error
	Patch(ctx context.Context, obj any, patch any, opts ...any) error
	DeleteAllOf(ctx context.Context, obj any, opts ...any) error
}

type Request struct {
	NamespacedName
}

type NamespacedName struct {
	Namespace string
	Name      string
}

type Result struct {
	Requeue      bool
	RequeueAfter time.Duration
}

type JobSpec struct{} // mock for batchv1.JobSpec
