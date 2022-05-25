package main

import (
	"fmt"
	"testing"
	"time"
)

var c1 = make(chan int, 2)
var c2 = make(chan int, 2)

func sendToChannel1() {
	for i := 0; i < cap(c1); i++ {
		time.Sleep(500 * time.Millisecond)
		c1 <- i
	}
	close(c1)
}

func sendToChannel2() {
	for i := 0; i < cap(c2); i++ {
		time.Sleep(500 * time.Millisecond)
		c2 <- i
	}
	close(c2)
}

func TestSelect(t *testing.T) {
	go sendToChannel1()
	go sendToChannel2()

	select {
	case c1Value := <-c1:
		fmt.Printf("received %v from channel 1\n", c1Value)
	case c2Value := <-c2:
		fmt.Printf("received %v from channel 2\n", c2Value)
	}
}
