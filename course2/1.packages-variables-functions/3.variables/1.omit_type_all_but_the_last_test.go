package main

import (
	"testing"
)

var packgeLevelVar1, packgeLevelVar2 string

func TestOmitTypeAllButTheLast(t *testing.T) {
	packgeLevelVar1 = "packgeLevelVar1"

	var functionLevelVar1, functionLevelVar2 string
	functionLevelVar1 = "functionLevelVar1"
	functionLevelVar2 = "functionLevelVar2"

	actual := packgeLevelVar1
	want := "packgeLevelVar1"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = functionLevelVar1
	want = "functionLevelVar1"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	actual = functionLevelVar2
	want = "functionLevelVar2"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
