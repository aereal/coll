package coll

import (
	"iter"
	"sync"
)

// NewSet returns a new [Set] containing the provided elements.
// Duplicates in the input are ignored.
func NewSet[E comparable](els ...E) *Set[E] {
	s := &Set[E]{
		values: map[E]struct{}{},
		mux:    sync.RWMutex{},
	}
	for _, v := range els {
		s.unsafeAppend(v)
	}
	return s
}

// Set represents a set of comparable elements.
// It is safe for concurrent use.
type Set[E comparable] struct {
	values map[E]struct{}
	mux    sync.RWMutex
}

// Len returns the number of elements in the set.
func (s *Set[E]) Len() int { return len(s.values) }

// Contains reports whether the element is present in the set.
// It is safe for concurrent use.
func (s *Set[E]) Contains(el E) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.unsafeContains(el)
}

func (s *Set[E]) unsafeContains(el E) bool {
	if s.values == nil {
		s.values = map[E]struct{}{}
	}
	_, found := s.values[el]
	return found
}

// Append adds the element to the set if it does not already exist.
// It is safe for concurrent use.
func (s *Set[E]) Append(el E) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.unsafeAppend(el)
}

func (s *Set[E]) unsafeAppend(el E) {
	if s.unsafeContains(el) {
		return
	}
	s.values[el] = struct{}{}
}

// Values returns an iterator over the elements of the set.
func (s *Set[E]) Values() iter.Seq[E] {
	return func(yield func(E) bool) {
		for el := range s.values {
			if !yield(el) {
				return
			}
		}
	}
}

func (s *Set[E]) Remove(removedEl E) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if !s.unsafeContains(removedEl) {
		// short circuit
		return
	}
	delete(s.values, removedEl)
}

// Diff returns a new [Set] containing elements that are in s or other but not in both.
func (s *Set[E]) Diff(other *Set[E]) *Set[E] {
	ret := NewSet[E]()
	buildDiff(ret, s, other)
	return ret
}

// Intersect returns a new [Set] containing elements that are present in both s and other.
func (s *Set[E]) Intersect(other *Set[E]) *Set[E] {
	ret := NewSet[E]()
	buildIntersection(ret, s, other)
	return ret
}

// Union returns a new [Set] containing all elements from both s and other.
func (s *Set[E]) Union(other *Set[E]) *Set[E] {
	ret := NewSet[E]()
	buildUnion(ret, s, other)
	return ret
}
