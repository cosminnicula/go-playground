package main

// a method is a function with a special receiver argument
import (
	"testing"
)

type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

// // area() is here written as a regular function, but with the same functionality
// func (r Rectangle) area() int {
// 	return r.height * r.width
// }

func TestMethodsOnStructType(t *testing.T) {
	r := Rectangle{3, 4}

	actual := r.area()
	want := 12

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
