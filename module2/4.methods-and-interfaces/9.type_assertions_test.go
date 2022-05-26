package main

import (
	"testing"
)

type MySecondString string

type MySecondStruct struct {
	myValue string
}

// a type assertion provides access to an interface value's underlying concrete value
func TestTypeAssertions(t *testing.T) {
	var i interface{} = MySecondStruct{myValue: "myValue"}

	// i is MySecondStruct
	mySecondStruct, ok := i.(MySecondStruct) // similar to reading from a map

	actual := mySecondStruct.myValue
	want := "myValue"

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	if ok != true {
		t.Errorf("actual %v expected %v", ok, true)
	}

	// i is MySecondString
	i = MySecondString("mySecondString")

	mySecondString, ok2 := i.(MySecondString)

	actual2 := string(mySecondString)
	want2 := "mySecondString"

	if actual2 != want2 {
		t.Errorf("actual %v expected %v", actual2, want2)
	}

	if ok2 != true {
		t.Errorf("actual %v expected %v", ok, true)
	}

	// i is a string
	i = "hi"

	aString := i.(string)

	actual3 := aString
	want3 := "hi"

	if actual3 != want3 {
		t.Errorf("actual %v expected %v", actual3, want3)
	}

	// something := i.(float64)
	// however, this will not panic: something, ok := i.(float64)
}
