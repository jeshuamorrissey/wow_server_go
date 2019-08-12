package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ServerAuthResponse is the initial message sent from the server
// to the client to start authorization.
type ServerAuthResponse struct {
	Error AuthErrorCode
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerAuthResponse) Bytes(stateBase session.State) []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Error)
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // unk
	binary.Write(buffer, binary.LittleEndian, uint8(0))  // unk
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // unk

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerAuthResponse) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerAuthResponse)
}
