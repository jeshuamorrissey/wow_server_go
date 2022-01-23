package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

func HandleClientAttackSwing(pkt *packet.ClientAttackSwing, state *system.State) ([]interfaces.ServerPacket, error) {
	target := state.OM.Get(pkt.Target)
	if target == nil {
		state.Log.Warnf("Received CLIENT_ATTACK_SWING with non-existant target %v (%v)", pkt.Target.Low(), pkt.Target.High())
		return []interfaces.ServerPacket{}, nil
	}

	state.Character.SendUpdates([]interface{}{
		&messages.UnitAttack{Target: target.GUID()},
	})

	return []interfaces.ServerPacket{}, nil
}
