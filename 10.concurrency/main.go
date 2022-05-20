package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type crawledData struct {
	url       string
	content   string
	startTime time.Time
	endTime   time.Time
}

func main() {
	crawlingIntervalInSeconds := 5
	c := make(chan crawledData)

	urls := map[string]int{
		"https://google.com":        0,
		"https://facebook.com":      0,
		"https://stackoverflow.com": 0,
		"https://golang.org":        0,
		"https://amazon.com":        0,
	}

	fmt.Printf("Crawling %v every %v seconds\n", urls, crawlingIntervalInSeconds)

	// initial content crawling for all urls
	for k := range urls {
		// launch subroutine for each url
		go func(url string) {
			// crawl
			c <- getContent(url, urls[url])
		}(k)
	}

	// crawl all urls every 5 seconds
	for cm := range c {
		// increment counter for url
		urls[cm.url] = urls[cm.url] + 1

		fmt.Println("-----------------")
		fmt.Println("[", cm.endTime, "][", cm.url, "][", urls[cm.url], "] done")
		fmt.Println("url: ", cm.url)
		fmt.Println("content: ", cm.content)
		fmt.Println("startTime: ", cm.startTime)
		fmt.Println("endTime: ", cm.endTime)
		fmt.Println("-----------------")

		// launch subroutine for current url
		go func(url string, counter int) {
			fmt.Println("[", time.Now().UTC(), "][", url, "][", counter+1, "] waiting")

			// throttle crawling
			time.Sleep(time.Duration(crawlingIntervalInSeconds) * time.Second)

			fmt.Println("[", time.Now().UTC(), "][", url, "][", counter+1, "] resuming")

			// crawl
			c <- getContent(url, counter)
		}(cm.url, urls[cm.url])
	}
}

func getContent(url string, counter int) crawledData {
	startTime := time.Now().UTC()
	fmt.Println("[", startTime, "][", url, "][", counter+1, "] initiating actual HTTP request")

	response, err := http.Get(url)

	endTime := time.Now().UTC()
	fmt.Println("[", endTime, "][", url, "][", counter+1, "] getting actual HTTP response")

	if err != nil {

		return crawledData{
			url:       url,
			content:   err.Error(),
			startTime: startTime,
			endTime:   endTime,
		}
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return crawledData{
		url:       url,
		content:   fmt.Sprintf("[%v]...truncated", strings.Trim(buf.String(), "\n")[:100]),
		startTime: startTime,
		endTime:   endTime,
	}
}
