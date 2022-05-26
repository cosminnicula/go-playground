package main

import (
	"testing"
)

type MyStruct struct {
	x int
	y int
}

func TestStructFields(t *testing.T) {

	m := MyStruct{x: 1, y: 1}

	actual := m.x
	want := 1

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = m.y
	want = 1

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
