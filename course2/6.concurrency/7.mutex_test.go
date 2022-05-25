package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	var m sync.Mutex
	fmt.Printf("Number of CPU cores: %v\n", runtime.NumCPU())
	fmt.Printf("Max parallelism: %v\n", runtime.GOMAXPROCS(0))
	start := time.Now()
	for i := 0; i < 10000000; i++ { //10 million threads
		wg.Add(1)
		go increment(&wg, &m)
	}
	wg.Wait()
	fmt.Printf("10 million threads concurrently incrementing a value took %s\n", time.Since(start))

	actual := x
	want := 10000000

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
