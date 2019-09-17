package dbc

// Race represents data within the ChrRaces.dbc file.
type Race struct {
	ID                    int
	Flags                 int
	FactionID             int
	Unk                   int
	MaleDisplayID         int
	FemaleDisplayID       int
	ClientPrefix          string
	MountScale            float32
	BaseLanguage          int
	CreatureType          int
	LoginEffectSpellID    int
	CombatStunSpellID     int
	ResSicknessSpellID    int
	SplashSoundID         int
	StartingTaxiNodes     int
	ClientFileString      string
	CinematicSequenceID   int
	Name                  string
	MaleFeatureName       string
	FemaleFeatureName     string
	HairCustomizationName string
}

var (
	// RaceByID is the primary source of truth, storing data for for this DBC.
	RaceByID map[int]*Race
)

// Indexes for this DBC file, to make querying easier.
var (
	RaceByIndex map[string]*Race

	RaceHuman    *Race
	RaceOrc      *Race
	RaceDwarf    *Race
	RaceNightElf *Race
	RaceUndead   *Race
	RaceTauren   *Race
	RaceGnome    *Race
	RaceTroll    *Race
	RaceGoblin   *Race
)

func init() {
	// Set the source of truth.
	RaceByID = map[int]*Race{
		1: &Race{
			ID:                    1,
			Flags:                 12,
			FactionID:             1,
			Unk:                   4140,
			MaleDisplayID:         49,
			FemaleDisplayID:       50,
			ClientPrefix:          "Hu",
			MountScale:            1.0,
			BaseLanguage:          7,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     2,
			ClientFileString:      "Human",
			CinematicSequenceID:   81,
			Name:                  "Human",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "PIERCINGS",
			HairCustomizationName: "NORMAL",
		},
		2: &Race{
			ID:                    2,
			Flags:                 12,
			FactionID:             2,
			Unk:                   4141,
			MaleDisplayID:         51,
			FemaleDisplayID:       52,
			ClientPrefix:          "Or",
			MountScale:            1.100000023841858,
			BaseLanguage:          1,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     4194304,
			ClientFileString:      "Orc",
			CinematicSequenceID:   21,
			Name:                  "Orc",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "PIERCINGS",
			HairCustomizationName: "NORMAL",
		},
		3: &Race{
			ID:                    3,
			Flags:                 12,
			FactionID:             3,
			Unk:                   4147,
			MaleDisplayID:         53,
			FemaleDisplayID:       54,
			ClientPrefix:          "Dw",
			MountScale:            1.0,
			BaseLanguage:          7,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1090,
			StartingTaxiNodes:     32,
			ClientFileString:      "Dwarf",
			CinematicSequenceID:   41,
			Name:                  "Dwarf",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "PIERCINGS",
			HairCustomizationName: "NORMAL",
		},
		4: &Race{
			ID:                    4,
			Flags:                 4,
			FactionID:             4,
			Unk:                   4145,
			MaleDisplayID:         55,
			FemaleDisplayID:       56,
			ClientPrefix:          "Ni",
			MountScale:            1.2000000476837158,
			BaseLanguage:          7,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     100663296,
			ClientFileString:      "NightElf",
			CinematicSequenceID:   61,
			Name:                  "Night Elf",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "MARKINGS",
			HairCustomizationName: "NORMAL",
		},
		5: &Race{
			ID:                    5,
			Flags:                 12,
			FactionID:             5,
			Unk:                   4142,
			MaleDisplayID:         57,
			FemaleDisplayID:       58,
			ClientPrefix:          "Sc",
			MountScale:            1.0,
			BaseLanguage:          1,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     1024,
			ClientFileString:      "Scourge",
			CinematicSequenceID:   2,
			Name:                  "Undead",
			MaleFeatureName:       "FEATURES",
			FemaleFeatureName:     "FEATURES",
			HairCustomizationName: "NORMAL",
		},
		6: &Race{
			ID:                    6,
			Flags:                 14,
			FactionID:             6,
			Unk:                   4143,
			MaleDisplayID:         59,
			FemaleDisplayID:       60,
			ClientPrefix:          "Ta",
			MountScale:            0.75,
			BaseLanguage:          1,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     2097152,
			ClientFileString:      "Tauren",
			CinematicSequenceID:   141,
			Name:                  "Tauren",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "HAIR",
			HairCustomizationName: "HORNS",
		},
		7: &Race{
			ID:                    7,
			Flags:                 12,
			FactionID:             115,
			Unk:                   4146,
			MaleDisplayID:         1563,
			FemaleDisplayID:       1564,
			ClientPrefix:          "Gn",
			MountScale:            0.800000011920929,
			BaseLanguage:          7,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     32,
			ClientFileString:      "Gnome",
			CinematicSequenceID:   101,
			Name:                  "Gnome",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "EARRINGS",
			HairCustomizationName: "NORMAL",
		},
		8: &Race{
			ID:                    8,
			Flags:                 14,
			FactionID:             116,
			Unk:                   4144,
			MaleDisplayID:         1478,
			FemaleDisplayID:       1479,
			ClientPrefix:          "Tr",
			MountScale:            1.0,
			BaseLanguage:          1,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     4194304,
			ClientFileString:      "Troll",
			CinematicSequenceID:   121,
			Name:                  "Troll",
			MaleFeatureName:       "TUSKS",
			FemaleFeatureName:     "TUSKS",
			HairCustomizationName: "NORMAL",
		},
		9: &Race{
			ID:                    9,
			Flags:                 1,
			FactionID:             1,
			Unk:                   0,
			MaleDisplayID:         1140,
			FemaleDisplayID:       1140,
			ClientPrefix:          "Gb",
			MountScale:            1.0,
			BaseLanguage:          7,
			CreatureType:          7,
			LoginEffectSpellID:    836,
			CombatStunSpellID:     1604,
			ResSicknessSpellID:    15007,
			SplashSoundID:         1096,
			StartingTaxiNodes:     0,
			ClientFileString:      "Goblin",
			CinematicSequenceID:   0,
			Name:                  "Goblin",
			MaleFeatureName:       "NORMAL",
			FemaleFeatureName:     "NONE",
			HairCustomizationName: "NORMAL",
		},
	}

	// Set the index.
	RaceByIndex = make(map[string]*Race)

	// Set the index values.
	RaceByIndex["Human"] = RaceByID[1]
	RaceByIndex["Orc"] = RaceByID[2]
	RaceByIndex["Dwarf"] = RaceByID[3]
	RaceByIndex["NightElf"] = RaceByID[4]
	RaceByIndex["Undead"] = RaceByID[5]
	RaceByIndex["Tauren"] = RaceByID[6]
	RaceByIndex["Gnome"] = RaceByID[7]
	RaceByIndex["Troll"] = RaceByID[8]
	RaceByIndex["Goblin"] = RaceByID[9]

	// As there is only a single index, add some special convenience values.
	RaceHuman = RaceByID[1]
	RaceOrc = RaceByID[2]
	RaceDwarf = RaceByID[3]
	RaceNightElf = RaceByID[4]
	RaceUndead = RaceByID[5]
	RaceTauren = RaceByID[6]
	RaceGnome = RaceByID[7]
	RaceTroll = RaceByID[8]
	RaceGoblin = RaceByID[9]
}
