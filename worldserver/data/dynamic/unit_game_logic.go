package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
)

// Constants related to unit logic.
const (
	HealthPerStamina = 10
	ManaPerIntellect = 20
)

// Unit interface methods (game-logic).
func (u *Unit) MeleeMainHandAttackRate() time.Duration {
	return time.Duration(1000)
}

func (u *Unit) MeleeOffHandAttackRate() time.Duration {
	return time.Duration(0)
}

func (u *Unit) ResolveMeleeAttack(target interfaces.Unit) *interfaces.AttackInfo {
	return &interfaces.AttackInfo{
		Damage: 100,
	}
}

// Utility methods.
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
