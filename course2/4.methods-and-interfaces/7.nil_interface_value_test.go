package main

import (
	"testing"
)

type MySecondInterface interface {
	myMethod()
}

func TestNilInterfaceValue(t *testing.T) {
	var i MySecondInterface

	actual := i

	// atempting to call i.myMethod() will trigger a runtime error

	if actual != nil {
		t.Errorf("actual %v expected %v", actual, nil)
	}
}
