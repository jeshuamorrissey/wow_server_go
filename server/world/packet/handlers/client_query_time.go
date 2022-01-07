package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientQueryTime(pkt *packet.ClientQueryTime, state *system.State) ([]system.ServerPacket, error) {
	return []system.ServerPacket{new(packet.ServerQueryTimeResponse)}, nil
}
