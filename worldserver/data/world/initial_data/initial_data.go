package initial_data

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
)

// NewCharacter creates a new character entry in the database and
// returns a pointer to it.
func NewCharacter(
	config *world.WorldConfig,
	name string,
	race *dbc.Race, class *dbc.Class, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) (*world.Character, error) {
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

		err := config.ObjectManager.Add(itemObj)
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

		err := config.ObjectManager.Add(itemObj)
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

			BaseHealth: startingStats.BaseHealth,
			Strength:   startingStats.Strength,
			Agility:    startingStats.Agility,
			Stamina:    startingStats.Stamina,
			Intellect:  startingStats.Intellect,
			Spirit:     startingStats.Spirit,

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

		Money: 10000,

		ZoneID: startingLocation.Zone,
		MapID:  startingLocation.Map,

		Equipment: equipment,
		Inventory: inventory,
	}

	err := config.ObjectManager.Add(charObj)
	if err != nil {
		return nil, err
	}

	for _, guid := range equipment {
		if config.ObjectManager.Exists(guid) {
			config.ObjectManager.Get(guid).(*object.Item).Owner = charObj.GUID()
			config.ObjectManager.Get(guid).(*object.Item).Container = charObj.GUID()
		}
	}

	for _, guid := range inventory {
		if config.ObjectManager.Exists(guid) {
			config.ObjectManager.Get(guid).(*object.Item).Owner = charObj.GUID()
			config.ObjectManager.Get(guid).(*object.Item).Container = charObj.GUID()
		}
	}

	return &world.Character{
		Name: name,
		GUID: charObj.GUID(),
	}, nil
}

func PopulateWorld(config *world.WorldConfig) error {
	err := config.ObjectManager.Add(&object.Unit{
		GameObject: object.GameObject{
			Entry:  uint32(dbc.UnitsByName["The Man"].Entry),
			ScaleX: 1.0,
		},

		Level:  1,
		Race:   dbc.RaceHuman,
		Class:  dbc.ClassRogue,
		Gender: c.GenderMale,

		HealthPercent: 1.0,
		PowerPercent:  1.0,

		MovementInfo: object.MovementInfo{
			Location: object.Location{
				X: -8945.95,
				Y: -132.493,
				Z: 83.5312,
				O: 180.0,
			},
		},
	})

	if err != nil {
		return err
	}

	err = config.ObjectManager.Add(&object.Unit{
		GameObject: object.GameObject{
			Entry:  uint32(dbc.UnitsByName["The Man"].Entry),
			ScaleX: 1.0,
		},

		Level:  1,
		Race:   dbc.RaceHuman,
		Class:  dbc.ClassRogue,
		Gender: c.GenderMale,

		HealthPercent: 1.0,
		PowerPercent:  1.0,

		MovementInfo: object.MovementInfo{
			Location: object.Location{
				X: -8942.95,
				Y: -132.493,
				Z: 83.5312,
				O: 180.0,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
