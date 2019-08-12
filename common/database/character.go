package database

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/objects"
	"github.com/jinzhu/gorm"
)

// Character represents a character in the game, linked to an account.
// The character's information is stored over this structure and a
// game object.
type Character struct {
	gorm.Model

	Name      string `gorm:"unique"`
	LastLogin *time.Time

	AccountID uint
	RealmID   uint

	GUID objects.GUID

	// Flags.
	HideHelm        bool
	HideCloak       bool
	IsGhost         bool
	RenameNextLogin bool
}

// Flags returns an set of flags based on the character's state.
func (char *Character) Flags() uint32 {
	var flags uint32
	if char.HideHelm {
		flags |= uint32(c.CharacterFlagHideHelm)
	}

	if char.HideCloak {
		flags |= uint32(c.CharacterFlagHideCloak)
	}

	if char.IsGhost {
		flags |= uint32(c.CharacterFlagGhost)
	}

	if char.RenameNextLogin {
		flags |= uint32(c.CharacterFlagRename)
	}

	return flags
}

// NewCharacter makes a new character with some basic information.
func NewCharacter(
	om *objects.ObjectManager,
	name string,
	account *Account, realm *Realm,
	class c.Class, race c.Race, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) *Character {
	startingEquipment, startingItems := data.GetStartingItems(class, race)

	equipment := make(map[c.EquipmentSlot]*objects.Item)
	for slot, item := range startingEquipment {
		equipment[slot] = om.Create(&objects.Item{
			BaseGameObject: objects.BaseGameObject{
				Entry: item.Entry,
			},
		}).(*objects.Item)
	}

	inventory := make(map[int]*objects.Item)
	for i, item := range startingItems {
		inventory[i] = om.Create(&objects.Item{
			BaseGameObject: objects.BaseGameObject{
				Entry: item.Entry,
			},
		}).(*objects.Item)
	}

	startingLocation := data.GetStartingLocation(class, race)

	return &Character{
		Name: name,
		GUID: om.Create(&objects.Player{
			Unit: objects.Unit{
				BaseGameObject: objects.BaseGameObject{
					Entry:  0,
					ScaleX: 1.0,
				},

				Location: objects.Location{
					X: startingLocation.X,
					Y: startingLocation.Y,
					Z: startingLocation.Z,
					O: startingLocation.O,
				},

				Level:  1,
				Race:   race,
				Class:  class,
				Gender: gender,
			},

			SkinColor: skinColor,
			Face:      face,
			HairStyle: hairStyle,
			HairColor: hairColor,
			Feature:   feature,

			ZoneID: startingLocation.Zone,
			MapID:  startingLocation.Map,

			Equipment: equipment,
			Inventory: inventory,
		}).GUID(),
		AccountID: account.ID,
		RealmID:   realm.ID,
	}
}
