package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

func HandleClientAttackStop(pkt *packet.ClientAttackStop, state *system.State) ([]system.ServerPacket, error) {
	state.CombatManager.StopAttack(state.Character)

	attacker := state.Character
	var targetGUID interfaces.GUID
	if target := state.CombatManager.GetTargetOf(attacker); target != nil {
		targetGUID = target.GUID()
	}

	return []system.ServerPacket{
		&packet.ServerAttackStop{
			Attacker: attacker.GUID(),
			Target:   targetGUID,
		},
	}, nil
}
