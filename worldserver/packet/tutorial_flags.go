package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ServerTutorialFlags is sent back in response to ClientPing.
type ServerTutorialFlags struct{}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerTutorialFlags) Bytes(stateBase session.State) []byte {
	buffer := bytes.NewBufferString("")

	// TODO(jeshua): implement tutorials.
	for i := 0; i < 8; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerTutorialFlags) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerTutorialFlags)
}
