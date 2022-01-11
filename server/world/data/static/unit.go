package static

// Unit represents a template within the world.
type Unit struct {
	// Basic information.
	Entry   int
	Name    string
	SubName string

	// Combat information.
	MaxHealth int
	MaxPower  int

	// Size information.
	BoundingRadius float32
	CombatReach    float32
	DisplayID      uint32

	// Entry    int
	// Name     string
	// SubName  string
	// MinLevel int
	// MaxLevel int

	// Models []ModelInfo

	// FactionAlliance      int
	// FactionHorde         int
	// Scale                float32
	// Family               int
	// CreatureType         int
	// InhabitType          int
	// RegenerateStats      int
	// RacialLeader         int
	// DynamicFlags         int
	// SpeedWalk            float32
	// SpeedRun             float32
	// UnitClass            int
	// Rank                 int
	// HealthMultiplier     float32
	// PowerMultiplier      float32
	// DamageMultiplier     float32
	// DamageVariance       float32
	// ArmorMultiplier      float32
	// ExperienceMultiplier float32
	// MinLevelHealth       int
	// MaxLevelHealth       int
	// MinLevelMana         int
	// MaxLevelMana         int
	// MinMeleeDmg          float32
	// MaxMeleeDmg          float32
	// MinRangedDmg         float32
	// MaxRangedDmg         float32
	// Armor                int
	// MeleeAttackPower     int
	// RangedAttackPower    int
	// MeleeBaseAttackTime  int
	// RangedBaseAttackTime int
	// DamageSchool         int
	// MinLootGold          int
	// MaxLootGold          int
	// LootID               int
	// PickpocketLootID     int
	// SkinningLootID       int
	// KillCredit1          int
	// KillCredit2          int
	// MechanicImmuneMask   int
	// SchoolImmuneMask     int
	// ResistanceHoly       int
	// ResistanceFire       int
	// ResistanceNature     int
	// ResistanceFrost      int
	// ResistanceShadow     int
	// ResistanceArcane     int
	// PetSpellDataID       int
	// MovementType         int
	// TrainerType          int
	// TrainerSpell         int
	// TrainerClass         int
	// TrainerRace          int
	// TrainerTemplateID    int
	// VendorTemplateID     int
	// GossipMenuID         int
	// EquipmentTemplateID  int
	// Civilian             int

	// // Flags.
	// HasGossip        bool
	// IsQuestgiver     bool
	// IsVendor         bool
	// IsFlightmaster   bool
	// IsTrainer        bool
	// IsSpirithealer   bool
	// IsSpiritguide    bool
	// IsInnkeeper      bool
	// IsBanker         bool
	// IsPetitioner     bool
	// IsTabarddesigner bool
	// IsBattlemaster   bool
	// IsAuctioneer     bool
	// IsStablemaster   bool
	// CanRepair        bool

	// // ExtraFlags.
	// IsInstanceBound    bool
	// NoAggro            bool
	// NoParry            bool
	// NoParryHasten      bool
	// NoBlock            bool
	// NoCrush            bool
	// NoXPAtKill         bool
	// IsInvisible        bool
	// IsNotTauntable     bool
	// HasAggroZone       bool
	// IsGuard            bool
	// NoCallAssist       bool
	// IsActive           bool
	// IsMMapForceEnable  bool
	// IsMMapForceDisable bool
	// WalksInWater       bool
	// HasNoSwimAnimation bool

	// // CreatureTypeFlags.
	// IsTameable     bool
	// IsGhostVisible bool
	// IsHerbLoot     bool
	// IsMiningLoot   bool
	// CanAssist      bool
	// IsEngineerLoot bool
}

// ModelInfo contains information about a in-game model.
type ModelInfo struct {
	ID             int
	BoundingRadius float32
	CombatReach    float32
	Gender         Gender
}

var (
	nextUnitEntry = 1

	// Units is a map of unit entry --> object.
	Units map[int]*Unit

	// UnitsByName is a map of unit name --> object.
	UnitsByName map[string]*Unit
)

func addUnit(unit *Unit) {
	entry := nextUnitEntry
	nextUnitEntry++

	unit.Entry = entry
	Units[entry] = unit
	UnitsByName[unit.Name] = unit
}

// Flags returns a bitmask of various flags.
func (u *Unit) Flags() int {
	var flags int
	// if u.HasGossip {
	// 	flags |= 0x00000001
	// }
	// if u.IsQuestgiver {
	// 	flags |= 0x00000002
	// }
	// if u.IsVendor {
	// 	flags |= 0x00000004
	// }
	// if u.IsFlightmaster {
	// 	flags |= 0x00000008
	// }
	// if u.IsTrainer {
	// 	flags |= 0x00000010
	// }
	// if u.IsSpirithealer {
	// 	flags |= 0x00000020
	// }
	// if u.IsSpiritguide {
	// 	flags |= 0x00000040
	// }
	// if u.IsInnkeeper {
	// 	flags |= 0x00000080
	// }
	// if u.IsBanker {
	// 	flags |= 0x00000100
	// }
	// if u.IsPetitioner {
	// 	flags |= 0x00000200
	// }
	// if u.IsTabarddesigner {
	// 	flags |= 0x00000400
	// }
	// if u.IsBattlemaster {
	// 	flags |= 0x00000800
	// }
	// if u.IsAuctioneer {
	// 	flags |= 0x00001000
	// }
	// if u.IsStablemaster {
	// 	flags |= 0x00002000
	// }
	// if u.CanRepair {
	// 	flags |= 0x00004000
	// }

	return flags
}

