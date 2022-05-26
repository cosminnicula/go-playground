package main

import (
	"testing"
)

type Numbers interface {
	int | float64
}

func isLess[T Numbers](a, b T) bool {
	return a < b
}

func TestTypeParameters(t *testing.T) {
	actual := isLess(1, 2)
	want := true

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	actual = isLess(1.23, 2.33)
	want = true

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
