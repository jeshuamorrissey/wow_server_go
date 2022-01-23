package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientNameQuery(pkt *packet.ClientNameQuery, state *system.State) ([]interfaces.ServerPacket, error) {
	if !state.OM.Exists(pkt.GUID) {
		return nil, nil
	}

	return []interfaces.ServerPacket{&packet.ServerNameQueryResponse{
		RealmName:     state.Config.Name,
		CharacterName: state.Account.Character.Name,
		PlayerGUID:    state.Character.GUID(),
		PlayerRace:    state.Character.Race,
		PlayerGender:  state.Character.Gender,
		PlayerClass:   state.Character.Class,
	}}, nil
}
