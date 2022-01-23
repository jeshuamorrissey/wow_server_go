package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerLoginVerifyWorld is sent back in response to ClientPing.
type ServerLoginVerifyWorld struct {
	MapID    int
	Location interfaces.Location
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLoginVerifyWorld) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.MapID))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Location.X))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Location.Y))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Location.Z))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Location.O))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLoginVerifyWorld) OpCode() static.OpCode {
	return static.OpCodeServerLoginVerifyWorld
}
