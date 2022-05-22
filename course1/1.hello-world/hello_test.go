package main

import "testing"

func TestHello(t *testing.T) {
	actual := hello()
	want := "Hello, world"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
