package main

import (
	"testing"
)

// the interface type that specifies zero methods is known as the empty interface
// empty interfaces are used by code that handles values of unknown type
func TestFunction(t *testing.T) {
	var i interface{}

	actual := i

	if actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}

	i = 1

	actual = i
	want := 1

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	i = "hi"

	actual = i
	want2 := "hi"

	if actual != want2 {
		t.Errorf("actual %v expected %v", actual, want2)
	}
}
