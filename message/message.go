package message

import (
	"encoding/binary"
)

const MagicHeader uint16 = 0x46

type MessageType uint8

const (
	Ping MessageType = iota + 1
	Pong
)

// bytes
const HeaderLength = 4

type MessageHeader struct {
	Type   MessageType
	Length uint16
}

type Message struct {
	Header   *MessageHeader
	PlayLoad []byte
}

func DecodeMessageHeader(data []byte) *MessageHeader {
	if len(data) != HeaderLength {
		return nil
	}
	// magic and type
	h1 := binary.BigEndian.Uint16(data[:2])
	// 4 bit space
	h1 >>= 4
	// message type
	messageType := h1 & 0xf
	h1 >>= 4
	// magic
	magciNum := h1 & 0xff
	if magciNum != MagicHeader {
		panic("")
	}
	// message playload length
	messageLength := binary.BigEndian.Uint16(data[2:4])

	mh := &MessageHeader{
		Type:   MessageType(messageType),
		Length: messageLength,
	}
	return mh
}

func EncodeMessageHeader(header *MessageHeader) []byte {
	data := make([]byte, 4)
	// magic and type
	var h1 uint16
	h1 = MagicHeader
	h1 = h1 << 4
	h1 = h1 | (uint16(header.Type))
	h1 = h1 << 4
	h1 = h1 | (header.Length)
	binary.BigEndian.PutUint16(data[:2], h1)
	binary.BigEndian.PutUint16(data[2:], header.Length)
	return data
}

func DecodeMessagePlayload(data []byte) {}

func DecodeMessage(data []byte) *Message {
	return &Message{}
}
