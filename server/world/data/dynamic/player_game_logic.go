package dynamic

import (
	"math/rand"
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Unit interface methods (game-logic).
func (p *Player) Initialize() {
	p.Unit.Initialize()

	// Players always start at logged out when initialized.
	p.Unit.IsActive = false
}

func (p *Player) MeleeMainHandAttackRate() time.Duration {
	return p.meleeAttackRate(static.EquipmentSlotMainHand)
}

func (p *Player) MeleeOffHandAttackRate() time.Duration {
	return p.meleeAttackRate(static.EquipmentSlotOffHand)
}

func (p *Player) ResolveMainHandAttack(target interfaces.Unit) *interfaces.AttackInfo {
	minDamage, maxDamage := p.getWeaponDamageRange(GetObjectManager().GetItem(p.Equipment[static.EquipmentSlotMainHand]))
	return &interfaces.AttackInfo{
		Damage: minDamage + rand.Intn(maxDamage-minDamage+1),
	}
}

func (p *Player) ResolveOffHandAttack(target interfaces.Unit) *interfaces.AttackInfo {
	minDamage, maxDamage := p.getWeaponDamageRange(GetObjectManager().GetItem(p.Equipment[static.EquipmentSlotMainHand]))
	if minDamage == 0 && maxDamage == 0 {
		return nil
	}

	return &interfaces.AttackInfo{
		Damage: minDamage + rand.Intn(maxDamage-minDamage+1),
	}
}

func (p *Player) SetInCombat(inCombat bool) {
	p.Unit.InCombat = true
}

// Utility methods.
// getWeaponRamageRange returns a pair of numbers: (min, max) for the given weapon.
func (p *Player) getWeaponDamageRange(weapon *Item) (int, int) {
	if weapon == nil {
		return 0, 0
	}

	return int(weapon.GetTemplate().Damages[static.SpellSchoolPhysical].Min), int(weapon.GetTemplate().Damages[static.SpellSchoolPhysical].Max)
}

// meleeAttackRate calculates the attack rate for a given equipment slot.
func (p *Player) meleeAttackRate(slot static.EquipmentSlot) time.Duration {
	weapon := GetObjectManager().GetItem(p.Equipment[slot])
	if weapon == nil {
		if slot == static.EquipmentSlotMainHand {
			return time.Duration(2000) * time.Millisecond
		} else {
			return time.Duration(0)
		}
	}

	return weapon.GetTemplate().AttackRate
}

// Resistances returns the final resitances for the player, after
// all modifications.
func (p *Player) Resistances() map[static.SpellSchool]int {
	resistances := map[static.SpellSchool]int{
		static.SpellSchoolPhysical: 0,
		static.SpellSchoolHoly:     0,
		static.SpellSchoolFire:     0,
		static.SpellSchoolNature:   0,
		static.SpellSchoolFrost:    0,
		static.SpellSchoolShadow:   0,
		static.SpellSchoolArcane:   0,
	}

	/// Add modifications based on items.
	for _, itemGUID := range p.Equipment {
		item := GetObjectManager().GetItem(itemGUID)

		for k, v := range item.GetTemplate().Resistances {
			resistances[k] += v
		}
	}

	/// Add modifications based on stats.
	// Each point in agility gives 2 armor.
	resistances[static.SpellSchoolPhysical] += p.Agility * 2

	// Each point in spirit increases resistances by 0.05.
	spiritBonus := int(0.05 * float32(p.Spirit))
	resistances[static.SpellSchoolHoly] += spiritBonus
	resistances[static.SpellSchoolFire] += spiritBonus
	resistances[static.SpellSchoolNature] += spiritBonus
	resistances[static.SpellSchoolFrost] += spiritBonus
	resistances[static.SpellSchoolShadow] += spiritBonus
	resistances[static.SpellSchoolArcane] += spiritBonus

	return resistances
}

func (p *Player) meleeAttackPower() int {
	return p.Unit.meleeAttackPower() + p.meleeAttackPowerMods()
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
