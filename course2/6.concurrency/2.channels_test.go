package main

import (
	"testing"
)

// ch <- v    // send v to channel ch
// v := <-ch  // receive from ch, and assign value to v

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestFunction(t *testing.T) {
	c := make(chan int)
	s := []int{1, 2, 3, 4}

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	partialSum1, partialSum2 := <-c, <-c

	actual := partialSum1 + partialSum2
	want := 10

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
