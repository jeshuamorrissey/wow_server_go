package initial_data

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/components"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// NewCharacter creates a new character entry in the database and
// returns a pointer to it.
func NewCharacter(
	cfg *config.Config,
	name string,
	race *static.Race, class *static.Class, gender static.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) (*config.Character, error) {
	startingEquipment, startingItems := static.GetStartingItems(class, race)

	equipment := make(map[static.EquipmentSlot]interfaces.GUID)
	for slot, item := range startingEquipment {
		itemObj := &dynamic.Item{
			GameObject: dynamic.GameObject{
				Entry:  uint32(item.Entry),
				ScaleX: 1.0,
			},

			Durability: item.MaxDurability,
		}

		cfg.ObjectManager.Add(itemObj)
		equipment[slot] = itemObj.GUID()
	}

	inventory := make(map[int]interfaces.GUID)
	for i, item := range startingItems {
		itemObj := &dynamic.Item{
			GameObject: dynamic.GameObject{
				Entry:  uint32(item.Entry),
				ScaleX: 1.0,
			},

			Durability: item.MaxDurability,
		}

		cfg.ObjectManager.Add(itemObj)
		inventory[i] = itemObj.GUID()
	}

	startingLocation := static.StartingLocationsByIndex[race]
	startingStats := static.StartingStatsByIndex[class][race]

	charObj := dynamic.InitializePlayer(&dynamic.Player{
		GameObject: dynamic.GameObject{
			Entry:  0,
			ScaleX: static.GetPlayerScale(race, gender),
		},

		Movement: components.Movement{
			MovementInfo: interfaces.MovementInfo{
				Location: interfaces.Location{
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
		},

		Unit: components.Unit{
			Level:  1,
			Race:   race,
			Class:  class,
			Gender: gender,
		},

		BasicStats: components.BasicStats{
			Strength:  startingStats.Strength,
			Agility:   startingStats.Agility,
			Stamina:   startingStats.Stamina,
			Intellect: startingStats.Intellect,
			Spirit:    startingStats.Spirit,
		},

		PlayerFeatures: components.PlayerFeatures{
			SkinColor: int(skinColor),
			Face:      int(face),
			HairStyle: int(hairStyle),
			HairColor: int(hairColor),
			Feature:   int(feature),
		},

		Player: components.Player{
			Money: 10000,
		},

		ZoneID: startingLocation.Zone,
		MapID:  startingLocation.Map,

		Equipment: equipment,
		Inventory: inventory,
	})

	cfg.ObjectManager.Add(charObj)
	for _, guid := range equipment {
		if cfg.ObjectManager.Exists(guid) {
			cfg.ObjectManager.GetItem(guid).Owner = charObj.GUID()
			cfg.ObjectManager.GetItem(guid).Container = charObj.GUID()
		}
	}

	for _, guid := range inventory {
		if cfg.ObjectManager.Exists(guid) {
			cfg.ObjectManager.GetItem(guid).Owner = charObj.GUID()
			cfg.ObjectManager.GetItem(guid).Container = charObj.GUID()
		}
	}

	return &config.Character{
		Name: name,
		GUID: charObj.GUID(),
	}, nil
}

func PopulateWorld(cfg *config.Config) error {
	cfg.ObjectManager.Add(dynamic.InitializeUnit(&dynamic.Unit{
		GameObject: dynamic.GameObject{
			Entry:  uint32(static.UnitsByName["The Man"].Entry),
			ScaleX: 1.0,
		},

		Unit: components.Unit{
			Level:  1,
			Race:   static.RaceHuman,
			Class:  static.ClassRogue,
			Gender: static.GenderMale,
		},

		Movement: components.Movement{
			MovementInfo: interfaces.MovementInfo{
				Location: interfaces.Location{
					X: -8945.95,
					Y: -132.493,
					Z: 83.5312,
					O: 180.0,
				},
			},

			SpeedWalk:         2.5,
			SpeedRun:          7.0,
			SpeedRunBackward:  4.5,
			SpeedSwim:         4.72,
			SpeedSwimBackward: 2.5,
			SpeedTurn:         3.14159,
		},

		RespawnTimeMS: 1000 * time.Millisecond,
	}))

	// cfg.ObjectManager.Add(dynamic.InitializeUnit(&dynamic.Unit{
	// 	GameObject: dynamic.GameObject{
	// 		Entry:  uint32(static.UnitsByName["The Man"].Entry),
	// 		ScaleX: 1.0,
	// 	},

	// 	Unit: components.Unit{
	// 		Level:  1,
	// 		Race:   static.RaceHuman,
	// 		Class:  static.ClassRogue,
	// 		Gender: static.GenderMale,
	// 	},

	// 	Movement: components.Movement{
	// 		MovementInfo: interfaces.MovementInfo{
	// 			Location: interfaces.Location{
	// 				X: -8942.95,
	// 				Y: -132.493,
	// 				Z: 83.5312,
	// 				O: 180.0,
	// 			},
	// 		},

	// 		SpeedWalk:         2.5,
	// 		SpeedRun:          7.0,
	// 		SpeedRunBackward:  4.5,
	// 		SpeedSwim:         4.72,
	// 		SpeedSwimBackward: 2.5,
	// 		SpeedTurn:         3.14159,
	// 	},
	// }))

	return nil
}
