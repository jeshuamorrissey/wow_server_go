package packet

import (
	"encoding/binary"
	"io"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientAttackSwing is sent from the client periodically.
type ClientAttackSwing struct {
	Target object.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientAttackSwing) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Target)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientAttackSwing) Handle(state *system.State) ([]system.ServerPacket, error) {
	target := state.OM.Get(pkt.Target)
	if target == nil {
		state.Log.Warnf("Received CLIENT_ATTACK_SWING with non-existant target %v (%v)", pkt.Target.Low(), pkt.Target.High())
		return []system.ServerPacket{}, nil
	}

	state.Log.Infof("Player is attacking %v", target.(*object.Unit).Template().Name)

	// Setup a callback using the battle syste.
	u1 := state.OM.Get(object.MakeGUID(1, c.HighGUIDUnit))
	u2 := state.OM.Get(object.MakeGUID(2, c.HighGUIDUnit))
	state.CombatManager.StartAttack(u1.(object.UnitInterface), u2.(object.UnitInterface))
	state.CombatManager.StartAttack(u2.(object.UnitInterface), u1.(object.UnitInterface))

	return []system.ServerPacket{
		&ServerAttackStart{Attacker: state.Character.GUID(), Target: target.GUID()},
	}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientAttackSwing) OpCode() system.OpCode {
	return system.OpCodeClientAttackswing
}
