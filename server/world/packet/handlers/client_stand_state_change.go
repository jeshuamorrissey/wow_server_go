package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientStandStateChange(pkt *packet.ClientStandStateChange, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Character.StandState = pkt.State

	response := new(packet.ServerStandStateUpdate)
	response.State = pkt.State

	return []interfaces.ServerPacket{response}, nil
}
