package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// the fmt package (and many others) look for this interface to print values
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestStringerInterface(t *testing.T) {
	person := Person{Name: "john", Age: 999}

	actual := fmt.Sprintf("%v", person)
	want := "john (999 years)"

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
