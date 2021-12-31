package dbc

import c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"

// Class represents data within the ChrClasses.dbc file.
type Class struct {
	ID          int
	PrimaryStat int
	PowerType   c.Power
	PetType     string
	Name        string
}

var (
	// ClassByID is the primary source of truth, storing data for for this DBC.
	ClassByID map[int]*Class
)

// Indexes for this DBC file, to make querying easier.
var (
	ClassByIndex map[string]*Class

	ClassWarrior *Class
	ClassPaladin *Class
	ClassHunter  *Class
	ClassRogue   *Class
	ClassPriest  *Class
	ClassShaman  *Class
	ClassMage    *Class
	ClassWarlock *Class
	ClassDruid   *Class
)

func init() {
	// Set the source of truth.
	ClassByID = map[int]*Class{
		1: {
			ID:          1,
			PrimaryStat: 0,
			PowerType:   c.PowerRage,
			PetType:     "PET",
			Name:        "Warrior",
		},
		2: {
			ID:          2,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Paladin",
		},
		3: {
			ID:          3,
			PrimaryStat: 1,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Hunter",
		},
		4: {
			ID:          4,
			PrimaryStat: 1,
			PowerType:   c.PowerEnergy,
			PetType:     "PET",
			Name:        "Rogue",
		},
		5: {
			ID:          5,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Priest",
		},
		7: {
			ID:          7,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Shaman",
		},
		8: {
			ID:          8,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Mage",
		},
		9: {
			ID:          9,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "DEMON",
			Name:        "Warlock",
		},
		11: {
			ID:          11,
			PrimaryStat: 0,
			PowerType:   c.PowerMana,
			PetType:     "PET",
			Name:        "Druid",
		},
	}

	// Set the index.
	ClassByIndex = make(map[string]*Class)

	// Set the index values.
	ClassByIndex["Warrior"] = ClassByID[1]
	ClassByIndex["Paladin"] = ClassByID[2]
	ClassByIndex["Hunter"] = ClassByID[3]
	ClassByIndex["Rogue"] = ClassByID[4]
	ClassByIndex["Priest"] = ClassByID[5]
	ClassByIndex["Shaman"] = ClassByID[7]
	ClassByIndex["Mage"] = ClassByID[8]
	ClassByIndex["Warlock"] = ClassByID[9]
	ClassByIndex["Druid"] = ClassByID[11]

	// As there is only a single index, add some special convenience values.
	ClassWarrior = ClassByID[1]
	ClassPaladin = ClassByID[2]
	ClassHunter = ClassByID[3]
	ClassRogue = ClassByID[4]
	ClassPriest = ClassByID[5]
	ClassShaman = ClassByID[7]
	ClassMage = ClassByID[8]
	ClassWarlock = ClassByID[9]
	ClassDruid = ClassByID[11]
}
