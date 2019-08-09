package data

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GetStartingEquipment is a utility which will return a mapping of equipment slot
// to the item that should be in that slot.
func GetStartingEquipment(class c.Class, race c.Race) map[c.EquipmentSlot]*Item {
	if class == c.ClassWarrior {
		if race == c.RaceHuman {
			return map[c.EquipmentSlot]*Item{
				c.EquipmentSlotChest:    Items[38],   // Recruit's Shirt
				c.EquipmentSlotLegs:     Items[39],   // Recruit's Pants
				c.EquipmentSlotFeet:     Items[40],   // Recruit's Boots
				c.EquipmentSlotMainHand: Items[25],   // Worn Shortsword
				c.EquipmentSlotOffHand:  Items[2362], // Worn Wooden Shield
			}
		}
	}

	return map[c.EquipmentSlot]*Item{}
}

// GetStartingItems is a utility which will return pointers to the item
// templates that a certain race/class combination should have.
func GetStartingItems(class c.Class, race c.Race) []*Item {
	items := []*Item{Items[6948]}
	if race == c.RaceHuman {
		items = append(items, Items[4540])
	}
	return items
}
