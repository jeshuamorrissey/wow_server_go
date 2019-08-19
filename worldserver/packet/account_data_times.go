package packet

import (
	"bytes"
	"encoding/binary"
)

// ServerAccountDataTimes is sent back in response to ClientPing.
type ServerAccountDataTimes struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAccountDataTimes) ToBytes(state *State) []byte {
	buffer := bytes.NewBufferString("")

	for i := 0; i < 32; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerAccountDataTimes) OpCode() OpCode {
	return OpCodeServerAccountDataTimes
}
