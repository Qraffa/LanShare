package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	io.WriteString(conn, "client")
}
