package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// fmt.Println(<-c) // blocking call; similar with JavaScript await, or Java CompletableFuture.get()
	// fmt.Println(<-c) // blocking call
	// fmt.Println(<-c) // blocking call

	// for {
	// 	go checkLink(<-c, c)
	// }

	//or

	for l := range c {
		// time.Sleep(5 * time.Second) // it means that we can execute a go routine every five seconds;
		// // during the waiting, messages send throguh the channel are not lost, but queued
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	go checkLink(l, c) // DON'T reference the same variable inside different goroutines. Here, l is shared.
		// }()
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkLink(url, c)
		}(l) // share information with the child routine by passing it as an argument via an anonymous function (see current example),
		// or communicate with the child routine over channels (see <-c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url) // blocking call
	if err != nil {
		fmt.Println(url, " might be down")
		c <- url
		return
	}

	fmt.Println(url, " is up")
	c <- url
}
