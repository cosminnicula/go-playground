package main

import (
	"testing"
)

func TestFunctionClosure(t *testing.T) {
	adder := func() func(int) int {
		sum := 0                 // closure
		return func(x int) int { // this anonymous function is bound to the enclosing function scope (e.g. can access sum)
			sum += x
			return sum
		}
	}

	someAdder := adder()
	sum := 0
	for i := 0; i < 10; i++ {
		sum = someAdder(i)
	}

	actual := sum
	want := 45

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	urls := []string{"https://google.com", "https://youtube.com"}

	c := make(chan string)
	for _, u := range urls {
		go func(url string) {
			c <- url + ": done"
		}(u) // closure (equivalent to closure in Java or JavaScript)
		// bounding the url to the scope of this function prevents the function accessing a value other than the one that was explicitly used when was called
	}

	newUrls := make([]string, len(urls))
	for i := 0; i < len(urls); i++ {
		newUrls[i] = <-c
	}

	actual2 := newUrls[0]
	want2 := "https://google.com: done"
	want3 := "https://youtube.com: done"

	if actual2 != want2 && actual2 != want3 {
		t.Errorf("actual %v expected %v or %v", actual2, want2, want3)
	}
}
