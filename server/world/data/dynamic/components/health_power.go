package components

import (
	"github.com/jeshuamorrissey/wow_server_go/lib/util"
)

const (
	HealthPerStamina = 10
	ManaPerIntellect = 20

	PlayerRegenPercentBasePerSecond      float64 = 0.05         // 5% mana regen per second (20 seconds to fully recover mana)
	PlayerRegenPercentPerSpiritPerSecond float64 = 0.05 / 100.0 // 5% mana regen per 100 spirit (~100 at level 60)
	PlayerRegenCombatPenaltyPercent      float64 = 0.80         // 80% reduction in regen when in combat

	UnitRegentPercentPerSecond     float64 = 0.3 // 30% regen per second
	UnitRegentCombatPenaltyPercent float64 = 1.0 // 100% reduction in regen when in combat
)

// HealthPower tracks
type HealthPower struct {
	CurrentHealth int
	CurrentPower  int
}

func (hp *HealthPower) ModHealth(healthMod int, maxHealth int) {
	hp.CurrentHealth = util.Clamp(0, hp.CurrentHealth+healthMod, maxHealth)
}

func (hp *HealthPower) ModPower(powerMod int, maxPower int) {
	hp.CurrentPower = util.Clamp(0, hp.CurrentPower+powerMod, maxPower)
}
