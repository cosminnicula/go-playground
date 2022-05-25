package main

import (
	"reflect"
	"testing"
)

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func TestGenericAnyType(t *testing.T) {
	actual1, actual2 := splitAnySlice([]int{0, 1, 2, 3, 4, 5})
	want1, want2 := []int{0, 1, 2}, []int{3, 4, 5}

	if !reflect.DeepEqual(actual1, want1) || !reflect.DeepEqual(actual2, want2) {
		t.Errorf("actual1 %v expected1 %v, actual2 %v, expected %v", actual1, want1, actual2, want2)
	}
}
