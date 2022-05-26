package main

import (
	"testing"
)

func TestPointers(t *testing.T) {
	i := 1

	changeI := func(pi *int) {
		*pi = 2
	}

	changeI(&i)

	actual := i
	want := 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	dontChangeI := func(vi int) {
		vi = 1
	}

	dontChangeI(i)

	actual = i
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
