package main

import (
	"testing"
)

// obviously, constants cannot be reassigned
const a = 123     // this is an untyped constant (type is inferred from its value)
const b int = 456 // this is a typed constant (type is explicitly declared)

// multi-line constant declaration
const (
	c = 789
	d = 987
)

func TestConstants(t *testing.T) {
	actual := a
	want := 123

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
