package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ServerLogoutResponse is sent back in response to ClientPing.
type ServerLogoutResponse struct {
	Reason        uint32
	InstantLogout bool
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLogoutResponse) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Reason))
	if pkt.InstantLogout {
		binary.Write(buffer, binary.LittleEndian, uint8(1))
	} else {
		binary.Write(buffer, binary.LittleEndian, uint8(0))
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLogoutResponse) OpCode() static.OpCode {
	return static.OpCodeServerLogoutResponse
}
