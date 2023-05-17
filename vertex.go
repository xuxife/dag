package dag

import (
	"fmt"

	"github.com/google/uuid"
)

type Vertex interface {
	fmt.Stringer
	comparable()
}

type BaseVertex struct {
	ID uuid.UUID
}

func NewVertex() *BaseVertex {
	return &BaseVertex{
		ID: uuid.New(),
	}
}

func (b *BaseVertex) comparable() {}

func (b *BaseVertex) String() string {
	return fmt.Sprintf("<ID: %s>", b.ID)
}
