package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ServerPong is sent back in response to ClientPing.
type ServerPong struct {
	Pong uint32
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerPong) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Pong)

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerPong) OpCode() static.OpCode {
	return static.OpCodeServerPong
}
