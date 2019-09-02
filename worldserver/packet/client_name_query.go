package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jinzhu/gorm"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientNameQuery is sent from the client periodically.
type ClientNameQuery struct {
	GUID object.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientNameQuery) FromBytes(state *system.State, buffer io.Reader) error {
	return binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
}

// Handle will ensure that the given account exists.
func (pkt *ClientNameQuery) Handle(state *system.State) ([]system.ServerPacket, error) {
	if !state.OM.Exists(pkt.GUID) {
		return nil, nil
	}

	response := new(ServerNameQueryResponse)
	response.Character = new(database.Character)
	err := state.DB.Where(&database.Character{GUID: pkt.GUID}).First(response.Character).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientNameQuery) OpCode() system.OpCode {
	return system.OpCodeClientNameQuery
}
