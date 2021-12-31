package object

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Constants related to unit logic.
const (
	HealthPerStamina = 10
	ManaPerIntellect = 20
)

func (u *Unit) powerType() c.Power {
	return u.Class.PowerType
}

func (u *Unit) maxHealth() int {
	return u.BaseHealth + HealthPerStamina*u.Stamina
}

func (u *Unit) maxPower() int {
	switch u.powerType() {
	case c.PowerMana:
		return u.Intellect * ManaPerIntellect
	case c.PowerRage:
	case c.PowerFocus:
	case c.PowerEnergy:
	case c.PowerHappiness:
		return 100
	}

	return 0
}

func (u *Unit) meleeAttackPower() int {
	switch u.Class {
	case dbc.ClassWarrior:
	case dbc.ClassPaladin:
		return u.Strength * 2
	case dbc.ClassHunter:
	case dbc.ClassRogue:
		return u.Strength*1 + u.Agility*1
	case dbc.ClassShaman:
		return u.Strength * 2
	case dbc.ClassPriest:
	case dbc.ClassMage:
	case dbc.ClassWarlock:
		return u.Strength * 1
	case dbc.ClassDruid:
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
	case dbc.ClassWarrior:
	case dbc.ClassHunter:
	case dbc.ClassRogue:
		return u.Agility * 2
	}

	return u.Agility * 1
}

type CombatInfo struct {
	Damage int32
}

func (u *Unit) Attack(target UnitInterface) AttackInfo {
	return AttackInfo{
		Damage: 100,
	}
}
