package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
)

// ServerLoginVerifyWorld is sent back in response to ClientPing.
type ServerLoginVerifyWorld struct {
	Character *object.Player
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerLoginVerifyWorld) ToBytes(state *State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Character.MapID))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location().X))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location().Y))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location().Z))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location().O))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLoginVerifyWorld) OpCode() OpCode {
	return OpCodeServerLoginVerifyWorld
}
