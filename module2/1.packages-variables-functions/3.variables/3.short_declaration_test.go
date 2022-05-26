package main

import (
	"testing"
)

// outside a function, every statement begins with var, func etc., so := is not available
// e.g. someVar := "something" won't work here

func TestShortDeclaration(t *testing.T) {
	var i, j int = 1, 2
	x, y := 3, "four" // inside a function, := can be used in place of a var declaration with implicit type

	if i != 1 {
		t.Errorf("actual %q expected %q", i, 1)
	}

	if j != 2 {
		t.Errorf("actual %q expected %q", j, 2)
	}

	if x != 3 {
		t.Errorf("actual %q expected %q", x, 3)
	}

	if y != "four" {
		t.Errorf("actual %q expected %q", y, "four")
	}
}
