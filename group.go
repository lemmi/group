// Package group provides grouping a slice via a less function
package group

import (
	"reflect"
)

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
}

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
