package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerAttackerStateUpdate is sent back in response to ClientPing.
type ServerAttackerStateUpdate struct {
	HitInfo        static.HitInfo
	Attacker       interfaces.GUID
	Target         interfaces.GUID
	Damage         int32
	OriginalDamage int32
	OverDamage     int32
	TargetState    static.AttackTargetState
	AttackerState  uint32
	MeleeSpellID   uint32
	BlockAmount    uint32
	RageGained     uint32
	Absorb         uint32
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerAttackerStateUpdate) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.HitInfo))
	buffer.Write(pkt.Attacker.Pack())
	buffer.Write(pkt.Target.Pack())
	binary.Write(buffer, binary.LittleEndian, int32(pkt.Damage))

	is_sub_damage := 1
	binary.Write(buffer, binary.LittleEndian, uint8(is_sub_damage)) // TODO: SubDamage
	if is_sub_damage != 0 {
		binary.Write(buffer, binary.LittleEndian, uint32(static.SpellSchoolPhysical)) // Damage school mask
		binary.Write(buffer, binary.LittleEndian, float32(pkt.Damage))                // sub damage
		binary.Write(buffer, binary.LittleEndian, uint32(pkt.Damage))                 // sub damage
		binary.Write(buffer, binary.LittleEndian, int32(pkt.Absorb))                  // absorbed
		binary.Write(buffer, binary.LittleEndian, int32(0))                           // reissted
	}

	binary.Write(buffer, binary.LittleEndian, uint8(pkt.TargetState))
	if pkt.Absorb == 0 {
		binary.Write(buffer, binary.LittleEndian, uint32(0))
	} else {
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // should be -1 by unsigned??
	}

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.MeleeSpellID))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.BlockAmount))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerAttackerStateUpdate) OpCode() static.OpCode {
	return static.OpCodeServerAttackerstateupdate
}
