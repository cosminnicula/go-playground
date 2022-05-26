package main

import (
	"strings"
	"testing"
)

func TestFunction(t *testing.T) {
	upperCase := func(s string) string {
		return strings.ToUpper(s)
	}

	actual := upperCase("hello")
	want := "HELLO"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
