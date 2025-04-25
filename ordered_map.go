package coll

import (
	"iter"
	"slices"
	"sync"
)

// NewOrderedMap returns a new instance of OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	om := &OrderedMap[K, V]{
		keys:  make([]K, 0),
		dirty: map[K]V{},
		mux:   sync.RWMutex{},
	}
	return om
}

// OrderedMap represents a map that preserves insertion order of keys.
// It is safe for concurrent use.
type OrderedMap[K comparable, V any] struct {
	dirty map[K]V
	keys  []K
	mux   sync.RWMutex
}

// Get retrieves the value associated with the given key.
// The second return value indicates whether the key was found.
// It is safe for concurrent use.
func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	val, ok := m.dirty[key]
	return val, ok
}

// Put inserts the key-value pair into the map if the key does not already exist.
// The insertion order of keys is preserved. It is safe for concurrent use.
func (m *OrderedMap[K, V]) Put(key K, value V) {
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, found := m.dirty[key]; found {
		return
	}
	m.dirty[key] = value
	m.keys = append(m.keys, key)
}

func (m *OrderedMap[K, V]) unsafeKeysIterator() iter.Seq[K] {
	return slices.Values(m.keys)
}

// Keys returns an iterator over the keys in insertion order.
// It is safe for concurrent use.
func (m *OrderedMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		m.mux.RLock()
		defer m.mux.RUnlock()
		m.unsafeKeysIterator()(yield)
	}
}

// Values returns an iterator over the values in insertion order of their corresponding keys.
// It is safe for concurrent use.
func (m *OrderedMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		m.mux.RLock()
		defer m.mux.RUnlock()
		for key := range m.unsafeKeysIterator() {
			if !yield(m.dirty[key]) {
				return
			}
		}
	}
}

// All returns an iterator over key-value pairs in insertion order.
// It is safe for concurrent use.
func (m *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.mux.RLock()
		defer m.mux.RUnlock()
		for key := range m.unsafeKeysIterator() {
			val := m.dirty[key]
			if !yield(key, val) {
				return
			}
		}
	}
}
