package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ServerAuthChallenge is the initial message sent from the server
// to the client to start authorization.
type ServerAuthChallenge struct {
	Seed uint32
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerAuthChallenge) Bytes(stateBase session.State) []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.BigEndian, pkt.Seed)

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerAuthChallenge) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerAuthChallenge)
}
