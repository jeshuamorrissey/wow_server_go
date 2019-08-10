package data

import (
	"encoding/json"
	"fmt"
	"os"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

var (
	// Map of "<Class>:<Race>" --> list of item entries.
	startingItems map[string][]int
)

// LoadStartingItems reads the starting item JSON file and
// populates the startingItems map.
func LoadStartingItems(jsonFile string) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(&startingItems)
}

// GetStartingItems is a utility which will return a mapping of equipment slot
// to the item that should be in that slot.
func GetStartingItems(class c.Class, race c.Race) (map[c.EquipmentSlot]*Item, []*Item) {
	items := startingItems[fmt.Sprintf("%d:%d", class, race)]
	equipment := make(map[c.EquipmentSlot]*Item)
	nonEquipment := make([]*Item, 0)

	for _, itemID := range items {
		item := Items[itemID]

		if item.InventoryType == c.InventoryTypeHead {
			equipment[c.EquipmentSlotHead] = item
		} else if item.InventoryType == c.InventoryTypeShoulders {
			equipment[c.EquipmentSlotShoulders] = item
		} else if item.InventoryType == c.InventoryTypeBody {
			equipment[c.EquipmentSlotBody] = item
		} else if item.InventoryType == c.InventoryTypeChest || item.InventoryType == c.InventoryTypeRobe {
			equipment[c.EquipmentSlotChest] = item
		} else if item.InventoryType == c.InventoryTypeWaist {
			equipment[c.EquipmentSlotWaist] = item
		} else if item.InventoryType == c.InventoryTypeLegs {
			equipment[c.EquipmentSlotLegs] = item
		} else if item.InventoryType == c.InventoryTypeFeet {
			equipment[c.EquipmentSlotFeet] = item
		} else if item.InventoryType == c.InventoryTypeWrists {
			equipment[c.EquipmentSlotWrists] = item
		} else if item.InventoryType == c.InventoryTypeHands {
			equipment[c.EquipmentSlotHands] = item
		} else if item.InventoryType == c.InventoryTypeWeapon || item.InventoryType == c.InventoryType2HWeapon || item.InventoryType == c.InventoryTypeWeaponMainHand {
			equipment[c.EquipmentSlotMainHand] = item
		} else if item.InventoryType == c.InventoryTypeShield || item.InventoryType == c.InventoryTypeWeaponOffHand {
			equipment[c.EquipmentSlotOffHand] = item
		} else if item.InventoryType == c.InventoryTypeThrown || item.InventoryType == c.InventoryTypeRanged || item.InventoryType == c.InventoryTypeRangedRight {
			equipment[c.EquipmentSlotRanged] = item
		} else if item.InventoryType == c.InventoryTypeTabard {
			equipment[c.EquipmentSlotTabard] = item
		} else if item.InventoryType == c.InventoryTypeCloak {
			equipment[c.EquipmentSlotBack] = item
		} else {
			nonEquipment = append(nonEquipment, item)
		}
	}

	return equipment, nonEquipment
}
