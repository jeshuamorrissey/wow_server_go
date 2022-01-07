package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientSetActiveMover(pkt *packet.ClientSetActiveMover, state *system.State) ([]system.ServerPacket, error) {
	if pkt.GUID != state.Character.GUID() {
		state.Log.Errorf("Incorrect mover GUID: it is %v, but should be %v", pkt.GUID, state.Character.GUID())
	}

	return nil, nil
}
