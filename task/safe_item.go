package task

import "sync"

type SafeItems[T any] struct {
	sync.Mutex
	Items []T
}

func NewSafeItem[T any](items ...T) *SafeItems[T] {
	return &SafeItems[T]{
		Items: items,
	}
}

func (s *SafeItems[T]) Append(items ...T) {
	s.Lock()
	defer s.Unlock()
	s.Items = append(s.Items, items...)
}

func (s *SafeItems[T]) Get() []T {
	s.Lock()
	defer s.Unlock()
	return s.Items
}
