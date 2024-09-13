package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, err := os.Open("test/src")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("test/dst")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}
