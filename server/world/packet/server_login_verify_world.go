package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ServerLoginVerifyWorld is sent back in response to ClientPing.
type ServerLoginVerifyWorld struct {
	Character *dynamic.Player
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLoginVerifyWorld) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Character.MapID))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.GetLocation().X))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.GetLocation().Y))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.GetLocation().Z))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.GetLocation().O))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLoginVerifyWorld) OpCode() static.OpCode {
	return static.OpCodeServerLoginVerifyWorld
}
