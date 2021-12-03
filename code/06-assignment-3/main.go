package main

import (
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	filePointer := file

	io.Copy(os.Stdout, filePointer)
}
