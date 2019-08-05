package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ServerPong is sent back in response to ClientPing.
type ServerPong struct {
	Pong uint32
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerPong) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Pong)

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerPong) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerPong)
}
