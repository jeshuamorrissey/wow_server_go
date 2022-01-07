package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

func HandleClientAttackSwing(pkt *packet.ClientAttackSwing, state *system.State) ([]system.ServerPacket, error) {
	target := state.OM.Get(pkt.Target)
	if target == nil {
		state.Log.Warnf("Received CLIENT_ATTACK_SWING with non-existant target %v (%v)", pkt.Target.Low(), pkt.Target.High())
		return []system.ServerPacket{}, nil
	}

	// Setup a callback using the battle syste.
	u1 := state.OM.GetUnit(interfaces.MakeGUID(1, static.HighGUIDUnit))
	u2 := state.OM.GetUnit(interfaces.MakeGUID(2, static.HighGUIDUnit))
	state.CombatManager.StartAttack(u1, u2)
	state.CombatManager.StartAttack(u2, u1)

	return []system.ServerPacket{
		&packet.ServerAttackStart{Attacker: state.Character.GUID(), Target: target.GUID()},
	}, nil
}
