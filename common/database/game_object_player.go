package database

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jinzhu/gorm"
)

// EquippedItem is a mapping of EquipmentSlot --> GameObjectItem.
type EquippedItem struct {
	gorm.Model

	Slot c.EquipmentSlot
	Item *GameObjectItem

	GameObjectPlayerID uint
}

// GameObjectPlayer is a game object which represents a
// player-controlled unit.
type GameObjectPlayer struct {
	GameObjectUnit

	Level uint8

	SkinColor uint8
	Face      uint8
	HairStyle uint8
	HairColor uint8
	Feature   uint8

	ZoneID uint32
	MapID  uint32

	Equipment []*EquippedItem
	Inventory []*BaggedItem
	Bags      []*GameObjectContainer

	CharacterID uint
}

// GetTypeID returns the type of the game object.
func (obj *GameObjectPlayer) GetTypeID() c.TypeID { return c.TypeIDPlayer }

// GetTypeMask returns the type mask for the game object.
func (obj *GameObjectPlayer) GetTypeMask() c.TypeMask {
	return obj.GameObjectUnit.GetTypeMask() | c.TypeMaskPlayer
}

// GetUpdateFlags returns the common update flags for the given type of game object.
func (obj *GameObjectPlayer) GetUpdateFlags() c.UpdateFlags {
	return obj.GameObjectUnit.GetUpdateFlags()
}

// GetHighGUID returns the high GUID component for the game object.
func (obj *GameObjectPlayer) GetHighGUID() c.HighGUID { return c.HighGUIDPlayer }

// EquipmentMap returns a map of equipment slot --> item.
func (obj *GameObjectPlayer) EquipmentMap() map[c.EquipmentSlot]*GameObjectItem {
	result := make(map[c.EquipmentSlot]*GameObjectItem)
	for _, item := range obj.Equipment {
		result[item.Slot] = item.Item
	}

	return result
}
