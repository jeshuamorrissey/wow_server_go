package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerAccountDataTimes is sent back in response to ClientPing.
type ServerAccountDataTimes struct{}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerAccountDataTimes) Bytes(session *system.Session) []byte {
	buffer := bytes.NewBufferString("")

	for i := 0; i < 32; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerAccountDataTimes) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerAccountDataTimes)
}
