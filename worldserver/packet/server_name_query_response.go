package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/config"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerNameQueryResponse is sent back in response to ClientPing.
type ServerNameQueryResponse struct {
	Character *config.Character
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerNameQueryResponse) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	player := state.OM.Get(pkt.Character.GUID).(*dynamic.Player)
	binary.Write(buffer, binary.LittleEndian, uint64(player.GUID()))

	buffer.WriteString(pkt.Character.Name)
	buffer.WriteByte('\x00')

	buffer.WriteString(state.Config.Name)
	buffer.WriteByte('\x00')

	binary.Write(buffer, binary.LittleEndian, uint32(player.Race.ID))
	binary.Write(buffer, binary.LittleEndian, uint32(player.Gender))
	binary.Write(buffer, binary.LittleEndian, uint32(player.Class.ID))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerNameQueryResponse) OpCode() static.OpCode {
	return static.OpCodeServerNameQueryResponse
}
