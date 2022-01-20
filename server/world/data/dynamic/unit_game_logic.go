package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/game"
)

// Constants related to unit logic.
const (
	HealthPerStamina = 10
	ManaPerIntellect = 20

	RegenTimeoutMS = 1000 // timeout (in ms) between regen events.
)

// Unit interface methods (game-logic).
func (u *Unit) Initialize() {
	go u.regenHealthAndPower()
}

func (u *Unit) MeleeMainHandAttackRate() time.Duration {
	return time.Duration(1000)
}

func (u *Unit) MeleeOffHandAttackRate() time.Duration {
	return time.Duration(0)
}

func (u *Unit) ResolveMainHandAttack(target interfaces.Unit) *interfaces.AttackInfo {
	return &interfaces.AttackInfo{
		Damage: 100,
	}
}

func (u *Unit) ResolveOffHandAttack(target interfaces.Unit) *interfaces.AttackInfo {
	return &interfaces.AttackInfo{
		Damage: 100,
	}
}

func (u *Unit) SetInCombat(inCombat bool) {
	u.InCombat = inCombat
}

// Utility methods.
func (u *Unit) regenHealthAndPower() {
	for range time.Tick(time.Millisecond * RegenTimeoutMS) {
		secondsInTimeout := RegenTimeoutMS / 1000.0

		u.CurrentHealth += game.UnitRegenPerSecond(u.maxHealth(), u.InCombat) * int(secondsInTimeout)
		if u.CurrentHealth >= u.maxHealth() {
			u.CurrentHealth = u.maxHealth()
		}

		u.CurrentPower += game.UnitRegenPerSecond(u.maxPower(), u.InCombat) * int(secondsInTimeout)
		if u.CurrentPower >= u.maxPower() {
			u.CurrentPower = u.maxPower()
		}

		// fmt.Printf("Regening unit %v %v\n", u.ID.Low(), u.InCombat)

		GetObjectManager().TriggerUpdateFor(u)
	}
}

func (u *Unit) powerType() static.Power {
	return static.PowerMana
	// return u.Class.PowerType
}

func (u *Unit) TakeDamage(damage int) {
	u.CurrentHealth = u.CurrentHealth - damage
	if u.CurrentHealth < 0 {
		u.CurrentHealth = 0

		// Trigger a respawn timer for this unit, if appropriate.
		if u.RespawnTimeMS != 0 {
			time.AfterFunc(u.RespawnTimeMS, func() {
				u.CurrentHealth = u.maxHealth()
				GetObjectManager().triggerUpdateFor(u)
			})
		}
	}
}

func (u *Unit) maxHealth() int {
	return u.Template().MaxHealth
}

func (u *Unit) maxPower() int {
	return u.Template().MaxPower
}

func (u *Unit) meleeAttackPower() int {
	switch u.Class {
	case static.ClassWarrior:
	case static.ClassPaladin:
		return u.Strength * 2
	case static.ClassHunter:
	case static.ClassRogue:
		return u.Strength*1 + u.Agility*1
	case static.ClassShaman:
		return u.Strength * 2
	case static.ClassPriest:
	case static.ClassMage:
	case static.ClassWarlock:
		return u.Strength * 1
	case static.ClassDruid:
		return u.Strength * 2
		// TODO: if in cat form, increase by u.Agility * 1
	}

	return 0
}

func (u *Unit) MeleeAttackRate() time.Duration {
	return time.Duration(1000) * time.Millisecond
}

func (u *Unit) rangedAttackPower() int {
	switch u.Class {
	case static.ClassWarrior:
	case static.ClassHunter:
	case static.ClassRogue:
		return u.Agility * 2
	}

	return u.Agility * 1
}
