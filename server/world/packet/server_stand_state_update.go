package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerStandStateUpdate is sent back in response to ClientPing.
type ServerStandStateUpdate struct {
	State static.StandState
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerStandStateUpdate) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.State)

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerStandStateUpdate) OpCode() static.OpCode {
	return static.OpCodeServerStandstateUpdate
}
