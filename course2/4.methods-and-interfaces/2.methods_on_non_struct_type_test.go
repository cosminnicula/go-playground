package main

import (
	"testing"
)

type MyInt int

func (m MyInt) square() int {
	return int(m * m)
}

func TestMethodsOnNonStructType(t *testing.T) {
	m := MyInt(5)

	actual := m.square()
	want := 25

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
