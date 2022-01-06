package static

var (
	startingItems = map[string]map[string][]string{
		"Warrior": {
			"Human": {
				"Recruit's Boots",
				"Recruit's Pants",
				"Recruit's Shirt",
				"Worn Shortsword",
				"Worn Wooden Shield",
			},
		},
	}
)

// GetStartingItems is a utility which will return a mapping of equipment slot
// to the item that should be in that slot.
func GetStartingItems(class *Class, race *Race) (map[EquipmentSlot]*Item, []*Item) {
	// items := startingItems[fmt.Sprintf("%d:%d", class, race)]
	items := startingItems[class.Name][race.Name]
	equipment := make(map[EquipmentSlot]*Item)
	nonEquipment := make([]*Item, 0)

	for _, itemID := range items {
		item := ItemsByName[itemID]

		if item.InventoryType == InventoryTypeHead {
			equipment[EquipmentSlotHead] = item
		} else if item.InventoryType == InventoryTypeShoulders {
			equipment[EquipmentSlotShoulders] = item
		} else if item.InventoryType == InventoryTypeBody {
			equipment[EquipmentSlotBody] = item
		} else if item.InventoryType == InventoryTypeChest || item.InventoryType == InventoryTypeRobe {
			equipment[EquipmentSlotChest] = item
		} else if item.InventoryType == InventoryTypeWaist {
			equipment[EquipmentSlotWaist] = item
		} else if item.InventoryType == InventoryTypeLegs {
			equipment[EquipmentSlotLegs] = item
		} else if item.InventoryType == InventoryTypeFeet {
			equipment[EquipmentSlotFeet] = item
		} else if item.InventoryType == InventoryTypeWrists {
			equipment[EquipmentSlotWrists] = item
		} else if item.InventoryType == InventoryTypeHands {
			equipment[EquipmentSlotHands] = item
		} else if item.InventoryType == InventoryTypeWeapon || item.InventoryType == InventoryType2HWeapon || item.InventoryType == InventoryTypeWeaponMainHand {
			equipment[EquipmentSlotMainHand] = item
		} else if item.InventoryType == InventoryTypeShield || item.InventoryType == InventoryTypeWeaponOffHand {
			equipment[EquipmentSlotOffHand] = item
		} else if item.InventoryType == InventoryTypeThrown || item.InventoryType == InventoryTypeRanged || item.InventoryType == InventoryTypeRangedRight {
			equipment[EquipmentSlotRanged] = item
		} else if item.InventoryType == InventoryTypeTabard {
			equipment[EquipmentSlotTabard] = item
		} else if item.InventoryType == InventoryTypeCloak {
			equipment[EquipmentSlotBack] = item
		} else {
			nonEquipment = append(nonEquipment, item)
		}
	}

	return equipment, nonEquipment
}
