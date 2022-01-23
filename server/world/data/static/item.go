package static

import (
	"time"
)

// ItemDamage represents the damage an item can do, as a min/max pair.
type ItemDamage struct {
	Min int
	Max int
}

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
	Entry         int
	Name          string
	Description   string
	DisplayID     DisplayID
	Class         ItemClass
	SubClass      ItemSubClass
	Quality       ItemQuality
	InventoryType InventoryType
	Level         int

	VendorBuyPrice  Coins
	VendorSellPrice Coins

	// Item stat information.
	AttackRate    time.Duration
	Resistances   map[SpellSchool]int
	Damages       map[SpellSchool]ItemDamage
	MaxDurability int
	SheathType    SheathType

	// Basic item information.
	// Entry              int
	// Bonding            Bonding
	// Class              ItemClass
	// Description        string
	// DisplayID          int
	// ItemLevel          int
	// Material           int
	// MaxStackSize       int
	// Name               string
	// PerCharacterLimit  int
	// Quality            ItemQuality
	// SubClass           ItemSubClass
	// TimeUntilDisappear int

	// // Equipment item information.
	// AttackRate    int
	// Block         float32
	// InventoryType InventoryType
	// MaxDurability int
	// SheathType    SheathType

	// Resistances map[SpellSchool]int
	// Stats       map[Stat]int
	// Spells      map[int]struct {
	// 	ID                int
	// 	Trigger           int
	// 	Charges           int
	// 	ProcPerMinuteRate float32
	// 	Cooldown          int
	// 	Category          SpellCategory
	// 	CategoryCooldown  int
	// }
	// Damage map[SpellSchool]struct {
	// 	Min float32
	// 	Max float32
	// }

	// // Food item information.
	// FoodType FoodType

	// // Readable item information.
	// Language     Language
	// PageMaterial int

	// // Bag item information.
	// BagFamily      BagFamily
	// ContainerSlots int
	// MaxMoneyLoot   int
	// MinMoneyLoot   int

	// // Ranged item information.
	// RangedModRange   float32
	// RequiredAmmoType ItemSubClass

	// // Item requirements.
	// AllowableClassMask        int
	// AllowableRaceMask         int
	// RequiredArea              int
	// RequiredHonorRank         int
	// RequiredLevel             int
	// RequiredMap               int
	// RequiredReputationFaction int // TODO(jeshua): load from DBC
	// RequiredReputationRank    int
	// RequiredSkill             int // TODO(jeshua): load from DBC
	// RequiredSkillRank         int
	// RequiredSpell             int // TODO(jeshua): load from DBC

	// // Vendor information.
	// VendorBuyPrice  int
	// VendorSellPrice int
	// VendorStackSize int

	// // Flags
	// IsCharter           bool
	// IsConjured          bool
	// IsIndestructible    bool
	// IsLetter            bool
	// IsLootable          bool
	// IsNoEquipCooldown   bool
	// IsNonConsumable     bool
	// IsPartyLoot         bool
	// IsPVPReward         bool
	// IsRealmTimeDuration bool
	// IsStackable         bool
	// IsUsable            bool
	// IsWrapper           bool
}

// Flags returns the various IsX attributes as a set of binary flags.
func (i *Item) Flags() ItemPrototypeFlag {
	var flags ItemPrototypeFlag
	// if i.IsCharter {
	// 	flags |= ItemPrototypeFlagCharter
	// }
	// if i.IsConjured {
	// 	flags |= ItemPrototypeFlagConjured
	// }
	// if i.IsIndestructible {
	// 	flags |= ItemPrototypeFlagIndestructible
	// }
	// if i.IsLetter {
	// 	flags |= ItemPrototypeFlagLetter
	// }
	// if i.IsLootable {
	// 	flags |= ItemPrototypeFlagLootable
	// }
	// if i.IsNoEquipCooldown {
	// 	flags |= ItemPrototypeFlagNoEquipCooldown
	// }
	// if i.IsPartyLoot {
	// 	flags |= ItemPrototypeFlagPartyLoot
	// }
	// if i.IsPVPReward {
	// 	flags |= ItemPrototypeFlagPVPReward
	// }
	// if i.IsStackable {
	// 	flags |= ItemPrototypeFlagStackable
	// }
	// if i.IsUsable {
	// 	flags |= ItemPrototypeFlagUsable
	// }
	// if i.IsWrapper {
	// 	flags |= ItemPrototypeFlagWrapper
	// }
	return flags
}

