package main

import (
	"testing"
)

func TestPointerToStructs(t *testing.T) {

	m := MyStruct{x: 1, y: 1}

	changeM := func(pm *MyStruct) {
		pm.x = 2 // convention for (*pm).x
	}

	changeM(&m)

	actual := m.x
	want := 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	dontChangeM := func(vm MyStruct) {
		vm.x = 1
	}

	dontChangeM(m)

	actual = m.x
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
