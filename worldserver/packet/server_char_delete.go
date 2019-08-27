package packet

import (
	"bytes"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerCharDelete is sent from the client when making a character.
type ServerCharDelete struct {
	Error c.CharErrorCode
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCharDelete) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCharDelete) OpCode() system.OpCode {
	return system.OpCodeServerCharDelete
}
