package main

import (
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]

	filePointer, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, filePointer)
}
