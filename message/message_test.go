package message

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, 0x4646)

	fmt.Println("hello")
	v := binary.BigEndian.Uint16(buf)
	fmt.Println(v)

	v >>= 4
	len := v & 0xf
	v >>= 4
	magic := v & 0xff

	fmt.Println(len)
	fmt.Println(magic)
}
