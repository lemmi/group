package group

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func intLess(s []int) func(int, int) bool {
	return func(i, j int) bool { return s[i] < s[j] }
}

func groupAndCompare(t *testing.T, slice interface{}, want interface{}, less func(i, j int) bool) {
	t.Helper()
	sort.Slice(slice, less)
	groups := Slice(slice, less).([][]int)
	if !reflect.DeepEqual(groups, want) {
		t.Fatalf("\nWanted: %v\n   Got: %v\n Slice: %v", want, groups, slice)
	}
}

func TestFixedIntSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 1, 2, 1, 3, 3, 2, 1, 1, 4, 4, 4, 5, 3}
	want := [][]int{
		[]int{1, 1, 1, 1, 1},
		[]int{2, 2, 2},
		[]int{3, 3, 3, 3},
		[]int{4, 4, 4, 4},
		[]int{5},
	}
	groupAndCompare(t, s, want, intLess(s))
}

func intSlice(value int, n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = value
	}

	return a
}

type IntSliceTest int

func (lt IntSliceTest) Test(t *testing.T) {
	var want [][]int
	l := int(lt)

	var n int
	for i := 0; n < l; i++ {
		max := l - n
		r := rand.Intn(max/2+1) + 1
		if r > max {
			r = max
		}
		want = append(want, intSlice(i, r))
		n += r
	}

	var slice []int
	for _, s := range want {
		slice = append(slice, s...)
	}
	t.Log(slice, want)
	groupAndCompare(t, slice, want, intLess(slice))
}

func TestIntSlices(t *testing.T) {
	for l := 1; l <= 1024; l *= 2 {
		t.Run(fmt.Sprintf("TestIntSliceL%d", l), IntSliceTest(l).Test)
	}
}

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
