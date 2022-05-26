package main

import (
	"testing"
)

func getType(i interface{}) interface{} {
	var thisValue interface{}

	switch v := i.(type) { // the declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type
	case string:
		thisValue = "string"
	case int:
		thisValue = 123
	default:
		thisValue = v
	}

	return thisValue
}

// a type switch is a construct that permits several type assertions in series
func TestTypeSwitches(t *testing.T) {
	// i is a string
	var i interface{} = "something"

	var actual interface{} = getType(i)
	var want interface{} = "string"

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	// i is an int
	i = 123

	actual = getType(i)
	want = 123

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	// i is a boolean
	i = false

	actual = getType(i)
	want = false

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
