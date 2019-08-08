package data

import (
	"compress/gzip"
	"encoding/json"
	"os"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// Item represents an item template within the world. This does not represent
// an individual item (this is done via a GameObject).
type Item struct {
	// DisenchantId  TODO
	// ItemSet TODO
	// LockId TODO
	// PageText TODO
	// RandomProperty TODO
	// StartQuest TODO

	// Basic item information.
	Bonding            c.Bonding
	Class              c.ItemClass
	Description        string
	DisplayID          int
	ItemLevel          int
	Material           int
	MaxStackSize       int
	Name               string
	PerCharacterLimit  int
	Quality            c.ItemQuality
	SubClass           c.ItemSubClass
	TimeUntilDisappear int

	// Equipment item information.
	AttackRate    int
	Block         float32
	InventoryType c.InventoryType
	MaxDurability int
	SheathType    c.SheathType

	Resistances map[c.SpellSchool]int
	Stats       map[c.Stat]int
	Spells      map[int]struct {
		ID                int
		Trigger           int
		Charges           int
		ProcPerMinuteRate float32
		Cooldown          int
		Category          c.SpellCategory
		CategoryCooldown  int
	}
	Damage map[c.SpellSchool]struct {
		Min float32
		Max float32
	}

	// Food item information.
	FoodType c.FoodType

	// Readable item information.
	Language     c.Language
	PageMaterial int

	// Bag item information.
	BagFamily      c.BagFamily
	ContainerSlots int
	MaxMoneyLoot   int
	MinMoneyLoot   int

	// Ranged item information.
	RangedModRange   float32
	RequiredAmmoType c.ItemSubClass

	// Item requirements.
	AllowableClassMask        int
	AllowableRaceMask         int
	RequiredArea              int
	RequiredHonorRank         int
	RequiredLevel             int
	RequiredMap               int
	RequiredReputationFaction int // TODO(jeshua): load from DBC
	RequiredReputationRank    int
	RequiredSkill             int // TODO(jeshua): load from DBC
	RequiredSkillRank         int
	RequiredSpell             int // TODO(jeshua): load from DBC

	// Vendor information.
	VendorBuyPrice  int
	VendorSellPrice int
	VendorStackSize int

	// Flags
	IsCharter           bool
	IsConjured          bool
	IsIndestructible    bool
	IsLetter            bool
	IsLootable          bool
	IsNoEquipCooldown   bool
	IsNonConsumable     bool
	IsPartyLoot         bool
	IsPVPReward         bool
	IsRealmTimeDuration bool
	IsStackable         bool
	IsUsable            bool
	IsWrapper           bool
}

var (
	// Items is a map of item entry --> object.
	Items map[int]Item
)

// LoadItems will load the items data from the given compressed JSON file.
func LoadItems(compressedJSONFile string) error {
	file, err := os.Open(compressedJSONFile)
	if err != nil {
		return err
	}

	gz, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	json.NewDecoder(gz).Decode(&Items)

	return nil
}
