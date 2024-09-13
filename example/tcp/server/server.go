package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening......")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept conn err")
			continue
		}

		fmt.Println("accept conn")
		go func(c net.Conn) {
			_, err := io.WriteString(c, "server")
			if err != nil {
				panic(err)
			}

			go func(c net.Conn) {
				buf, err := io.ReadAll(c)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(buf))
			}(c)
		}(conn)
	}

}
