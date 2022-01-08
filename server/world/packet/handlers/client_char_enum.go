package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientCharEnum(pkt *packet.ClientCharEnum, state *system.State) ([]system.ServerPacket, error) {
	response := new(packet.ServerCharEnum)
	response.Characters = append(response.Characters, state.Account.Character)
	response.ObjectManager = state.Config.ObjectManager
	return []system.ServerPacket{response}, nil
}
