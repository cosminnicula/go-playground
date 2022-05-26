package main

import (
	"fmt"
	"testing"
)

// the error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//     Error() string
// }

type MyError struct {
	Level string
	What  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Level, e.What)
}

func run() error {
	return &MyError{Level: "error", What: "it didn't work"}
}

func TestErrorInterface(t *testing.T) {
	var actual interface{}

	if err := run(); err != nil {
		// As with fmt.Stringer, the fmt package looks for the error interface when printing values
		actual = fmt.Sprintf("%v", err)
	}

	want := "[error] it didn't work"

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
