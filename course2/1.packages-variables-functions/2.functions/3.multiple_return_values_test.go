package main

import (
	"strings"
	"testing"
)

func TestMultiple(t *testing.T) {
	upperCase := func(s string) (string, int) {
		return strings.ToUpper(s), len(s)
	}

	actual_s, actual_len := upperCase("hello")
	want_s, want_len := "HELLO", 5

	if actual_s != want_s {
		t.Errorf("actual %q expected %q", actual_s, want_s)
	}

	if actual_len != want_len {
		t.Errorf("actual %q expected %q", actual_len, want_len)
	}
}
