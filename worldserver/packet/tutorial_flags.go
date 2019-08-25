package packet

import (
	"bytes"
	"encoding/binary"
)

// ServerTutorialFlags is sent back in response to ClientPing.
type ServerTutorialFlags struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerTutorialFlags) ToBytes(state *State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	// TODO(jeshua): implement tutorials.
	for i := 0; i < 8; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerTutorialFlags) OpCode() OpCode {
	return OpCodeServerTutorialFlags
}
