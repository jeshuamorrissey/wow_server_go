package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientPlayerLogin(pkt *packet.ClientPlayerLogin, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Log.Infof("%v %v", pkt, state.OM.Players)
	if !state.OM.Exists(pkt.GUID) {
		state.Log.Errorf("Attempt to log in with unknown GUID %v!", pkt.GUID)
		return []interfaces.ServerPacket{}, nil
	}

	player := state.OM.GetPlayer(pkt.GUID)
	state.Updater.Login(player.GUID(), state.Session)
	state.Character = player

	return []interfaces.ServerPacket{
		&packet.ServerLoginVerifyWorld{
			MapID:    player.MapID,
			Location: player.Location,
		},
		&packet.ServerAccountDataTimes{},
		&packet.ServerTutorialFlags{
			Tutorials: player.Tutorials,
		},
		&packet.ServerInitWorldStates{
			Map:  uint32(player.MapID),
			Zone: uint32(player.ZoneID),
		},
	}, nil
}
