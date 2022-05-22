package main

import (
	"testing"
)

// arrays cannot be resized
func TestArrays(t *testing.T) {
	var myArray [2]int
	myArray[0] = 1
	myArray[1] = 2

	actual := myArray[0]
	want := 1

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
