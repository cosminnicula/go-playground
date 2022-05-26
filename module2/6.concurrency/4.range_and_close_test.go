package main

import (
	"testing"
	"time"
)

func sendToChannel(n int, c chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(500 * time.Millisecond)
		c <- i
	}
	close(c)
}

// only the sender should close the channel, and not the receiver
func TestRangeAndClose(t *testing.T) {
	c := make(chan int, 2)

	go sendToChannel(cap(c), c)

	j := 0
	for i := range c {
		time.Sleep(510 * time.Millisecond)

		actual := i
		want := j
		j++

		if actual != want {
			t.Errorf("[iteration %v] actual %v expected %v", i, actual, want)
		}
	}

	_, actual2 := <-c
	want2 := false

	if actual2 != want2 {
		t.Errorf("actual %v expected %v", actual2, want2)
	}
}
