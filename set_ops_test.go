package coll_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/aereal/coll"
)

func Test_set_ops(t *testing.T) {
	t.Run("Diff()", func(t *testing.T) {
		xs := coll.NewSet(1, 2, 3, 4, 5)
		ys := coll.NewOrderedSet(1, 2, 3)
		got := slices.Sorted(coll.Diff(xs, ys).Values())
		want := []int{4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", want, got)
		}
	})
	t.Run("Intersect()", func(t *testing.T) {
		xs := coll.NewSet(1, 2, 3, 4, 5)
		ys := coll.NewOrderedSet(1, 2, 3)
		got := slices.Sorted(coll.Intersect(xs, ys).Values())
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", want, got)
		}
	})
	t.Run("Union()", func(t *testing.T) {
		xs := coll.NewSet(1, 2, 3)
		ys := coll.NewOrderedSet(4, 5)
		got := slices.Sorted(coll.Union(xs, ys).Values())
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("mismatch:\n\twant: %#v\n\t got: %#v", want, got)
		}
	})
}
