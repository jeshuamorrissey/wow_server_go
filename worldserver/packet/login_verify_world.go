package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/objects"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ServerLoginVerifyWorld is sent back in response to ClientPing.
type ServerLoginVerifyWorld struct {
	Character *objects.Player
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginVerifyWorld) Bytes(stateBase session.State) []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Character.MapID))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location.X))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location.Y))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location.Z))
	binary.Write(buffer, binary.LittleEndian, float32(pkt.Character.Location.O))

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerLoginVerifyWorld) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerLoginVerifyWorld)
}
