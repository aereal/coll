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
}

func TestOrderedSet_empty(t *testing.T) {
	nums := new(coll.OrderedSet[int])
	if got := nums.Contains(42); got {
		t.Errorf("Contains(42) reports true unexpectedly")
	}
}
