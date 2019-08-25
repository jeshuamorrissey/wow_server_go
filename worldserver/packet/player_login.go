package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
)

// ClientPlayerLogin is sent from the client periodically.
type ClientPlayerLogin struct {
	GUID object.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientPlayerLogin) FromBytes(state *State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientPlayerLogin) Handle(state *State) ([]ServerPacket, error) {
	if !state.OM.Exists(pkt.GUID) {
		state.Log.Errorf("Attempt to log in with unknown GUID %v!", pkt.GUID)
		return []ServerPacket{}, nil
	}

	player := state.OM.Get(pkt.GUID).(*object.Player)
	state.Log.Infof("player = %v", player)

	return []ServerPacket{
		&ServerLoginVerifyWorld{
			Character: player,
		},
		&ServerAccountDataTimes{},
		&ServerTutorialFlags{},
		&ServerInitWorldStates{
			Map:  uint32(player.MapID),
			Zone: uint32(player.ZoneID),
		},
	}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientPlayerLogin) OpCode() OpCode {
	return OpCodeClientPlayerLogin
}
