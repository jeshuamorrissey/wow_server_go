package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerAttackStop is sent back in response to ClientPing.
type ServerAttackStop struct {
	Attacker interfaces.GUID
	Target   interfaces.GUID
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAttackStop) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.Attacker.Pack())
	if pkt.Target > 0 {
		binary.Write(buffer, binary.LittleEndian, pkt.Target.Pack())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // unk

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAttackStop) OpCode() static.OpCode {
	return static.OpCodeServerAttackstop
}
