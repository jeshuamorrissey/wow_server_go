package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

func HandleClientAttackStop(pkt *packet.ClientAttackStop, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Character.SendUpdates([]interface{}{
		&messages.UnitStopAttack{},
	})

	return []interfaces.ServerPacket{
		&packet.ServerAttackStop{
			Attacker: state.Character.GUID(),
			Target:   state.Character.Target,
		},
	}, nil
}
