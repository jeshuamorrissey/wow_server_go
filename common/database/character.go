package database

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
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

	GUID object.GUID

	// Flags.
	HideHelm        bool
	HideCloak       bool
	IsGhost         bool
	RenameNextLogin bool
}

// NewCharacter creates a new character entry in the database and
// returns a pointer to it.
func NewCharacter(
	om *object.Manager,
	account *Account, realm *Realm,
	name string,
	race c.Race, class c.Class, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) (*Character, error) {
	startingEquipment, startingItems := dbc.GetStartingItems(class, race)

	equipment := make(map[c.EquipmentSlot]object.GUID)
	for slot, item := range startingEquipment {
		itemObj := &object.Item{
			GameObject: object.GameObject{
				Entry: uint32(item.Entry),
			},
		}

		err := om.Add(itemObj)
		if err != nil {
			return nil, err
		}

		equipment[slot] = itemObj.GUID()
	}

	inventory := make(map[int]object.GUID)
	for i, item := range startingItems {
		itemObj := &object.Item{
			GameObject: object.GameObject{
				Entry: uint32(item.Entry),
			},
		}

		err := om.Add(itemObj)
		if err != nil {
			return nil, err
		}

		inventory[i] = itemObj.GUID()
	}

	startingLocation := dbc.GetStartingLocation(class, race)

	charObj := &object.Player{
		Unit: object.Unit{
			GameObject: object.GameObject{
				Entry:  0,
				ScaleX: dbc.GetPlayerScale(race, gender),
			},

			Loc: object.Location{
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
	}

	err := om.Add(charObj)
	if err != nil {
		return nil, err
	}

	return &Character{
		Name:      name,
		GUID:      charObj.GUID(),
		AccountID: account.ID,
		RealmID:   realm.ID,
	}, nil
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
