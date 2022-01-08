package packet

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerQueryTimeResponse is sent back in response to ClientPing.
type ServerQueryTimeResponse struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerQueryTimeResponse) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.BigEndian, uint32(time.Now().Unix()))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerQueryTimeResponse) OpCode() static.OpCode {
	return static.OpCodeServerQueryTimeResponse
}
