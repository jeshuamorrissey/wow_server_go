package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerNameQueryResponse is sent back in response to ClientPing.
type ServerNameQueryResponse struct {
	RealmName string
	Character *config.Character
	Player    *dynamic.Player
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerNameQueryResponse) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint64(pkt.Player.GUID()))

	buffer.WriteString(pkt.Character.Name)
	buffer.WriteByte('\x00')

	buffer.WriteString(pkt.RealmName)
	buffer.WriteByte('\x00')

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Player.Race.ID))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Player.Gender))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Player.Class.ID))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerNameQueryResponse) OpCode() static.OpCode {
	return static.OpCodeServerNameQueryResponse
}
