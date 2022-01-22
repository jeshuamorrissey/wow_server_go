package channels

import "github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"

type CombatUpdate struct {
	Attacker   interfaces.Object
	Target     interfaces.Object
	AttackInfo *interfaces.AttackInfo
}

var (
	CombatUpdates chan *CombatUpdate   = make(chan *CombatUpdate)
	ObjectUpdates chan interfaces.GUID = make(chan interfaces.GUID)
)
