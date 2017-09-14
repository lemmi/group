package group

import (
	"reflect"
	"sort"
	"testing"
)

func TestGrouper(t *testing.T) {
	s := []int{1, 2, 3, 4, 1, 2, 1, 3, 3, 2, 1, 1, 4, 4, 4, 5, 3}
	want := [][]int{
		[]int{1, 1, 1, 1, 1},
		[]int{2, 2, 2},
		[]int{3, 3, 3, 3},
		[]int{4, 4, 4, 4},
		[]int{5},
	}

	is := sort.IntSlice(s)
	is.Sort()
	var groups [][]int
	var g Grouper
	for g.Scan(is) {
		group := s[g.L:g.R]
		groups = append(groups, group)
	}
	if !reflect.DeepEqual(groups, want) {
		t.Fatalf("\nWanted: %v\n   Got: %v\n Slice: %v", want, groups, s)
	}
}
