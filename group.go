// Package group provides grouping a slice via a less function
package group

import (
	"reflect"
)

// Interface is a subset of ``sort.Interface''. Any type that can be sorted by
// the sort package can be grouped.
// A group is a range of elements that compare equal according to the
// collection's Less method.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with index i should sort before the
	// element with index j.
	Less(i, j int) bool
}

// A Grouper holds the lower and upper index of a sorted collection where
// all elements in that range compare equal according to the Less() Method of
// the collection
type Grouper struct {
	// L is the lower index of the element in the current group
	L int
	// R is the upper index of the element in the current group
	R int
}

// Scan checks Data for the next group. It will return false when no items are
// left in the collection
func (g *Grouper) Scan(data Interface) bool {
	if g.R == data.Len() {
		return false
	}
	g.L = g.R
	for {
		g.R++
		if g.R == data.Len() {
			break
		}
		if data.Less(g.L, g.R) {
			break
		}
	}

	return true
}

// Slice groups any slice []T with the supplied less function and returns the
// grouped result [][]T
func Slice(slice interface{}, less func(i, j int) bool) interface{} {
	ret := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice)), 0, 0)
	sv := reflect.ValueOf(slice)
	sl := sv.Len()
	var l, k int
	for k = 1; k < sl; k++ {
		if less(l, k) {
			group := sv.Slice3(l, k, k)
			ret = reflect.Append(ret, group)
			l = k
		}
	}

	group := sv.Slice3(l, k, k)
	ret = reflect.Append(ret, group)

	return ret.Interface()
}
