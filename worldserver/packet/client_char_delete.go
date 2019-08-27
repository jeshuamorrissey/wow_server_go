package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientCharDelete is sent from the client when deleting a character.
type ClientCharDelete struct {
	HighGUID c.HighGUID
	ID       uint32
}

// FromBytes loads the packet from the given data.
func (pkt *ClientCharDelete) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.HighGUID)
	binary.Read(buffer, binary.LittleEndian, &pkt.ID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharDelete) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCharDelete)
	response.Error = c.CharErrorCodeDeleteSuccess

	// Get the object.
	var char database.Character
	err := state.DB.Where("GUID = ?", uint64(pkt.HighGUID)<<32|uint64(pkt.ID)).First(&char).Error
	if err != nil {
		response.Error = c.CharErrorCodeDeleteFailed
		return []system.ServerPacket{response}, nil
	}

	state.DB.Unscoped().Delete(&char)

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharDelete) OpCode() system.OpCode {
	return system.OpCodeClientCharDelete
}
