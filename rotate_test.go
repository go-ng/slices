package slices

import (
	"testing"
)

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	r := []int{1, 2, 3, 4}
	Rotate(r, 1)
	if !intsEqual([]int{4, 1, 2, 3}, r) {
		t.Fatalf("%v", r)
	}
	Rotate(r, -1)
	if !intsEqual([]int{1, 2, 3, 4}, r) {
		t.Fatalf("%v", r)
	}
}
