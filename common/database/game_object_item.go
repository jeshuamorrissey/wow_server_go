package database

import (
	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GameObjectItem represents an item instance in the game.
type GameObjectItem struct {
	GameObjectBase

	EquippedItemID     uint // may be 0 for items which aren't equipment
	GameObjectPlayerID uint // the player this item belongs to
}

// GetTypeID returns the type of the game object.
func (obj *GameObjectItem) GetTypeID() c.TypeID { return c.TypeIDItem }

// GetTypeMask returns the type mask for the game object.
func (obj *GameObjectItem) GetTypeMask() c.TypeMask {
	return obj.GameObjectBase.GetTypeMask() | c.TypeMaskItem
}

// GetUpdateFlags returns the common update flags for the given type of game object.
func (obj *GameObjectItem) GetUpdateFlags() c.UpdateFlags { return c.UpdateFlagsAll }

// GetHighGUID returns the high GUID component for the game object.
func (obj *GameObjectItem) GetHighGUID() c.HighGUID { return c.HighGUIDItem }

// Template returns a pointer to the item template for this item.
func (obj *GameObjectItem) Template() *data.Item {
	return data.Items[obj.Entry]
}
