package main

import (
	"testing"
)

func TestSlicesOfSlices(t *testing.T) {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"

	actual := board[0][0]
	want := "X"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
