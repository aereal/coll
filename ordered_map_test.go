package coll_test

import (
	"reflect"
	"testing"

	"github.com/aereal/coll"
)

func TestOrderedMap_put_and_get(t *testing.T) {
	type entry struct {
		key   string
		value int
	}

	tests := []struct {
		name      string
		getKey    string
		entries   []entry
		wantVal   int
		wantFound bool
	}{
		{
			name: "existing key",
			entries: []entry{
				{"apple", 1},
				{"banana", 2},
			},
			getKey:    "banana",
			wantVal:   2,
			wantFound: true,
		},
		{
			name: "non-existing key",
			entries: []entry{
				{"a", 1},
			},
			getKey:    "z",
			wantVal:   0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := coll.NewOrderedMap[string, int]()
			for _, e := range tt.entries {
				m.Put(e.key, e.value)
			}
			got, ok := m.Get(tt.getKey)
			if ok != tt.wantFound {
				t.Errorf("Get(%q): found = %v, want %v", tt.getKey, ok, tt.wantFound)
			}
			if got != tt.wantVal {
				t.Errorf("Get(%q): got = %v, want %v", tt.getKey, got, tt.wantVal)
			}
		})
	}
}

func TestOrderedMap_keys_and_values(t *testing.T) {
	m := coll.NewOrderedMap[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	var gotKeys []string
	keysIter := m.Keys()
	keysIter(func(k string) bool {
		gotKeys = append(gotKeys, k)
		return true
	})

	wantKeys := []string{"a", "b", "c"}
	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("Keys: got %v, want %v", gotKeys, wantKeys)
	}

	var gotVals []int
	valsIter := m.Values()
	valsIter(func(v int) bool {
		gotVals = append(gotVals, v)
		return true
	})

	wantVals := []int{1, 2, 3}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("Values: got %v, want %v", gotVals, wantVals)
	}
}

func TestOrderedMap_All(t *testing.T) {
	m := coll.NewOrderedMap[string, int]()
	m.Put("x", 10)
	m.Put("y", 20)

	var gotKeys []string
	var gotVals []int
	allIter := m.All()
	allIter(func(k string, v int) bool {
		gotKeys = append(gotKeys, k)
		gotVals = append(gotVals, v)
		return true
	})

	wantKeys := []string{"x", "y"}
	wantVals := []int{10, 20}

	if !reflect.DeepEqual(gotKeys, wantKeys) {
		t.Errorf("All (keys): got %v, want %v", gotKeys, wantKeys)
	}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("All (values): got %v, want %v", gotVals, wantVals)
	}
}
