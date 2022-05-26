package main

import (
	"testing"
)

type MyInterface interface {
	myFunction() string
}

type MyStruct struct {
	myString string
}

func (m *MyStruct) myFunction() string {
	if m == nil { // in Go it is common to write methods that gracefully handle being called with a nil receiver
		return "nil"
	}

	return m.myString
}

// if the concrete value inside the interface itself is nil, the method will be called with a nil receiver
func TestInterfaceValuesWithNilUnderlyingValues(t *testing.T) {
	var i MyInterface

	var m *MyStruct
	i = m

	actual := i.myFunction() // does not trigger a null pointer error
	want := "nil"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
