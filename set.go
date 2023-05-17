package dag

import "sync"

type Set[T comparable] interface {
	Add(...T)
	Delete(...T)
	Contains(T) bool
	List() []T
	Len() int
	Duplicate() Set[T]
}

type set[T comparable] struct {
	m map[T]struct{}
	l sync.RWMutex
}

func NewSet[T comparable](vs ...T) Set[T] {
	s := &set[T]{
		m: make(map[T]struct{}),
		l: sync.RWMutex{},
	}
	s.Add(vs...)
	return s
}

func (s *set[T]) Add(t ...T) {
	s.l.Lock()
	defer s.l.Unlock()
	for _, v := range t {
		s.m[v] = struct{}{}
	}
}

func (s *set[T]) Delete(t ...T) {
	s.l.Lock()
	defer s.l.Unlock()
	for _, v := range t {
		delete(s.m, v)
	}
}

func (s *set[T]) Contains(t T) bool {
	s.l.RLock()
	defer s.l.RUnlock()
	_, ok := s.m[t]
	return ok
}

func (s *set[T]) List() []T {
	s.l.RLock()
	defer s.l.RUnlock()
	l := make([]T, 0, len(s.m))
	for k := range s.m {
		l = append(l, k)
	}
	return l
}

func (s *set[T]) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return len(s.m)
}

func (s *set[T]) Duplicate() Set[T] {
	s.l.RLock()
	defer s.l.RUnlock()
	m := make(map[T]struct{})
	for k, v := range s.m {
		m[k] = v
	}
	return &set[T]{
		m: m,
		l: sync.RWMutex{},
	}
}
