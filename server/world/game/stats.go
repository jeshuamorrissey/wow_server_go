package game

const (
	HealthPerStamina = 10

	PlayerRegenPercentBasePerSecond      float64 = 0.05         // 5% mana regen per second (20 seconds to fully recover mana)
	PlayerRegenPercentPerSpiritPerSecond float64 = 0.05 / 100.0 // 5% mana regen per 100 spirit (~100 at level 60)
	PlayerRegenCombatPenaltyPercent      float64 = 0.80         // 80% reduction in regen when in combat

	UnitRegentPercentPerSecond     float64 = 0.3 // 30% regen per second
	UnitRegentCombatPenaltyPercent float64 = 1.0 // 100% reduction in regen when in combat
)

func HealthFromStamina(stamina int) int {
	return stamina * HealthPerStamina
}

func PlayerRegenPerSecond(max int, spirit int, inCombat bool) int {
	combatPenalty := 1.0
	if inCombat {
		combatPenalty = 1.0 - PlayerRegenCombatPenaltyPercent
	}

	healthPerSecond := int(combatPenalty * PlayerRegenPercentBasePerSecond * float64(max))
	healthPerSecondFromSpirit := int(combatPenalty * PlayerRegenPercentPerSpiritPerSecond * float64(spirit))

	return healthPerSecond + healthPerSecondFromSpirit
}

func UnitRegenPerSecond(max int, inCombat bool) int {
	combatPenalty := 1.0
	if inCombat {
		combatPenalty = 1.0 - UnitRegentCombatPenaltyPercent
	}

	return int(combatPenalty * UnitRegentPercentPerSecond * float64(max))
}
