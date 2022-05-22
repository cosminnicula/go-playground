package main

import (
	"strings"
	"testing"
)

func Function(t *testing.T) {
	upperCase := func(s string) string {
		return strings.ToUpper(s)
	}

	actual := upperCase("hello")
	want := "HELLO"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}

func OmmitTypeAllButTheLast(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	actual := add(1, 2)
	want := 3

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
