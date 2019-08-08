package database

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jinzhu/gorm"
)

// GameObject represents a generic game object.
type GameObject interface {
	GetTypeID() c.TypeID
	GetTypeMask() c.TypeMask
	GetUpdateFlags() c.UpdateFlags
	GetHighGUID() c.HighGUID

	// GetFields() map[string]interface{}
}

// GameObjectBase represents the base, "generic" game object.
type GameObjectBase struct {
	gorm.Model

	Entry int
}

// GetTypeID returns the type of the game object.
func (obj *GameObjectBase) GetTypeID() c.TypeID { return c.TypeIDObject }

// GetTypeMask returns the type mask for the game object.
func (obj *GameObjectBase) GetTypeMask() c.TypeMask { return c.TypeMaskObject }

// GetUpdateFlags returns the common update flags for the given type of game object.
func (obj *GameObjectBase) GetUpdateFlags() c.UpdateFlags { return c.UpdateFlagsNone }

// GetHighGUID returns the high GUID component for the game object.
func (obj *GameObjectBase) GetHighGUID() c.HighGUID { return 0 }
