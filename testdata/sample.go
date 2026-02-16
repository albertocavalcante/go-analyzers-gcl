// Package testdata contains sample code that should trigger all go-analyzers diagnostics.
// Used by CI to verify the golangci-lint plugin works end-to-end.
package testdata

import "sort"

// TriggerMakeCopy should be flagged by makecopy:
// make+copy can be simplified to slices.Clone.
func TriggerMakeCopy() []int {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

// TriggerSearchMigrate should be flagged by searchmigrate:
// sort.Search can be replaced with slices.BinarySearch.
func TriggerSearchMigrate() int {
	s := []int{1, 2, 3, 4, 5}
	return sort.Search(len(s), func(i int) bool { return s[i] >= 3 })
}

// TriggerClampCheck should be flagged by clampcheck:
// clamp pattern can be simplified to min(max(x, lo), hi).
func TriggerClampCheck() int {
	x := 50
	lo := 0
	hi := 100
	if x < lo {
		x = lo
	} else if x > hi {
		x = hi
	}
	return x
}
