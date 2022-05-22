package main

// The expression T(v) converts the value v to the type T.
// Assignment between items of different type requires an explicit conversion

import (
	"math"
	"testing"
)

func TestExplicitConversion(t *testing.T) {
	var i = 25
	var float float64 = float64(i) // "var float float64 = i" won't work

	actual := float
	var want float64 = 25

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	actual = math.Sqrt(float64(i))
	want = 5

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
