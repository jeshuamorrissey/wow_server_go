package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientPlayerLogin(pkt *packet.ClientPlayerLogin, state *system.State) ([]system.ServerPacket, error) {
	state.Log.Infof("%v %v", pkt, state.OM.Players)
	if !state.OM.Exists(pkt.GUID) {
		state.Log.Errorf("Attempt to log in with unknown GUID %v!", pkt.GUID)
		return []system.ServerPacket{}, nil
	}

	player := state.OM.GetPlayer(pkt.GUID)
	state.Updater.Login(player.GUID(), state.Session)
	state.Character = player

	return []system.ServerPacket{
		&packet.ServerLoginVerifyWorld{
			Character: player,
		},
		&packet.ServerAccountDataTimes{},
		&packet.ServerTutorialFlags{},
		&packet.ServerInitWorldStates{
			Map:  uint32(player.MapID),
			Zone: uint32(player.ZoneID),
		},
	}, nil
}
