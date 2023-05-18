package dag

import (
	"fmt"
)

type Vertex interface {
	fmt.Stringer
	// TODO
}

type BaseVertex struct {
	Name string
}

func NewVertex(name string) *BaseVertex {
	return &BaseVertex{name}
}

func (b *BaseVertex) String() string {
	return fmt.Sprintf("<%s>", b.Name)
}
