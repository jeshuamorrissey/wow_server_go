package object

import (
	"math/rand"
	"time"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Resistances returns the final resitances for the player, after
// all modifications.
func (p *Player) Resistances() map[c.SpellSchool]int {
	resistances := map[c.SpellSchool]int{
		c.SpellSchoolPhysical: 0,
		c.SpellSchoolHoly:     0,
		c.SpellSchoolFire:     0,
		c.SpellSchoolNature:   0,
		c.SpellSchoolFrost:    0,
		c.SpellSchoolShadow:   0,
		c.SpellSchoolArcane:   0,
	}

	/// Add modifications based on items.
	for _, itemGUID := range p.Equipment {
		item := p.Manager().Get(itemGUID).(*Item)

		for k, v := range item.Template().Resistances {
			resistances[k] += v
		}
	}

	/// Add modifications based on stats.
	// Each point in agility gives 2 armor.
	resistances[c.SpellSchoolPhysical] += p.Agility * 2

	// Each point in spirit increases resistances by 0.05.
	spiritBonus := int(0.05 * float32(p.Spirit))
	resistances[c.SpellSchoolHoly] += spiritBonus
	resistances[c.SpellSchoolFire] += spiritBonus
	resistances[c.SpellSchoolNature] += spiritBonus
	resistances[c.SpellSchoolFrost] += spiritBonus
	resistances[c.SpellSchoolShadow] += spiritBonus
	resistances[c.SpellSchoolArcane] += spiritBonus

	return resistances
}

func (p *Player) meleeAttackPower() int {
	return p.Unit.meleeAttackPower() + p.meleeAttackPowerMods()
}

func (p *Player) meleeAttackRate() int {
	weapon := p.Manager().Get(p.Equipment[c.EquipmentSlotMainHand])
	if weapon == nil {
		return 2000
	}

	return int(weapon.(*Item).Template().AttackRate.Milliseconds())
}

func (p *Player) MeleeAttackRate() time.Duration {
	weapon := p.Manager().Get(p.Equipment[c.EquipmentSlotMainHand])
	if weapon == nil {
		return time.Duration(2000) * time.Millisecond
	}

	return weapon.(*Item).Template().AttackRate
}

func (p *Player) Attack(target UnitInterface) AttackInfo {
	weapon := p.Manager().Get(p.Equipment[c.EquipmentSlotMainHand])
	if weapon == nil {
		return AttackInfo{
			Damage: 0,
		}
	}

	minDamage := int(weapon.(*Item).Template().Damages[c.SpellSchoolPhysical].Min)
	maxDamage := int(weapon.(*Item).Template().Damages[c.SpellSchoolPhysical].Max)

	finalDamage := minDamage + rand.Intn(maxDamage-minDamage+1)

	// Calculate what % of the max this is, and reduce the % by that.
	finalDamagePercent := float32(finalDamage) / float32(target.(*Unit).Template().MaxHealth)

	target.(*Unit).HealthPercent -= finalDamagePercent
	p.manager.Update(target.GUID())

	return AttackInfo{
		Damage: finalDamage,
	}
}

func (p *Player) meleeAttackPowerMods() int {
	// TODO(jeshua): account for items
	return 0
}

func (p *Player) rangedAttackPower() int {
	return p.Unit.rangedAttackPower() + p.rangedAttackPowerMods()
}

func (p *Player) rangedAttackPowerMods() int {
	// TODO(jeshua): account for items
	return 0
}

func (p *Player) damageModPercentage() float32 {
	// TODO(jeshua): account for active spell effects
	// TODO(jeshua): account for items
	return 1.0
}
