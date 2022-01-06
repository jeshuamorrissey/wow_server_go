package packet

import (
	"bytes"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ServerCharDelete is sent from the client when making a character.
type ServerCharDelete struct {
	Error static.CharErrorCode
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCharDelete) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCharDelete) OpCode() static.OpCode {
	return static.OpCodeServerCharDelete
}
