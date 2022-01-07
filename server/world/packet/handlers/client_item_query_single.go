package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientItemQuerySingle(pkt *packet.ClientItemQuerySingle, state *system.State) ([]system.ServerPacket, error) {
	response := new(packet.ServerItemQuerySingleResponse)

	response.Entry = pkt.Entry
	response.Item = nil
	if item, ok := static.Items[int(pkt.Entry)]; ok {
		response.Item = item
	} else if pkt.GUID != 0 && state.OM.Exists(pkt.GUID) {
		if item := state.OM.GetItem(pkt.GUID); item != nil {
			response.Item = item.GetTemplate()
		} else if container := state.OM.GetContainer(pkt.GUID); container != nil {
			response.Item = container.GetTemplate()
		}
	}

	return []system.ServerPacket{response}, nil
}
