package operator_test

import (
	"context"
	"fmt"

	"github.com/xuxife/dag"
)

// # DAG for Kubernetes Operator
//
// Now let's assume we are going to build a Kubernetes Operator,
// which watches a CRD `Plan`, and reconcile the jobs defined in Plan in topological order.
type Plan struct {
	// metav1.TypeMeta   `json:",inline"`
	// metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanSpec   `json:"spec,omitempty"`
	Status PlanStatus `json:"status,omitempty"`
}

type PlanList struct {
	// metav1.TypeMeta `json:",inline"`
	// metav1.ListMeta `json:"metadata,omitempty"`
	Items []Plan `json:"items"`
}

type PlanSpec struct {
	TaskSpecs []TaskSpec `json:"tasks,omitempty"`
}

type TaskSpec struct {
	Name      string   `json:"name,omitempty"`
	DependsOn []string `json:"dependsOn,omitempty"`
	// JobSpec batchv1.JobSpec `json:"job,omitempty"`
	JobSpec JobSpec `json:"job,omitempty"`
}

type JobSpec struct{} // mock for batchv1.JobSpec

type PlanStatus struct {
	TaskStatuses []TaskStatus `json:"tasks,omitempty"`
}

type TaskStatus struct {
	Name   string         `json:"name,omitempty"`
	Status TaskStatusType `json:"status,omitempty"`
}

type TaskStatusType string

const (
	TaskStatusTypePending   TaskStatusType = "Pending"
	TaskStatusTypeRunning   TaskStatusType = "Running"
	TaskStatusTypeSucceeded TaskStatusType = "Succeeded"
	TaskStatusTypeFailed    TaskStatusType = "Failed"
	TaskStatusTypeUnknown   TaskStatusType = "Unknown"
	TaskStatusTypeCancelled TaskStatusType = "Cancelled"
)

// ## Reconciler
// Now let's buld a reconciler for Plan.
type PlanReconciler struct {
	Client K8sClient
}

// some internal types in reconciler
type task struct {
	dag.BaseVertex // implements dag.Vertex

	TaskSpec
	Status TaskStatusType
}

func (r *PlanReconciler) Reconcile(ctx context.Context, req Request) (Result, error) {
	var plan Plan
	_ = r.Client.Get(ctx, req.NamespacedName, &plan)

	tasks := map[string]*task{}
	for _, spec := range plan.Spec.TaskSpecs {
		tasks[spec.Name] = &task{
			BaseVertex: *dag.NewVertex(spec.Name),
			TaskSpec:   spec,
		}
	}
	for _, status := range plan.Status.TaskStatuses {
		tasks[status.Name].Status = status.Status
	}

	// update status with latest job status
	// r.Client.List(ctx, &JobList{})

	// build DAG
	d, _ := dag.New()
	for _, t := range tasks {
		d.AddVertex(t)
		for _, dep := range t.DependsOn {
			d.AddEdge(dag.From(tasks[dep]).To(t)...)
		}
	}
	_, err := d.Sort()
	if err != nil {
		return Result{}, fmt.Errorf("Task dependency cycle detected: %w", err)
	}

	// remove succeeded tasks to determine which tasks to run next
	for _, v := range d.GetVertices() {
		t := v.(*task)
		switch t.Status {
		case TaskStatusTypeSucceeded:
			d.DeleteVertex(t)
		case TaskStatusTypeFailed:
			// handle failure
		}
	}

	// start tasks without dependencies and in pending
	for _, v := range d.GetVerticesNoIn() {
		t := v.(*task)
		if t.Status != TaskStatusTypePending {
			continue
		}
		// start task
		_ = r.Client.Create(ctx, t.JobSpec)
	}

	return Result{}, nil
}
