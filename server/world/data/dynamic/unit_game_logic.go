package dynamic

import (
	"math"
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Constants related to unit logic.
const (
	HealthPerStamina = 10
	ManaPerIntellect = 20

	ManaRegenPercentBasePerSecond      float64 = 0.05         // 5% mana regen per second (20 seconds to fully recover mana)
	ManaRegenPercentPerSpiritPerSecond float64 = 0.05 / 100.0 // 5% mana regen per 100 spirit (~100 at level 60)
	RegenTimeoutMS                             = 1000         // timeout (in ms) between regen events.
)

// Unit interface methods (game-logic).
func (u *Unit) Initialize() {
	u.IsActive = true

	go u.restoreHealthPower()
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
	u.InCombat = true
}

// Utility methods.
func (u *Unit) restoreHealthPower() {
	for range time.Tick(time.Millisecond * RegenTimeoutMS) {
		if u.IsActive {
			secondsInTimeout := RegenTimeoutMS / 1000.0
			manaPercentToRestore := ManaRegenPercentBasePerSecond*secondsInTimeout + ManaRegenPercentPerSpiritPerSecond*secondsInTimeout*float64(u.Spirit)

			// If we are in combat, penalize regen.
			combatMultiplier := 1.0
			if u.InCombat {
				combatMultiplier = 0.1
			}

			u.PowerPercent = float32(math.Min(float64(u.PowerPercent)+manaPercentToRestore*combatMultiplier, 1.0))
			GetObjectManager().TriggerUpdateFor(u)
		}
	}
}

func (u *Unit) powerType() static.Power {
	return static.PowerMana
	// return u.Class.PowerType
}

func (u *Unit) maxHealth() int {
	return u.BaseHealth + HealthPerStamina*u.Stamina
}

func (u *Unit) maxPower() int {
	switch u.powerType() {
	case static.PowerMana:
		return u.Intellect * ManaPerIntellect
	case static.PowerRage:
	case static.PowerFocus:
	case static.PowerEnergy:
	case static.PowerHappiness:
		return 100
	}

	return 0
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
