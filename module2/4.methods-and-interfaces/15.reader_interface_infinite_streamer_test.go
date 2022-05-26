package main

import (
	"fmt"
	"io"
	"testing"

	"golang.org/x/tour/reader"
)

type Reader struct {
	content string
}

func (r Reader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte(r.content[i])
	}

	return len(b), nil
}

// this test doesn't assert anything, since is an infinite stream
func TestReaderInterfaceInfiniteStreamer(t *testing.T) {
	r := Reader{content: "A...etc., etc."}
	b := make([]byte, 1)

	iteration := 0
	for {
		n, err := r.Read(b)
		fmt.Printf("(iteration %v) bytes read: %v, err: %v, raw bytes: %v, bytes: %q\n", iteration, n, err, b, b[:n])
		iteration++
		if err == io.EOF {
			break
		}
	}

	reader.Validate(Reader{})

}
