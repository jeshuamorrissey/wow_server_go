package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientCharEnum is sent from the client when first connecting.
type ClientCharEnum struct {
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientCharEnum) FromBytes(state *system.State, buffer io.Reader) error {
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharEnum) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCharEnum)

	err := state.DB.Where(&database.Character{AccountID: state.Account.ID, RealmID: state.Realm.ID}).Find(&response.Characters).Error
	if err != nil {
		return nil, err
	}

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharEnum) OpCode() system.OpCode {
	return system.OpCodeClientCharEnum
}
