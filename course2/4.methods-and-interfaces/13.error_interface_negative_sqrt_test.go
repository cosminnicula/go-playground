package main

import (
	"fmt"
	"math"
	"testing"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("square root of %f is impossible", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func TestErrorInterfaceNegativeSqrt(t *testing.T) {
	sqrt, _ := Sqrt(25)

	actual := sqrt
	want := float64(5)

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	_, err2 := Sqrt(-1)

	actual2 := fmt.Sprintf("%v", err2)
	want2 := "square root of -1.000000 is impossible"

	if actual2 != want2 {
		t.Errorf("actual %v expected %v", actual2, want2)
	}

	actual2 = fmt.Sprintf("%f", err2)
	want2 = "-1.000000"

	if actual2 != want2 {
		t.Errorf("actual %v expected %v", actual2, want2)
	}
}
