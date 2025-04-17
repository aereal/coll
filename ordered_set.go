package coll

import (
	"iter"
	"sync"
)

func NewOrderedSet[E comparable](els ...E) *OrderedSet[E] {
	s := &OrderedSet[E]{existence: map[E]struct{}{}}
	for _, v := range els {
		s.unsafeAppend(v)
	}
	return s
}

type OrderedSet[E comparable] struct {
	existence map[E]struct{}
	values    []E
	mux       sync.RWMutex
}

func (s *OrderedSet[E]) Len() int { return len(s.values) }

func (s *OrderedSet[E]) Contains(el E) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.unsafeContains(el)
}

func (s *OrderedSet[E]) unsafeContains(el E) bool {
	if s.existence == nil {
		s.existence = map[E]struct{}{}
	}
	_, found := s.existence[el]
	return found
}

func (s *OrderedSet[E]) Append(el E) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.unsafeAppend(el)
}

func (s *OrderedSet[E]) unsafeAppend(el E) {
	if s.unsafeContains(el) {
		return
	}
	s.existence[el] = struct{}{}
	s.values = append(s.values, el)
}

func (s *OrderedSet[E]) Values() iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, el := range s.values {
			if !yield(el) {
				return
			}
		}
	}
}