var (
	nextItemEntry = 1

	// Items is a map of item entry --> object.
	Items map[int]*Item

	// ItemsByName is a map of item name --> object.
	ItemsByName map[string]*Item
)

func addItem(item *Item) {
	entry := nextItemEntry
	nextItemEntry++

	item.Entry = entry
	Items[entry] = item
	ItemsByName[item.Name] = item
}

func makeDuration(duration string) time.Duration {
	d, _ := time.ParseDuration(duration)
	return d
}

func init() {
	Items = make(map[int]*Item)
	ItemsByName = make(map[string]*Item)

	addItem(&Item{
		Name:          "Worn Wooden Shield",
		DisplayID:     DisplayIDInvShield09,
		Class:         ItemClassArmor,
		SubClass:      ItemSubClassArmorShield,
		Quality:       ItemQualityPoor,
		InventoryType: InventoryTypeShield,
		Level:         1,

		VendorBuyPrice:  MakeCoins(1, 0, 0),
		VendorSellPrice: MakeCoins(15, 0, 0),

		Resistances: map[SpellSchool]int{
			SpellSchoolPhysical: 5,
		},
		MaxDurability: 20,
	})
	addItem(&Item{
		Name:          "Worn Shortsword",
		DisplayID:     DisplayIDInvSword04,
		Class:         ItemClassWeapon,
		SubClass:      ItemSubClassWeaponSword,
		Quality:       ItemQualityNormal,
		InventoryType: InventoryTypeWeaponMainHand,
		Level:         2,

		VendorBuyPrice:  MakeCoins(1, 0, 0),
		VendorSellPrice: MakeCoins(15, 0, 0),

		AttackRate: makeDuration("1.9s"),
		Damages: map[SpellSchool]ItemDamage{
			SpellSchoolPhysical: {
				Min: 1,
				Max: 3,
			},
		},
		MaxDurability: 20,
		SheathType:    SheathTypeLargeWeaponLeft,
	})
	addItem(&Item{
		Name:          "Recruit's Shirt",
		DisplayID:     DisplayIDInvShirt05,
		Class:         ItemClassArmor,
		Quality:       ItemQualityNormal,
		InventoryType: InventoryTypeChest,
		Level:         1,

		VendorBuyPrice:  MakeCoins(1, 0, 0),
		VendorSellPrice: MakeCoins(15, 0, 0),
	})
	addItem(&Item{
		Name:          "Recruit's Boots",
		DisplayID:     DisplayIDInvBoots06,
		Class:         ItemClassArmor,
		Quality:       ItemQualityNormal,
		InventoryType: InventoryTypeFeet,
		Level:         1,

		VendorBuyPrice:  MakeCoins(1, 0, 0),
		VendorSellPrice: MakeCoins(15, 0, 0),
	})
	addItem(&Item{
		Name:          "Recruit's Pants",
		DisplayID:     DisplayIDInvPants02,
		Class:         ItemClassArmor,
		Quality:       ItemQualityPoor,
		InventoryType: InventoryTypeLegs,
		Level:         1,

		VendorBuyPrice:  MakeCoins(1, 0, 0),
		VendorSellPrice: MakeCoins(15, 0, 0),

		Resistances: map[SpellSchool]int{
			SpellSchoolPhysical: 2,
		},
		MaxDurability: 25,
	})
}
