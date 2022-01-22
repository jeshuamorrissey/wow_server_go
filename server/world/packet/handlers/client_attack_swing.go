package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

func HandleClientAttackSwing(pkt *packet.ClientAttackSwing, state *system.State) ([]system.ServerPacket, error) {
	target := state.OM.Get(pkt.Target)
	if target == nil {
		state.Log.Warnf("Received CLIENT_ATTACK_SWING with non-existant target %v (%v)", pkt.Target.Low(), pkt.Target.High())
		return []system.ServerPacket{}, nil
	}

	state.Character.SendUpdates([]interface{}{
		&messages.UnitAttack{Target: target.GUID()},
	})

	target.SendUpdates([]interface{}{
		&messages.UnitRegisterAttack{Attacker: state.Character.GUID()},
	})

	return []system.ServerPacket{
		&packet.ServerAttackStart{Attacker: state.Character.GUID(), Target: target.GUID()},
	}, nil
}
