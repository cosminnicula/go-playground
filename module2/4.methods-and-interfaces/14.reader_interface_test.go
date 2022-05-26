package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// the io.Reader interface has a Read method: func (T) Read(b []byte) (n int, err error)
func TestReaderInterface(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")
	var result []string

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		result = append(result, fmt.Sprintf("n = %v err = %v b = %q", n, err, b[:n]))
		if err == io.EOF {
			break
		}
	}

	actual := result[0]
	want := "n = 8 err = <nil> b = \"Hello, R\""

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
