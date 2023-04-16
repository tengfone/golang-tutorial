package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func httpPackage() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp) // This only prints the header not the body. Need to use the reader to read the body

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) // Esentially the respon bondy will take the byte slice and fill it with the body of the response
	// fmt.Println(string(bs))

	// This line is the same as the 2 lines above.
	// this is a writer interface that take info and send it to somewhere
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

type logWriter struct{}

// This is a custom writer that implements the writer interface, notice its called Write and it takes in a byte slice and returns
// an int and an error (check official docs).
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
