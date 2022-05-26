package main

// make function allocates a zeroed array and returns a slice that refers to that array:

import (
	"testing"
)

func TestSlicesWithMake(t *testing.T) {
	// slice initialization (length 6)
	s := make([]int, 6)

	actual := len(s)
	want := 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s)
	want = 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// slice initialization (length 6)
	s2 := make([]int, 0, 6)

	actual = len(s2)
	want = 0

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s2)
	want = 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// shrink the length of the slice
	s3 := s[:2]

	actual = len(s3)
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s3)
	want = 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// drop its first two values
	s4 := s[2:]

	actual = len(s4)
	want = 4

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s4)
	want = 4

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
