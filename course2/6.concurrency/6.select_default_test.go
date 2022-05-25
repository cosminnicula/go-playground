package main

import (
	"fmt"
	"testing"
	"time"
)

func process(c chan int) {
	time.Sleep(500 * time.Millisecond)
	c <- 0
}

func TestSelectDefault(t *testing.T) {
	c := make(chan int, 2)
	go process(c)

	var actual int
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case actual := <-c:
			fmt.Println("received value: ", actual)
			return
		default:
			fmt.Println("no value received")
		}
	}

	want := 0

	if actual != want {
		t.Errorf("actual %q expected %q", actual, want)
	}
}
