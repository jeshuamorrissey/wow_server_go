package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientUpdateAccountData(pkt *packet.ClientUpdateAccountData, state *system.State) ([]interfaces.ServerPacket, error) {
	// Not implemented.
	return nil, nil
}
