package main

import (
	"testing"
)

func TestOmitTypeAllButTheLast(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	actual := add(1, 2)
	want := 3

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
