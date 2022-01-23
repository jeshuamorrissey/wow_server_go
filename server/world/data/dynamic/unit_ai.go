package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
)

func (u *Unit) HandleAttack(attacker interfaces.GUID) {
	// TODO(jeshua): more complex AI.
	u.SendUpdates([]interface{}{
		&messages.UnitAttack{Target: attacker},
	})
}

func (u *Unit) HandleAttackStop(attacker interfaces.GUID) {
}
