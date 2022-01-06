package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientPlayerLogin is sent from the client periodically.
type ClientPlayerLogin struct {
	GUID interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientPlayerLogin) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientPlayerLogin) Handle(state *system.State) ([]system.ServerPacket, error) {
	state.Log.Infof("%v %v", pkt, state.OM.Players)
	if !state.OM.Exists(pkt.GUID) {
		state.Log.Errorf("Attempt to log in with unknown GUID %v!", pkt.GUID)
		return []system.ServerPacket{}, nil
	}

	player := state.OM.GetPlayer(pkt.GUID)
	state.Updater.Login(player.GUID(), state.Session)
	state.Character = player

	return []system.ServerPacket{
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
func (*ClientPlayerLogin) OpCode() static.OpCode {
	return static.OpCodeClientPlayerLogin
}
