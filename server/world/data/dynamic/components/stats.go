package components

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

type BasicStats struct {
	Strength  int
	Agility   int
	Stamina   int
	Intellect int
	Spirit    int
}

func (bs *BasicStats) regenForTimePeriod(max int, timePassed time.Duration, isInCombat bool) int {
	combatPenalty := 1.0
	if isInCombat {
		combatPenalty = 1.0 - PlayerRegenCombatPenaltyPercent
	}

	statPerSecondFromSpirit := int(combatPenalty * PlayerRegenPercentPerSpiritPerSecond * float64(bs.Spirit))
	statPerSecond := int(combatPenalty*PlayerRegenPercentBasePerSecond*float64(max)) + statPerSecondFromSpirit

	return int(float64(statPerSecond) * (float64(timePassed) / float64(time.Second)))
}

// HealthRegen returns the amount of health generated in a particular time period.
func (bs *BasicStats) HealthRegen(timePassed time.Duration, isInCombat bool) int {
	return bs.regenForTimePeriod(bs.MaxHealth(), timePassed, isInCombat)
}

// PowerRegen returns the amount of health generated in a particular time period.
func (bs *BasicStats) PowerRegen(timePassed time.Duration, isInCombat bool) int {
	return bs.regenForTimePeriod(bs.MaxPower(), timePassed, isInCombat)
}

// MaxHealth returns the maximum health given the current stats.
func (bs *BasicStats) MaxHealth() int {
	return HealthPerStamina * bs.Stamina
}

// MaxPower returns the maximum power given the current stats.
func (bs *BasicStats) MaxPower() int {
	return ManaPerIntellect * bs.Intellect
}

// MeleeAttackPower returns the melee attack power contribution from stats.
func (bs *BasicStats) MeleeAttackPower(class *static.Class) int {
	switch class {
	case static.ClassHunter:
	case static.ClassRogue:
		return bs.Strength*1 + bs.Agility*1
	case static.ClassPriest:
	case static.ClassMage:
	case static.ClassWarlock:
		return bs.Strength * 1
	}

	return bs.Strength * 2
}

// RangedAttackPower returns the melee attack power contribution from stats.
func (bs *BasicStats) RangedAttackPower(class *static.Class) int {
	switch class {
	case static.ClassWarrior:
	case static.ClassHunter:
	case static.ClassRogue:
		return bs.Agility * 2
	}

	return bs.Agility * 1
}
