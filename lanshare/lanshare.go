package lanshare

import (
	"fmt"
	"io"
	"net"
)

type ClientType int32

const (
	Sender ClientType = iota
	Receiver
)

type Client struct {
	Type ClientType
}

func (c *Client) Run() {
	// client set config
	c.serve()
}

func (c *Client) serve() {
	var conn net.Conn
	if c.Type == Sender {
		serverConn, err := net.Listen("tcp", "127.0.0.1:12345")
		if err != nil {
			panic(err)
		}
		conn, err = serverConn.Accept()
		if err != nil {
			panic(err)
		}
	} else if c.Type == Receiver {
		var err error
		conn, err = net.Dial("tcp", "127.0.0.1:12345")
		if err != nil {
			panic(err)
		}
	}

	// TODO: for test
	if c.Type == Receiver {
		conn.Write([]byte("client...ping..."))
	}

	// 轮训处理
	for {
		// TODO: 使用tcp read（message格式）
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		c.processMessage(buf)
		io.WriteString(conn, "hello")
	}
}

func (c *Client) processMessage(data []byte) {
	fmt.Printf("type: %d, %s\n", c.Type, string(data))
}
