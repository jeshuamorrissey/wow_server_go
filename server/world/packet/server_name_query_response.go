package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerNameQueryResponse is sent back in response to ClientPing.
type ServerNameQueryResponse struct {
	RealmName     string
	CharacterName string
	PlayerGUID    interfaces.GUID
	PlayerRace    *static.Race
	PlayerGender  static.Gender
	PlayerClass   *static.Class
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerNameQueryResponse) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint64(pkt.PlayerGUID))

	buffer.WriteString(pkt.CharacterName)
	buffer.WriteByte('\x00')

	buffer.WriteString(pkt.RealmName)
	buffer.WriteByte('\x00')

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.PlayerRace.ID))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.PlayerGender))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.PlayerClass.ID))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerNameQueryResponse) OpCode() static.OpCode {
	return static.OpCodeServerNameQueryResponse
}
