package group_test

import (
	"fmt"
	"sort"

	"github.com/lemmi/group"
)

func ExampleGrouper_Scan() {
	s := sort.IntSlice([]int{2, 2, 1, 2, 1, 1})
	s.Sort()

	var groups [][]int
	var g group.Grouper

	for g.Scan(s) {
		group := s[g.L:g.R]
		groups = append(groups, group)
	}

	fmt.Println(groups)
	// Output:
	// [[1 1 1] [2 2 2]]
}
