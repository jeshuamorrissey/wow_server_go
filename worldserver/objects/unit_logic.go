package objects

import c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"

// Constants related to unit logic.
const (
	HealthPerStamina = 10
	ManaPerIntellect = 20
)

func (u *Unit) powerType() c.Power {
	return c.PowerMana
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
