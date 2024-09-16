package lanshare

import (
	"fmt"
	"io"
	"lanshare/message"
	"net"
	"time"
)

type ClientType int32

const (
	Sender ClientType = iota
	Receiver
)

type Client struct {
	Type ClientType
	Conn net.Conn
}

func (c *Client) Run() {
	// client set config
	c.serve()
}

func (c *Client) serve() {
	if c.Type == Sender {
		serverConn, err := net.Listen("tcp", "127.0.0.1:12345")
		if err != nil {
			panic(err)
		}
		c.Conn, err = serverConn.Accept()
		if err != nil {
			panic(err)
		}
	} else if c.Type == Receiver {
		var err error
		c.Conn, err = net.Dial("tcp", "127.0.0.1:12345")
		if err != nil {
			panic(err)
		}
	}

	// TODO: for test
	if c.Type == Receiver {
		pingMsg := &message.Message{
			Header: &message.MessageHeader{
				Type:   message.Ping,
				Length: 4,
			},
			PlayLoad: []byte("ping"),
		}
		if err := c.Send(pingMsg); err != nil {
			panic(err)
		}
	}

	// 轮训处理
	for {
		msg, err := c.Receive()
		if err != nil {
			panic(err)
		}
		c.processMessage(msg)
	}
}

func (c *Client) Receive() (msg *message.Message, err error) {
	header := make([]byte, message.HeaderLength)
	_, err = io.ReadFull(c.Conn, header)
	if err != nil {
		return
	}
	messageHeader := message.DecodeMessageHeader(header)
	data := make([]byte, messageHeader.Length)
	_, err = io.ReadFull(c.Conn, data)
	if err != nil {
		return
	}
	msg = &message.Message{
		Header:   messageHeader,
		PlayLoad: data,
	}
	return
}

func (c *Client) Send(msg *message.Message) error {
	header := message.EncodeMessageHeader(msg.Header)
	output := append(header, msg.PlayLoad...)
	n, err := c.Conn.Write(output)
	if err != nil {
		return err
	}
	if n != len(output) {
		panic(err)
	}
	return nil
}

func (c *Client) processMessage(msg *message.Message) {
	// fmt.Printf("type: %d, %s\n", c.Type, string(data))
	fmt.Printf("process message. type is %d\n", msg.Header.Type)
	switch msg.Header.Type {
	case message.Ping:
		fmt.Println("receive ping message...")
		time.Sleep(time.Second)
		pongMsg := &message.Message{
			Header: &message.MessageHeader{
				Type:   message.Pong,
				Length: 4,
			},
			PlayLoad: []byte("pong"),
		}
		c.Send(pongMsg)
	case message.Pong:
		// do nothing
		fmt.Println("receive pong message...")
		fmt.Println(string(msg.PlayLoad))
		pingMsg := &message.Message{
			Header: &message.MessageHeader{
				Type:   message.Ping,
				Length: 4,
			},
			PlayLoad: []byte("ping"),
		}
		c.Send(pingMsg)
	}
}
