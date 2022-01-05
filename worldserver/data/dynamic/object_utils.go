package dynamic

import (
	"fmt"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
)

// HighGUID returns the high GUID for a given type of object. Will
// panic if an unknown type is passed.
func HighGUID(obj interfaces.Object) static.HighGUID {
	switch o := obj.(type) {
	case *Item:
		return static.HighGUIDItem
	case *Container:
		return static.HighGUIDContainer
	case *Unit:
		return static.HighGUIDUnit
	case *Player:
		return static.HighGUIDPlayer
	default:
		panic(fmt.Sprintf("HighGUID: Unknown object type %T!", o))
	}
}

// TypeID returns the type ID for a given game object.
func TypeID(obj interfaces.Object) static.TypeID {
	switch o := obj.(type) {
	case *Item:
		return static.TypeIDItem
	case *Container:
		return static.TypeIDContainer
	case *Unit:
		return static.TypeIDUnit
	case *Player:
		return static.TypeIDPlayer
	default:
		panic(fmt.Sprintf("TypeID: Unknown object type %T!", o))
	}
}

// TypeMask returns the type mask for a given game object.
func TypeMask(obj interfaces.Object) static.TypeMask {
	switch o := obj.(type) {
	case *Item:
		return static.TypeMaskObject | static.TypeMaskItem
	case *Container:
		return static.TypeMaskObject | static.TypeMaskItem | static.TypeMaskContainer
	case *Unit:
		return static.TypeMaskObject | static.TypeMaskUnit
	case *Player:
		return static.TypeMaskObject | static.TypeMaskUnit | static.TypeMaskPlayer
	default:
		panic(fmt.Sprintf("TypeMask: Unknown object type %T!", o))
	}
}

// UpdateFlags returns the update flags for a given game object.
func UpdateFlags(obj interfaces.Object) static.UpdateFlags {
	switch o := obj.(type) {
	case *Item:
		return static.UpdateFlagsAll
	case *Container:
		return static.UpdateFlagsAll
	case *Unit:
		return static.UpdateFlagsAll | static.UpdateFlagsLiving | static.UpdateFlagsHasPosition
	case *Player:
		return static.UpdateFlagsAll | static.UpdateFlagsLiving | static.UpdateFlagsHasPosition
	default:
		panic(fmt.Sprintf("UpdateFlags: Unknown object type %T!", o))
	}
}

// NumUpdateFields returns the number of bytes in the mask for an object type.
func NumUpdateFields(obj interfaces.Object) int {
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
