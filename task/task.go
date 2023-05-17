package task

import (
	"context"
	"fmt"

	"github.com/xuxife/dag"
)

type Task interface {
	dag.Vertex

	Run(context.Context) error

	Input(...Task)
	Output() []any
}

type TaskStatus int

const (
	StatusUnknown TaskStatus = iota
	StatusPending
	StatusRunning
	StatusSuccess
	StatusFailed
	StatusCancelled
)

func (s TaskStatus) String() string {
	switch s {
	case StatusUnknown:
		return "unknown"
	case StatusPending:
		return "pending"
	case StatusRunning:
		return "running"
	case StatusSuccess:
		return "success"
	case StatusFailed:
		return "failed"
	case StatusCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

type TaskResult struct {
	Task   Task
	Err    error
	Status TaskStatus
}

func (r TaskResult) String() string {
	if r.Err != nil {
		return fmt.Sprintf("<Task: %s, Status: %s, Err: %v>", r.Task, r.Status, r.Err)
	}
	return fmt.Sprintf("<Task: %s, Status: %s>", r.Task, r.Status)
}

func (r TaskResult) Error() string {
	return r.String()
}
