package channels

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
)

type CombatUpdate struct {
	Attacker   interfaces.Object
	Target     interfaces.Object
	AttackInfo *interfaces.AttackInfo
}

type PacketUpdate struct {
	SendTo   interfaces.GUID // can be set to 0 to mean "broadcast"
	Packet   interfaces.ServerPacket
	Location *interfaces.Location
}

var (
	CombatUpdates chan *CombatUpdate   = make(chan *CombatUpdate)
	ObjectUpdates chan interfaces.GUID = make(chan interfaces.GUID)
	PacketUpdates chan *PacketUpdate   = make(chan *PacketUpdate)
)
