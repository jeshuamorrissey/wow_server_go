package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientAttackSwing is sent from the client periodically.
type ClientAttackSwing struct {
	Target interfaces.GUID
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

	// Setup a callback using the battle syste.
	u1 := state.OM.GetUnit(interfaces.MakeGUID(1, static.HighGUIDUnit))
	u2 := state.OM.GetUnit(interfaces.MakeGUID(2, static.HighGUIDUnit))
	state.CombatManager.StartAttack(u1, u2)
	state.CombatManager.StartAttack(u2, u1)

	return []system.ServerPacket{
		&ServerAttackStart{Attacker: state.Character.GUID(), Target: target.GUID()},
	}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientAttackSwing) OpCode() static.OpCode {
	return static.OpCodeClientAttackswing
}
