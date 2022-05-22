package main

// length is the number of elements it contains
// capancity is the number of elements in the underlying array
import (
	"testing"
)

func TestSlicesLengthAndCapacity(t *testing.T) {
	// slice initialization
	s := []int{2, 3, 5, 7, 11, 13}

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

	// shrink the length of the slice
	s = s[:2]

	actual = len(s)
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s)
	want = 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// extend the length of the slice
	s = s[:4]

	actual = len(s)
	want = 4

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s)
	want = 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// drop its first two values
	s = s[2:]

	actual = len(s)
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = cap(s)
	want = 4

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
