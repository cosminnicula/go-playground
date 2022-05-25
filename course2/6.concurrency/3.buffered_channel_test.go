package main

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	c := make(chan int, 2)

	c <- 1
	c <- 2

	actual := <-c
	want := 1

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}

	c <- 3
	select {
	case c <- 4: // Put 4 in the channel unless it is full
	default:
		fmt.Println("channel is full. discarding value")
	}

	actual = <-c
	want = 2

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
