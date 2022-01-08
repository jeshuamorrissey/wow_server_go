package packet

import (
	"bytes"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerLogoutComplete is sent back in response to ClientPing.
type ServerLogoutComplete struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLogoutComplete) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLogoutComplete) OpCode() static.OpCode {
	return static.OpCodeServerLogoutComplete
}
