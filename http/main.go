package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// bs := make([]byte, 99999)
	// _, err = resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(bs))

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
