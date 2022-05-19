package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(resp.Body)
	// fmt.Println(string(buf.String()))

	// io.Copy(os.Stdout, resp.Body)

	mw := myWriter{}
	io.Copy(mw, resp.Body)
}

func (myWriter) Write(p []byte) (int, error) {
	fmt.Println("---custom writer---")
	fmt.Println(string(p))
	fmt.Println("---custom writer---")

	return len(p), nil
}
