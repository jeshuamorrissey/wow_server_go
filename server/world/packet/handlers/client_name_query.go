package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientNameQuery(pkt *packet.ClientNameQuery, state *system.State) ([]system.ServerPacket, error) {
	if !state.OM.Exists(pkt.GUID) {
		return nil, nil
	}

	response := new(packet.ServerNameQueryResponse)
	response.RealmName = state.Config.Name
	response.Character = state.Account.Character
	response.Player = state.Character
	return []system.ServerPacket{response}, nil
}
