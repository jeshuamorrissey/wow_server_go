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
	race *dbc.Race, class *dbc.Class, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) (*Character, error) {
	startingEquipment, startingItems := dbc.GetStartingItems(class, race)

	equipment := make(map[c.EquipmentSlot]object.GUID)
	for slot, item := range startingEquipment {
		itemObj := &object.Item{
			GameObject: object.GameObject{
				Entry:  uint32(item.Entry),
				ScaleX: 1.0,
			},

			Durability: item.MaxDurability,
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
				Entry:  uint32(item.Entry),
				ScaleX: 1.0,
			},

			Durability: item.MaxDurability,
		}

		err := om.Add(itemObj)
		if err != nil {
			return nil, err
		}

		inventory[i] = itemObj.GUID()
	}

	startingLocation := dbc.StartingLocationsByIndex[race]
	startingStats := dbc.StartingStatsByIndex[class][race]

	charObj := &object.Player{
		Unit: object.Unit{
			GameObject: object.GameObject{
				Entry: 0,
				// ScaleX: dbc.GetPlayerScale(race, gender),
				ScaleX: 1.0,
			},

			MovementInfo: object.MovementInfo{
				Location: object.Location{
					X: startingLocation.X,
					Y: startingLocation.Y,
					Z: startingLocation.Z,
					O: startingLocation.O,
				},
			},

			SpeedWalk:         2.5,
			SpeedRun:          7.0,
			SpeedRunBackward:  4.5,
			SpeedSwim:         4.72,
			SpeedSwimBackward: 2.5,
			SpeedTurn:         3.14159,

			HealthPercent: 1.0,
			PowerPercent:  1.0,

			Strength: startingStats.Strength,
			// Agility:   startingStats.Agility,
			// Stamina:   startingStats.Stamina,
			// Intellect: startingStats.Intellect,
			// Spirit:    startingStats.Spirit,

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

	for _, guid := range equipment {
		if om.Exists(guid) {
			om.Get(guid).(*object.Item).Owner = charObj.GUID()
			om.Get(guid).(*object.Item).Container = charObj.GUID()
		}
	}

	for _, guid := range inventory {
		if om.Exists(guid) {
			om.Get(guid).(*object.Item).Owner = charObj.GUID()
			om.Get(guid).(*object.Item).Container = charObj.GUID()
		}
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
