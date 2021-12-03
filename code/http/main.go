package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.de")
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %v\n", resp)

	fmt.Println("response body:")
	io.Copy(os.Stdout, resp.Body)
}
