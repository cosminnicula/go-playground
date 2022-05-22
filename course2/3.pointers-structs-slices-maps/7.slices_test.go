package main

import (
	"reflect"
	"testing"
)

// a slice is a dynamically-sized, flexible view into the elements of an array
// a type []T is a slice
func TestSlices(t *testing.T) {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	mySlice := myArray[0:2]

	actual := mySlice
	want := []int{1, 2}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
