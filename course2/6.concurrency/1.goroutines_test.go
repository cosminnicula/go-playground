package main

import (
	"testing"
	"time"
)

var myArray []int

func insertInArray() {
	i := 0
	time.Sleep(500 * time.Millisecond)
	myArray = append(myArray, i)
}

func TestGoroutines(t *testing.T) {
	go insertInArray()

	actual := len(myArray)
	want := 0

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	time.Sleep(510 * time.Millisecond)

	actual = len(myArray)
	want = 1

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
