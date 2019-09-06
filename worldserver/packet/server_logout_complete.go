package packet

import (
	"bytes"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerLogoutComplete is sent back in response to ClientPing.
type ServerLogoutComplete struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLogoutComplete) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")
	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLogoutComplete) OpCode() system.OpCode {
	return system.OpCodeServerLogoutComplete
}
