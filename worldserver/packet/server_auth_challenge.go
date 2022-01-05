package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerAuthChallenge is the initial message sent from the server
// to the client to start authorization.
type ServerAuthChallenge struct {
	Seed uint32
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAuthChallenge) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.BigEndian, pkt.Seed)

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAuthChallenge) OpCode() static.OpCode {
	return static.OpCodeServerAuthChallenge
}