// ExtraFlags returns a bitmask of the ExtraFlags.
func (u *Unit) ExtraFlags() int {
	var flags int
	// if u.IsInstanceBound {
	// 	flags |= 0x00000001
	// }
	// if u.NoAggro {
	// 	flags |= 0x00000002
	// }
	// if u.NoParry {
	// 	flags |= 0x00000004
	// }
	// if u.NoParryHasten {
	// 	flags |= 0x00000008
	// }
	// if u.NoBlock {
	// 	flags |= 0x00000010
	// }
	// if u.NoCrush {
	// 	flags |= 0x00000020
	// }
	// if u.NoXPAtKill {
	// 	flags |= 0x00000040
	// }
	// if u.IsInvisible {
	// 	flags |= 0x00000080
	// }
	// if u.IsNotTauntable {
	// 	flags |= 0x00000100
	// }
	// if u.HasAggroZone {
	// 	flags |= 0x00000200
	// }
	// if u.IsGuard {
	// 	flags |= 0x00000400
	// }
	// if u.NoCallAssist {
	// 	flags |= 0x00000800
	// }
	// if u.IsActive {
	// 	flags |= 0x00001000
	// }
	// if u.IsMMapForceEnable {
	// 	flags |= 0x00002000
	// }
	// if u.IsMMapForceDisable {
	// 	flags |= 0x00004000
	// }
	// if u.WalksInWater {
	// 	flags |= 0x00008000
	// }
	// if u.HasNoSwimAnimation {
	// 	flags |= 0x00010000
	// }
	return flags
}

// CreatureTypeFlags returns a bitmask of the CreatureTypeFlags.
func (u *Unit) CreatureTypeFlags() int {
	var flags int
	// if u.IsTameable {
	// 	flags |= 0x00000001
	// }
	// if u.IsGhostVisible {
	// 	flags |= 0x00000002
	// }
	// if u.IsHerbLoot {
	// 	flags |= 0x00000100
	// }
	// if u.IsMiningLoot {
	// 	flags |= 0x00000200
	// }
	// if u.CanAssist {
	// 	flags |= 0x00001000
	// }
	// if u.IsEngineerLoot {
	// 	flags |= 0x00008000
	// }
	return flags
}

// GetPlayerModelInfo gets the ModelInfo structure for player races.
// Taken from ChrRaces.dbc
func GetPlayerModelInfo(race *Race, gender Gender) ModelInfo {
	if gender == GenderMale {
		switch race.ID {
		case RaceHuman.ID:
			return ModelInfo{
				ID:             49,
				BoundingRadius: 0.306,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceOrc.ID:
			return ModelInfo{
				ID:             51,
				BoundingRadius: 0.372,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceDwarf.ID:
			return ModelInfo{
				ID:             53,
				BoundingRadius: 0.347,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceNightElf.ID:
			return ModelInfo{
				ID:             55,
				BoundingRadius: 0.389,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceUndead.ID:
			return ModelInfo{
				ID:             57,
				BoundingRadius: 0.383,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceTauren.ID:
			return ModelInfo{
				ID:             59,
				BoundingRadius: 0.9747,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceGnome.ID:
			return ModelInfo{
				ID:             1563,
				BoundingRadius: 0.3519,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceTroll.ID:
			return ModelInfo{
				ID:             1478,
				BoundingRadius: 0.306,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceGoblin.ID:
			return ModelInfo{
				ID:             1140,
				BoundingRadius: 0.347222,
				CombatReach:    1.5,
				Gender:         gender,
			}
		}
	} else {
		switch race.ID {
		case RaceHuman.ID:
			return ModelInfo{
				ID:             50,
				BoundingRadius: 0.208,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceOrc.ID:
			return ModelInfo{
				ID:             52,
				BoundingRadius: 0.236,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceDwarf.ID:
			return ModelInfo{
				ID:             54,
				BoundingRadius: 0.347,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceNightElf.ID:
			return ModelInfo{
				ID:             56,
				BoundingRadius: 0.306,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceUndead.ID:
			return ModelInfo{
				ID:             58,
				BoundingRadius: 0.383,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceTauren.ID:
			return ModelInfo{
				ID:             60,
				BoundingRadius: 0.8725,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceGnome.ID:
			return ModelInfo{
				ID:             1564,
				BoundingRadius: 0.3519,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceTroll.ID:
			return ModelInfo{
				ID:             1479,
				BoundingRadius: 0.306,
				CombatReach:    1.5,
				Gender:         gender,
			}
		case RaceGoblin.ID:
			return ModelInfo{
				ID:             1140,
				BoundingRadius: 0.347222,
				CombatReach:    1.5,
				Gender:         gender,
			}
		}
	}

	return ModelInfo{}
}

// GetPlayerScale returns the ScaleX factor for a given race/class/gender.
func GetPlayerScale(race *Race, gender Gender) float32 {
	if race.ID == RaceTauren.ID {
		if gender == GenderFemale {
			return 1.25
		}

		return 1.35
	}

	return 1.0
}

func init() {
	Units = make(map[int]*Unit)
	UnitsByName = make(map[string]*Unit)

	addUnit(&Unit{
		Name: "The Man",

		MaxHealth: 15,
		MaxPower:  100,

		BoundingRadius: 0.306,
		CombatReach:    1.5,
		DisplayID:      5035,
	})
}
