package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerAttackStart is sent back in response to ClientPing.
type ServerAttackStart struct {
	Attacker interfaces.GUID
	Target   interfaces.GUID
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAttackStart) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Attacker)
	binary.Write(buffer, binary.LittleEndian, pkt.Target)

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAttackStart) OpCode() static.OpCode {
	return static.OpCodeServerAttackstart
}
