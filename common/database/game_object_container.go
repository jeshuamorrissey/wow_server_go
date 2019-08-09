package database

import (
	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jinzhu/gorm"
)

// BaggedItem represents an item within a container.
type BaggedItem struct {
	gorm.Model

	Slot int
	Item *GameObjectItem

	GameObjectPlayerID    uint // for items in the inventory
	GameObjectContainerID uint // for items in a bag
}

// GameObjectContainer represents an item instance in the game.
type GameObjectContainer struct {
	GameObjectItem

	Items []*BaggedItem
}

// GetTypeID returns the type of the game object.
func (obj *GameObjectContainer) GetTypeID() c.TypeID { return c.TypeIDContainer }

// GetTypeMask returns the type mask for the game object.
func (obj *GameObjectContainer) GetTypeMask() c.TypeMask {
	return obj.GameObjectItem.GetTypeMask() | c.TypeMaskContainer
}

// GetUpdateFlags returns the common update flags for the given type of game object.
func (obj *GameObjectContainer) GetUpdateFlags() c.UpdateFlags { return c.UpdateFlagsAll }

// GetHighGUID returns the high GUID component for the game object.
func (obj *GameObjectContainer) GetHighGUID() c.HighGUID { return c.HighGUIDContainer }

// Template returns a pointer to the item template for this item.
func (obj *GameObjectContainer) Template() *data.Item {
	return data.Items[obj.Entry]
}
