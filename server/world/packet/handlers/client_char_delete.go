package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientCharDelete(pkt *packet.ClientCharDelete, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerCharDelete)
	response.Error = static.CharErrorCodeDeleteFailed
	return []interfaces.ServerPacket{response}, nil
}
