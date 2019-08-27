package packet

import (
	"bytes"
	"encoding/binary"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerAuthResponse is the initial message sent from the server
// to the client to start authorization.
type ServerAuthResponse struct {
	Error c.AuthErrorCode
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAuthResponse) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Error)
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // unk
	binary.Write(buffer, binary.LittleEndian, uint8(0))  // unk
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // unk

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAuthResponse) OpCode() system.OpCode {
	return system.OpCodeServerAuthResponse
}
