package object

import c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"

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
