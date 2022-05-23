package main

// 1. methods with pointer receivers can modify the value to which the receiver points
// 2. avoids copying the value on each method call (more efficient for large objects)
// note: in general, all methods on a given type should have either value or pointer receivers, but not a mixture of both
import (
	"testing"
)

type Circle struct {
	radius float64
}

func (c Circle) diameter() float64 {
	return float64(2 * c.radius)
}

// // // diameter() is here written as a regular function, but with the same functionality
// func diameter(c Circle) float64 {
// 	return float64(2 * c.radius)
// }

// if c wouldn't have been a pointer, then c's radius would have remained unchanged after calling this function
func (c *Circle) scale(f float64) {
	c.radius = c.radius * f
}

// // scale() is here written as a regular function, but with the same functionality
// func scale(c *Circle, f float64) {
// 	c.radius = c.radius * f
// }

func TestMethodsWithPointerReceivers(t *testing.T) {
	c := Circle{radius: 5}

	c.scale(2) // short notation for (&c).scale(2)
	// calling the regular function with no receiver: scale(&c, 2)

	actual := c.diameter()
	want := float64(20)

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
