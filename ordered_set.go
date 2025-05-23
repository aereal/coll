package coll

import (
	"iter"
	"sync"
)

// NewOrderedSet returns a new [OrderedSet] containing the provided elements.
// Duplicates in the input are ignored, and insertion order is preserved.
func NewOrderedSet[E comparable](els ...E) *OrderedSet[E] {
	s := &OrderedSet[E]{
		existence: map[E]struct{}{},
		mux:       sync.RWMutex{},
		values:    []E{},
	}
	for _, v := range els {
		s.unsafeAppend(v)
	}
	return s
}

// OrderedSet represents a set of comparable elements that maintains insertion order.
// It is safe for concurrent use.
type OrderedSet[E comparable] struct {
	existence map[E]struct{}
	values    []E
	mux       sync.RWMutex
}

// Len returns the number of elements in the set.
func (s *OrderedSet[E]) Len() int { return len(s.values) }

// Contains reports whether the element is present in the set.
// It is safe for concurrent use.
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

// Append adds the element to the set if it does not already exist.
// The insertion order is preserved. It is safe for concurrent use.
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

// Values returns an iterator over the elements of the set in insertion order.
func (s *OrderedSet[E]) Values() iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, el := range s.values {
			if !yield(el) {
				return
			}
		}
	}
}
