package database

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
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

	Object GameObjectPlayer

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
	name string,
	account *Account, realm *Realm,
	class c.Class, race c.Race, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) *Character {
	equipment := []*EquippedItem{}
	for slot, item := range data.GetStartingEquipment(class, race) {
		equipment = append(equipment, &EquippedItem{
			Slot: slot,
			Item: &GameObjectItem{
				GameObjectBase: GameObjectBase{
					Entry: item.Entry,
				},
			},
		})
	}

	inventory := []*BaggedItem{}
	for i, item := range data.GetStartingItems(class, race) {
		inventory = append(inventory, &BaggedItem{
			Slot: i,
			Item: &GameObjectItem{
				GameObjectBase: GameObjectBase{
					Entry: item.Entry,
				},
			},
		})
	}

	return &Character{
		Name: name,
		Object: GameObjectPlayer{
			GameObjectUnit: GameObjectUnit{
				Race:   race,
				Class:  class,
				Gender: gender,

				// TODO(jeshua): make this based on the race.
				X: 0.0,
				Y: 0.0,
				Z: 0.0,
				O: 0.0,
			},

			Level: 1,

			SkinColor: skinColor,
			Face:      face,
			HairStyle: hairStyle,
			HairColor: hairColor,
			Feature:   feature,

			ZoneID: 1,
			MapID:  1,

			Equipment: equipment,
			Inventory: inventory,
			Bags:      []*GameObjectContainer{},
		},
		AccountID: account.ID,
		RealmID:   realm.ID,
	}
}
