package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientTutorialFlag(pkt *packet.ClientTutorialFlag, state *system.State) ([]system.ServerPacket, error) {
	state.Character.Tutorials[pkt.Flag] = true
	return nil, nil
}
