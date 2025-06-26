package coll_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/aereal/coll"
)

func TestOrderedSet(t *testing.T) {
	strSet := coll.NewOrderedSet("c", "a", "b")
	if gotLen := strSet.Len(); gotLen != 3 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	gotStrs := slices.Collect(strSet.Values())
	wantStrs := []string{"c", "a", "b"}
	if !reflect.DeepEqual(gotStrs, []string{"c", "a", "b"}) {
		t.Errorf("Values() returns the unexpected value:\n\twant: %#v\n\t got: %#v", wantStrs, gotStrs)
	}
	if !strSet.Contains("a") {
		t.Error("the set says it DOES NOT contain 'a'")
	}
	if strSet.Contains("z") {
		t.Error("the set says it DOES contain 'z'")
	}
	strSet.Append("z")
	if !strSet.Contains("z") {
		t.Error("the appended set says it DOES contain 'z'")
	}
	if gotLen := strSet.Len(); gotLen != 4 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	strSet.Append("a") // try to append existent element
	if gotLen := strSet.Len(); gotLen != 4 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	strSet.Remove("a")
	if gotLen := strSet.Len(); gotLen != 3 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	if strSet.Contains("a") {
		t.Error("the set says it DOES contain 'a'")
	}

	// try to remove the element that is not in the set
	strSet.Remove("a")
	if gotLen := strSet.Len(); gotLen != 3 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	if strSet.Contains("a") {
		t.Error("the set says it DOES contain 'a'")
	}
}

func TestOrderedSet_Diff(t *testing.T) {
	testCases := []struct {
		lhs  *coll.OrderedSet[string]
		rhs  *coll.OrderedSet[string]
		want *coll.OrderedSet[string]
		name string
	}{
		{
			name: "empty vs empty",
			lhs:  coll.NewOrderedSet[string](),
			rhs:  coll.NewOrderedSet[string](),
			want: coll.NewOrderedSet[string](),
		},
		{
			name: "lhs == rhs",
			lhs:  coll.NewOrderedSet("a", "b", "c"),
			rhs:  coll.NewOrderedSet("a", "b", "c"),
			want: coll.NewOrderedSet[string](),
		},
		{
			name: "lhs > rhs",
			lhs:  coll.NewOrderedSet("a", "b", "c"),
			rhs:  coll.NewOrderedSet("a"),
			want: coll.NewOrderedSet("b", "c"),
		},
		{
			name: "lhs < rhs",
			lhs:  coll.NewOrderedSet("a"),
			rhs:  coll.NewOrderedSet("a", "b", "c"),
			want: coll.NewOrderedSet("b", "c"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Diff(tc.rhs)
			gotSlice := slices.Collect(got.Values())
			wantSlice := slices.Collect(tc.want.Values())
			if !reflect.DeepEqual(wantSlice, gotSlice) {
				t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", wantSlice, gotSlice)
			}
		})
	}
}

func TestOrderedSet_Intersect(t *testing.T) {
	testCases := []struct {
		lhs  *coll.OrderedSet[string]
		rhs  *coll.OrderedSet[string]
		want *coll.OrderedSet[string]
		name string
	}{
		{
			name: "empty vs empty",
			lhs:  coll.NewOrderedSet[string](),
			rhs:  coll.NewOrderedSet[string](),
			want: coll.NewOrderedSet[string](),
		},
		{
			name: "lhs == rhs",
			lhs:  coll.NewOrderedSet("a", "b", "c"),
			rhs:  coll.NewOrderedSet("a", "b", "c"),
			want: coll.NewOrderedSet("a", "b", "c"),
		},
		{
			name: "lhs > rhs",
			lhs:  coll.NewOrderedSet("a", "b", "c"),
			rhs:  coll.NewOrderedSet("a"),
			want: coll.NewOrderedSet("a"),
		},
		{
			name: "lhs < rhs",
			lhs:  coll.NewOrderedSet("a"),
			rhs:  coll.NewOrderedSet("a", "b", "c"),
			want: coll.NewOrderedSet("a"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Intersect(tc.rhs)
			gotSlice := slices.Collect(got.Values())
			wantSlice := slices.Collect(tc.want.Values())
			if !reflect.DeepEqual(wantSlice, gotSlice) {
				t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", wantSlice, gotSlice)
			}
		})
	}
}

func TestOrderedSet_empty(t *testing.T) {
	nums := new(coll.OrderedSet[int])
	if got := nums.Contains(42); got {
		t.Errorf("Contains(42) reports true unexpectedly")
	}
}
