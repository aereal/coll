package coll_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/aereal/coll"
)

func TestSet(t *testing.T) {
	strSet := coll.NewSet("c", "a", "b")
	if gotLen := strSet.Len(); gotLen != 3 {
		t.Errorf("Len() returns unexpected value: %d", gotLen)
	}
	gotStrs := slices.Sorted(strSet.Values())
	wantStrs := []string{"c", "a", "b"}
	if !reflect.DeepEqual(gotStrs, []string{"a", "b", "c"}) {
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

func TestSet_Diff(t *testing.T) {
	testCases := []struct {
		lhs  *coll.Set[string]
		rhs  *coll.Set[string]
		want *coll.Set[string]
		name string
	}{
		{
			name: "empty vs empty",
			lhs:  coll.NewSet[string](),
			rhs:  coll.NewSet[string](),
			want: coll.NewSet[string](),
		},
		{
			name: "lhs == rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet[string](),
		},
		{
			name: "lhs > rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("a"),
			want: coll.NewSet("b", "c"),
		},
		{
			name: "lhs < rhs",
			lhs:  coll.NewSet("a"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet("b", "c"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Diff(tc.rhs)
			gotSlice := slices.Sorted(got.Values())
			wantSlice := slices.Sorted(tc.want.Values())
			if !reflect.DeepEqual(wantSlice, gotSlice) {
				t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", wantSlice, gotSlice)
			}
		})
	}
}

func TestSet_Intersect(t *testing.T) {
	testCases := []struct {
		lhs  *coll.Set[string]
		rhs  *coll.Set[string]
		want *coll.Set[string]
		name string
	}{
		{
			name: "empty vs empty",
			lhs:  coll.NewSet[string](),
			rhs:  coll.NewSet[string](),
			want: coll.NewSet[string](),
		},
		{
			name: "lhs == rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet("a", "b", "c"),
		},
		{
			name: "lhs > rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("a"),
			want: coll.NewSet("a"),
		},
		{
			name: "lhs < rhs",
			lhs:  coll.NewSet("a"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet("a"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Intersect(tc.rhs)
			gotSlice := slices.Sorted(got.Values())
			wantSlice := slices.Sorted(tc.want.Values())
			if !reflect.DeepEqual(wantSlice, gotSlice) {
				t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", wantSlice, gotSlice)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	testCases := []struct {
		lhs  *coll.Set[string]
		rhs  *coll.Set[string]
		want *coll.Set[string]
		name string
	}{
		{
			name: "empty vs empty",
			lhs:  coll.NewSet[string](),
			rhs:  coll.NewSet[string](),
			want: coll.NewSet[string](),
		},
		{
			name: "lhs == rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet("a", "b", "c"),
		},
		{
			name: "lhs > rhs",
			lhs:  coll.NewSet("a", "b", "c"),
			rhs:  coll.NewSet("d"),
			want: coll.NewSet("a", "b", "c", "d"),
		},
		{
			name: "lhs < rhs",
			lhs:  coll.NewSet("d"),
			rhs:  coll.NewSet("a", "b", "c"),
			want: coll.NewSet("d", "a", "b", "c"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Union(tc.rhs)
			gotSlice := slices.Sorted(got.Values())
			wantSlice := slices.Sorted(tc.want.Values())
			if !reflect.DeepEqual(wantSlice, gotSlice) {
				t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", wantSlice, gotSlice)
			}
		})
	}
}

func TestSet_empty(t *testing.T) {
	nums := new(coll.Set[int])
	if got := nums.Contains(42); got {
		t.Errorf("Contains(42) reports true unexpectedly")
	}
}
