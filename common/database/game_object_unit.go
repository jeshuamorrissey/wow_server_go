package database

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GameObjectUnit is a game object which represents a unit.
type GameObjectUnit struct {
	GameObjectBase

	Race   c.Race
	Class  c.Class
	Gender c.Gender

	X float32
	Y float32
	Z float32
	O float32
}

// GetTypeID returns the type of the game object.
func (obj *GameObjectUnit) GetTypeID() c.TypeID { return c.TypeIDUnit }

// GetTypeMask returns the type mask for the game object.
func (obj *GameObjectUnit) GetTypeMask() c.TypeMask {
	return obj.GameObjectBase.GetTypeMask() | c.TypeMaskUnit
}

// GetUpdateFlags returns the common update flags for the given type of game object.
func (obj *GameObjectUnit) GetUpdateFlags() c.UpdateFlags {
	return c.UpdateFlagsAll | c.UpdateFlagsLiving | c.UpdateFlagsHasPosition
}

// GetHighGUID returns the high GUID component for the game object.
func (obj *GameObjectUnit) GetHighGUID() c.HighGUID { return c.HighGUIDUnit }
