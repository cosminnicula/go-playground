package main

import (
	"testing"
)

// built-in function for appending: func append(s []T, vs ...T) []T
// if the backing array of s is too small to fit all the given values a bigger array will be allocated. the returned slice will point to the newly allocated array.

func TestSliceAppend(t *testing.T) {
	var s []int

	s = append(s, 1, 2, 3, 4, 5)

	actual := len(s)
	want := 5

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s)
	want = 6 // ? if length is odd and len > 5, then capacity is len + 1; otherwise, is len

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
