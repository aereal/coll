package coll

import "iter"

type SetLike[E comparable] interface {
	Len() int
	Values() iter.Seq[E]
	Contains(E) bool
}

type almostSet[E comparable] interface {
	SetLike[E]
	unsafeAppend(E)
}

// Diff returns a new [Set] containing elements that are in xs or ys but not in both.
func Diff[E comparable](xs, ys SetLike[E]) SetLike[E] {
	ret := NewSet[E]()
	buildDiff(ret, xs, ys)
	return ret
}

func buildDiff[E comparable](ret almostSet[E], xs, ys SetLike[E]) {
	lhs := xs
	rhs := ys
	if lhs.Len() < rhs.Len() {
		lhs = ys
		rhs = xs
	}
	for lv := range lhs.Values() {
		if !rhs.Contains(lv) {
			ret.unsafeAppend(lv)
		}
	}
}

// Intersect returns a new [Set] containing elements that are present in both xs and ys.
func Intersect[E comparable](xs, ys SetLike[E]) SetLike[E] {
	ret := NewSet[E]()
	buildIntersection(ret, xs, ys)
	return ret
}

func buildIntersection[E comparable](ret almostSet[E], xs, ys SetLike[E]) {
	lhs := xs
	rhs := ys
	if lhs.Len() < rhs.Len() {
		lhs = ys
		rhs = xs
	}
	for lv := range lhs.Values() {
		if rhs.Contains(lv) {
			ret.unsafeAppend(lv)
		}
	}
}

// Union returns a new [Set] containing all elements from both xs and ys.
func Union[E comparable](xs, ys SetLike[E]) SetLike[E] {
	ret := NewSet[E]()
	buildUnion(ret, xs, ys)
	return ret
}

func buildUnion[E comparable](ret almostSet[E], xs, ys SetLike[E]) {
	for v := range xs.Values() {
		ret.unsafeAppend(v)
	}
	for v := range ys.Values() {
		ret.unsafeAppend(v)
	}
}
