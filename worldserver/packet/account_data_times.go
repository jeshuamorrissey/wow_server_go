package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerAccountDataTimes is sent back in response to ClientPing.
type ServerAccountDataTimes struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAccountDataTimes) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	for i := 0; i < 32; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAccountDataTimes) OpCode() system.OpCode {
	return system.OpCodeServerAccountDataTimes
}
