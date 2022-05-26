package main

import (
	"testing"
)

var var1, var2 int = 1, 2 // variable initializers, one per variable

func TestInitializers(t *testing.T) {
	var var3, var4 = "var3", 123 // type is omitted => variable infers the type of the initializer

	actual_1 := var1
	want_1 := 1

	if actual_1 != want_1 {
		t.Errorf("actual %q expected %q", actual_1, want_1)
	}

	actual_1 = var2
	want_1 = 2

	if actual_1 != want_1 {
		t.Errorf("actual %q expected %q", actual_1, want_1)
	}

	actual_2 := var3
	want_2 := "var3"

	if actual_2 != want_2 {
		t.Errorf("actual %q expected %q", actual_2, want_2)
	}

	actual_3 := var4
	want_3 := 123

	if actual_3 != want_3 {
		t.Errorf("actual %q expected %q", actual_3, want_3)
	}
}
