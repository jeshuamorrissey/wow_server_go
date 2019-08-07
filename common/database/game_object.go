package database

import (
	"github.com/jinzhu/gorm"
)

// TypeID is a enum value which maps to a object's type ID.
//go:generate stringer -type=TypeID -trimprefix=TypeID
type TypeID int

// TypeID values.
const (
	TypeIDObject        TypeID = 0
	TypeIDItem          TypeID = 1
	TypeIDContainer     TypeID = 2
	TypeIDUnit          TypeID = 3
	TypeIDPlayer        TypeID = 4
	TypeIDGameObject    TypeID = 5
	TypeIDDynamicObject TypeID = 6
	TypeIDCorpse        TypeID = 7
)

// TypeMask is a enum value which maps to a object's type ID.
//go:generate stringer -type=TypeMask -trimprefix=TypeMask
type TypeMask int

// TypeMask values.
const (
	TypeMaskObject        TypeMask = 0x0001
	TypeMaskItem          TypeMask = 0x0002
	TypeMaskContainer     TypeMask = 0x0004
	TypeMaskUnit          TypeMask = 0x0008
	TypeMaskPlayer        TypeMask = 0x0010
	TypeMaskGameObject    TypeMask = 0x0020
	TypeMaskDynamicObject TypeMask = 0x0040
	TypeMaskCorpse        TypeMask = 0x0080
)

// UpdateFlags is a enum value which maps to a object's type ID.
//go:generate stringer -type=UpdateFlags -trimprefix=UpdateFlags
type UpdateFlags int

// UpdateFlags values.
const (
	UpdateFlagsNone        UpdateFlags = 0x0000
	UpdateFlagsSelf        UpdateFlags = 0x0001
	UpdateFlagsTransport   UpdateFlags = 0x0002
	UpdateFlagsFullGUID    UpdateFlags = 0x0004
	UpdateFlagsHighGUID    UpdateFlags = 0x0008
	UpdateFlagsAll         UpdateFlags = 0x0010
	UpdateFlagsLiving      UpdateFlags = 0x0020
	UpdateFlagsHasPosition UpdateFlags = 0x0040
)

// HighGUID is a enum value which maps to a object's type ID.
//go:generate stringer -type=HighGUID -trimprefix=HighGUID
type HighGUID uint32

// HighGUID values.
const (
	HighGUIDItem          HighGUID = 0x40000000
	HighGUIDContainer     HighGUID = 0x40000000
	HighGUIDPlayer        HighGUID = 0x00000000
	HighGUIDGameobject    HighGUID = 0xF1100000
	HighGUIDTransport     HighGUID = 0xF1200000
	HighGUIDUnit          HighGUID = 0xF1300000
	HighGUIDPet           HighGUID = 0xF1400000
	HighGUIDDynamicObject HighGUID = 0xF1000000
	HighGUIDCorpse        HighGUID = 0xF1010000
	HighGUIDMoTransport   HighGUID = 0x1FC00000
)

// GameObject represents a generic game object.
type GameObject interface {
	GetTypeID() TypeID
	GetTypeMask() TypeMask
	GetUpdateFlags() UpdateFlags
	GetHighGUID() HighGUID

	// GetFields() map[string]interface{}
}

// GameObjectBase represents the base, "generic" game object.
type GameObjectBase struct {
	gorm.Model
}

func (obj *GameObjectBase) GetTypeID() TypeID           { return TypeIDObject }
func (obj *GameObjectBase) GetTypeMask() TypeMask       { return TypeMaskObject }
func (obj *GameObjectBase) GetUpdateFlags() UpdateFlags { return UpdateFlagsNone }
func (obj *GameObjectBase) GetHighGUID() HighGUID       { return 0 }

// GameObjectUnit is a game object which represents a unit.
type GameObjectUnit struct {
	GameObjectBase

	Race   Race
	Class  Class
	Gender Gender

	X float32
	Y float32
	Z float32
	O float32
}

func (obj *GameObjectUnit) GetTypeID() TypeID { return TypeIDUnit }
func (obj *GameObjectUnit) GetTypeMask() TypeMask {
	return obj.GameObjectBase.GetTypeMask() | TypeMaskUnit
}
func (obj *GameObjectUnit) GetUpdateFlags() UpdateFlags {
	return UpdateFlagsAll | UpdateFlagsLiving | UpdateFlagsHasPosition
}
func (obj *GameObjectUnit) GetHighGUID() HighGUID { return HighGUIDUnit }

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

	CharacterID uint
}

func (obj *GameObjectPlayer) GetTypeID() TypeID { return TypeIDPlayer }
func (obj *GameObjectPlayer) GetTypeMask() TypeMask {
	return obj.GameObjectUnit.GetTypeMask() | TypeMaskPlayer
}
func (obj *GameObjectPlayer) GetUpdateFlags() UpdateFlags { return obj.GameObjectUnit.GetUpdateFlags() }
func (obj *GameObjectPlayer) GetHighGUID() HighGUID       { return HighGUIDPlayer }
