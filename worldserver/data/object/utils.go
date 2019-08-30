package object

import (
	"fmt"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// HighGUID returns the high GUID for a given type of object. Will
// panic if an unknown type is passed.
func HighGUID(obj Object) c.HighGUID {
	switch o := obj.(type) {
	case *Item:
		return c.HighGUIDItem
	case *Container:
		return c.HighGUIDContainer
	case *Unit:
		return c.HighGUIDUnit
	case *Player:
		return c.HighGUIDPlayer
	default:
		panic(fmt.Sprintf("HighGUID: Unknown object type %T!", o))
	}
}

// TypeID returns the type ID for a given game object.
func TypeID(obj Object) c.TypeID {
	switch o := obj.(type) {
	case *Item:
		return c.TypeIDItem
	case *Container:
		return c.TypeIDContainer
	case *Unit:
		return c.TypeIDUnit
	case *Player:
		return c.TypeIDPlayer
	default:
		panic(fmt.Sprintf("TypeID: Unknown object type %T!", o))
	}
}

// TypeMask returns the type mask for a given game object.
func TypeMask(obj Object) c.TypeMask {
	switch o := obj.(type) {
	case *Item:
		return c.TypeMaskObject | c.TypeMaskItem
	case *Container:
		return c.TypeMaskObject | c.TypeMaskItem | c.TypeMaskContainer
	case *Unit:
		return c.TypeMaskObject | c.TypeMaskUnit
	case *Player:
		return c.TypeMaskObject | c.TypeMaskUnit | c.TypeMaskPlayer
	default:
		panic(fmt.Sprintf("TypeMask: Unknown object type %T!", o))
	}
}

// UpdateFlags returns the update flags for a given game object.
func UpdateFlags(obj Object) c.UpdateFlags {
	switch o := obj.(type) {
	case *Item:
		return c.UpdateFlagsAll
	case *Container:
		return c.UpdateFlagsAll
	case *Unit:
		return c.UpdateFlagsAll | c.UpdateFlagsLiving | c.UpdateFlagsHasPosition
	case *Player:
		return c.UpdateFlagsAll | c.UpdateFlagsLiving | c.UpdateFlagsHasPosition
	default:
		panic(fmt.Sprintf("UpdateFlags: Unknown object type %T!", o))
	}
}

// NumUpdateFields returns the number of bytes in the mask for an object type.
func NumUpdateFields(obj Object) int {
	switch o := obj.(type) {
	case *Item:
		return 48
	case *Container:
		return 106
	case *Unit:
		return 188
	case *Player:
		return 1282
	default:
		panic(fmt.Sprintf("UpdateFlags: Unknown object type %T!", o))
	}
}
