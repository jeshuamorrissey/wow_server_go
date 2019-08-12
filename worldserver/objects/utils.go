package objects

import (
	"math"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GUID represents the unique ID of an object within the manager.
type GUID uint64

// High returns the high-part of the GUID. This correspondings to a HighGUID
// value.
func (guid GUID) High() uint32 {
	return uint32(guid >> 32)
}

// Low returns the low-part of the GUID. This is the part that can be modified
// by the server.
func (guid GUID) Low() uint32 {
	return uint32(guid)
}

// Location represents a location in 3D space.
type Location struct {
	X, Y, Z, O float64
}

// Distance calculates the distance between two locations.
func (loc *Location) Distance(other *Location) float64 {
	return math.Sqrt(
		math.Pow(loc.X-other.X, 2) +
			math.Pow(loc.X-other.X, 2) +
			math.Pow(loc.X-other.X, 2))
}

// TypeID returns the type ID for a given game object.
func TypeID(obj GameObject) c.TypeID {
	switch obj.(type) {
	case *Item:
		return c.TypeIDItem
	case *Container:
		return c.TypeIDContainer
	case *Unit:
		return c.TypeIDUnit
	case *Player:
		return c.TypeIDPlayer
	}

	return c.TypeIDObject
}

// TypeMask returns the type mask for a given game object.
func TypeMask(obj GameObject) c.TypeMask {
	switch obj.(type) {
	case *Item:
		return c.TypeMaskObject | c.TypeMaskItem
	case *Container:
		return c.TypeMaskObject | c.TypeMaskItem | c.TypeMaskContainer
	case *Unit:
		return c.TypeMaskObject | c.TypeMaskUnit
	case *Player:
		return c.TypeMaskObject | c.TypeMaskUnit | c.TypeMaskPlayer
	}

	return c.TypeMaskObject
}

// UpdateFlags returns the update flags for a given game object.
func UpdateFlags(obj GameObject) c.UpdateFlags {
	switch obj.(type) {
	case *Item:
		return c.UpdateFlagsAll
	case *Container:
		return c.UpdateFlagsAll
	case *Unit:
		return c.UpdateFlagsAll | c.UpdateFlagsLiving | c.UpdateFlagsHasPosition
	case *Player:
		return c.UpdateFlagsAll | c.UpdateFlagsLiving | c.UpdateFlagsHasPosition
	}

	return c.UpdateFlagsNone
}

func mergeUpdateFields(first, second map[c.UpdateField]interface{}) map[c.UpdateField]interface{} {
	for k, v := range second {
		first[k] = v
	}

	return first
}
