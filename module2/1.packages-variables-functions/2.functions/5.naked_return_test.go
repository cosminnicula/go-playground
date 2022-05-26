package main

import (
	"strings"
	"testing"
)

func TestNakedReturn(t *testing.T) {
	upperCase := func(s string) (upperCaseString string, stringLength int) {
		upperCaseString = strings.ToUpper(s)
		stringLength = len(s)

		return // naked return
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
