package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/game"
)

// Unit interface methods (game-logic).
func (u *Unit) Initialize() {
	go u.regenHealthAndPower()
}

// Utility methods.
func (u *Unit) regenHealthAndPower() {
	for range time.Tick(static.RegenTimeout) {
		if u.CurrentHealth == 0 {
			return
		}

		secondsInTimeout := static.RegenTimeout / time.Second
		healthMod := game.UnitRegenPerSecond(u.maxHealth(), u.IsInCombat()) * int(secondsInTimeout)
		powerMod := game.UnitRegenPerSecond(u.maxPower(), u.IsInCombat()) * int(secondsInTimeout)

		u.UpdateChannel() <- []interface{}{
			&messages.UnitModHealth{Amount: healthMod},
			&messages.UnitModPower{Amount: powerMod},
		}
	}
}

func (u *Unit) maxHealth() int {
	return u.Template().MaxHealth
}

func (u *Unit) maxPower() int {
	return u.Template().MaxPower
}
