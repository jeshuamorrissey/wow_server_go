package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Unit interface methods (game-logic).
func (p *Player) Initialize() {
	go p.restoreHealthPower()
}

// Utility methods.
func (p *Player) restoreHealthPower() {
	for range time.Tick(static.RegenTimeout) {
		if p.IsLoggedIn {
			p.SendUpdates([]interface{}{
				&messages.ModHealth{Amount: p.HealthRegen(static.RegenTimeout, p.IsInCombat())},
				&messages.ModPower{Amount: p.PowerRegen(static.RegenTimeout, p.IsInCombat())},
			})
		}
	}
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

func (p *Player) resistances() map[static.SpellSchool]int {
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

func (p *Player) weapons() []*Item {
	weapons := make([]*Item, 0)
	for _, slot := range []static.EquipmentSlot{static.EquipmentSlotMainHand, static.EquipmentSlotOffHand, static.EquipmentSlotRanged} {
		weaponGUID, ok := p.Equipment[slot]
		if ok {
			if weapon := GetObjectManager().GetItem(weaponGUID); weapon != nil {
				if weapon.GetTemplate().AttackRate > 0 {
					weapons = append(weapons, weapon)
				}
			}
		}
	}

	return weapons
}
