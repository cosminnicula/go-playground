package main

// an interface type is defined as a set of method signatures
// a value of interface type can hold any value that implements those methods
// a type implements an interface by implementing its methods. there is no explicit declaration of intent, no "implements" keyword
// under the hood, interface values can be thought of as a tuple of a value and a concrete type: fmt.Printf("(%v, %T)\n", i, i) will output the tuple (value, type)
import (
	"testing"
)

type Vehicle interface {
	start() string
}

type Something string

func (s Something) start() string {
	return "something started"
}

type Car struct {
	model string
}

func (c *Car) start() string {
	return "car started"
}

func TestInterfaces(t *testing.T) {
	var v Vehicle
	something := Something("something")
	car := Car{model: "tesla"}
	v = something

	actual := v.start()
	want := "something started"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	v = &car

	actual = v.start()
	want = "car started"

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	// v is a Car (not *Car), which means it doesn't implement the Vehicle interface => "v = car" won't compile
	// v = car
}
