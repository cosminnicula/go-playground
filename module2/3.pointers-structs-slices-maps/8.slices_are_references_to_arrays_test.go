package main

// a slice does not store any data, it just describes a section of an underlying array
// changing the elements of a slice modifies the corresponding elements of its underlying array

import (
	"testing"
)

// arrays cannot be resized
func TestSlicesAreReferencesToArrays(t *testing.T) {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	b[0] = "X"

	actual := a[1]
	want := "X"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
