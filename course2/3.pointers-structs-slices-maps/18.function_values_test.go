package main

import (
	"testing"
)

func TestFunctionValues(t *testing.T) {
	// function returning function
	getArea := func() func(int, int) int {
		return func(x, y int) int {
			return x * y
		}
	}

	// function receiving function as parameter
	squareSomeArea := func(fn func(int, int) int) int {
		area := fn(2, 3)
		return area * area
	}

	actual := getArea()(2, 3) // equivalent to currying in JavaScript; similar with currying in Java
	want := 6

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = squareSomeArea(getArea())
	want = 36

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
