package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientPing(pkt *packet.ClientPing, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerPong)
	response.Pong = pkt.Ping

	return []interfaces.ServerPacket{response}, nil
}
