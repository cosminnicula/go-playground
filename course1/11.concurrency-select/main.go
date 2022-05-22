package main

import (
	"fmt"
	"time"
)

func main() {
	waitOnTwoChannels()

	// waitOnTwoChannelsDeadlock()
}

// https://gobyexample.com/select
func waitOnTwoChannels() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("[", time.Now().UTC(), "][emit]one")
		c1 <- "one"

		time.Sleep(2 * time.Second)
		fmt.Println("[", time.Now().UTC(), "][emit]one")
		c1 <- "one"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("[", time.Now().UTC(), "][emit]two")
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("[", time.Now().UTC(), "][start]iteration ", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
		fmt.Println("[", time.Now().UTC(), "][end]iteration ", i)
	}

	// possible outcomes:
	// iteration 0 received one, iteration 1 received one
	// iteration 0 received one, iteration 2 received two
}

func waitOnTwoChannelsDeadlock() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for m1 := range c1 {
		fmt.Println(m1)
		select { //deadlock
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
