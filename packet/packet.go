package packet

import (
	"bytes"
	"encoding/binary"
)

func maybeSetSize(currSize int, data []byte) int {
	if currSize == 0 && len(data) != 0 {
		return len(data)
	}

	return currSize
}

func maybeAllocate(data *[]byte, expectedLen int) {
	if len(*data) != expectedLen {
		*data = make([]byte, expectedLen)
	}
}

func MakeReader(buffer *bytes.Buffer) ProcessFunc {
	return func(byteOrder binary.ByteOrder, data interface{}) error {
		return binary.Read(buffer, byteOrder, data)
	}
}

func MakeWriter(buffer *bytes.Buffer) ProcessFunc {
	return func(byteOrder binary.ByteOrder, data interface{}) error {
		return binary.Write(buffer, byteOrder, data)
	}
}

type ProcessFunc func(binary.ByteOrder, interface{}) error

type ClientPacketA struct {
	A    uint8
	B    [4]byte
	C    [3]uint8
	DLen uint8
	D    []byte
}

func (pkt *ClientPacketA) Process(process ProcessFunc) error {
	process(binary.LittleEndian, &pkt.A)
	process(binary.LittleEndian, &pkt.B)

	if pkt.A == 0 {
		process(binary.LittleEndian, &pkt.C)
	}

	pkt.DLen = uint8(maybeSetSize(int(pkt.DLen), pkt.D))
	process(binary.LittleEndian, &pkt.DLen)
	maybeAllocate(&pkt.D, int(pkt.DLen))
	process(binary.LittleEndian, &pkt.D)
	return nil
}

type ServerPacketA struct {
	a int16
	b string
	c []struct {
		x float32
		y float32
		z float32
	}
}
