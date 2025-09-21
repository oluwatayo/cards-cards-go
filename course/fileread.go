package main

import (
	"fmt"
	"io"
	"os"
)

func readFile() {
	fileName := os.Args[1]
	content, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	io.Copy(os.Stdout, content)
}
