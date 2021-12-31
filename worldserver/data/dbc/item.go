package dbc

import (
	"time"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// ItemDamage represents the damage an item can do, as a min/max pair.
type ItemDamage struct {
	Min float32
	Max float32
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
	DisplayID     int
	Class         c.ItemClass
	SubClass      c.ItemSubClass
	Quality       c.ItemQuality
	InventoryType c.InventoryType
	Level         int

	VendorBuyPrice  c.Coins
	VendorSellPrice c.Coins

	// Item stat information.
	AttackRate    time.Duration
	Resistances   map[c.SpellSchool]int
	Damages       map[c.SpellSchool]ItemDamage
	MaxDurability int
	SheathType    c.SheathType

	// Basic item information.
	// Entry              int
	// Bonding            c.Bonding
	// Class              c.ItemClass
	// Description        string
	// DisplayID          int
	// ItemLevel          int
	// Material           int
	// MaxStackSize       int
	// Name               string
	// PerCharacterLimit  int
	// Quality            c.ItemQuality
	// SubClass           c.ItemSubClass
	// TimeUntilDisappear int

	// // Equipment item information.
	// AttackRate    int
	// Block         float32
	// InventoryType c.InventoryType
	// MaxDurability int
	// SheathType    c.SheathType

	// Resistances map[c.SpellSchool]int
	// Stats       map[c.Stat]int
	// Spells      map[int]struct {
	// 	ID                int
	// 	Trigger           int
	// 	Charges           int
	// 	ProcPerMinuteRate float32
	// 	Cooldown          int
	// 	Category          c.SpellCategory
	// 	CategoryCooldown  int
	// }
	// Damage map[c.SpellSchool]struct {
	// 	Min float32
	// 	Max float32
	// }

	// // Food item information.
	// FoodType c.FoodType

	// // Readable item information.
	// Language     c.Language
	// PageMaterial int

	// // Bag item information.
	// BagFamily      c.BagFamily
	// ContainerSlots int
	// MaxMoneyLoot   int
	// MinMoneyLoot   int

	// // Ranged item information.
	// RangedModRange   float32
	// RequiredAmmoType c.ItemSubClass

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
func (i *Item) Flags() c.ItemPrototypeFlag {
	var flags c.ItemPrototypeFlag
	// if i.IsCharter {
	// 	flags |= c.ItemPrototypeFlagCharter
	// }
	// if i.IsConjured {
	// 	flags |= c.ItemPrototypeFlagConjured
	// }
	// if i.IsIndestructible {
	// 	flags |= c.ItemPrototypeFlagIndestructible
	// }
	// if i.IsLetter {
	// 	flags |= c.ItemPrototypeFlagLetter
	// }
	// if i.IsLootable {
	// 	flags |= c.ItemPrototypeFlagLootable
	// }
	// if i.IsNoEquipCooldown {
	// 	flags |= c.ItemPrototypeFlagNoEquipCooldown
	// }
	// if i.IsPartyLoot {
	// 	flags |= c.ItemPrototypeFlagPartyLoot
	// }
	// if i.IsPVPReward {
	// 	flags |= c.ItemPrototypeFlagPVPReward
	// }
	// if i.IsStackable {
	// 	flags |= c.ItemPrototypeFlagStackable
	// }
	// if i.IsUsable {
	// 	flags |= c.ItemPrototypeFlagUsable
	// }
	// if i.IsWrapper {
	// 	flags |= c.ItemPrototypeFlagWrapper
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
		DisplayID:     c.DisplayIDInvShield09,
		Class:         c.ItemClassArmor,
		SubClass:      c.ItemSubClassArmorShield,
		Quality:       c.ItemQualityPoor,
		InventoryType: c.InventoryTypeShield,
		Level:         1,

		VendorBuyPrice:  c.MakeCoins(1, 0, 0),
		VendorSellPrice: c.MakeCoins(15, 0, 0),

		Resistances: map[c.SpellSchool]int{
			c.SpellSchoolPhysical: 5,
		},
		MaxDurability: 20,
	})
	addItem(&Item{
		Name:          "Worn Shortsword",
		DisplayID:     c.DisplayIDInvSword04,
		Class:         c.ItemClassWeapon,
		SubClass:      c.ItemSubClassWeaponSword,
		Quality:       c.ItemQualityNormal,
		InventoryType: c.InventoryTypeWeaponMainHand,
		Level:         2,

		VendorBuyPrice:  c.MakeCoins(1, 0, 0),
		VendorSellPrice: c.MakeCoins(15, 0, 0),

		AttackRate: makeDuration("1.9s"),
		Damages: map[c.SpellSchool]ItemDamage{
			c.SpellSchoolPhysical: {
				Min: 1.0,
				Max: 3.0,
			},
		},
		MaxDurability: 20,
		SheathType:    c.SheathTypeLargeWeaponLeft,
	})
	addItem(&Item{
		Name:          "Recruit's Shirt",
		DisplayID:     c.DisplayIDInvShirt05,
		Class:         c.ItemClassArmor,
		Quality:       c.ItemQualityNormal,
		InventoryType: c.InventoryTypeChest,
		Level:         1,

		VendorBuyPrice:  c.MakeCoins(1, 0, 0),
		VendorSellPrice: c.MakeCoins(15, 0, 0),
	})
	addItem(&Item{
		Name:          "Recruit's Boots",
		DisplayID:     c.DisplayIDInvBoots06,
		Class:         c.ItemClassArmor,
		Quality:       c.ItemQualityNormal,
		InventoryType: c.InventoryTypeFeet,
		Level:         1,

		VendorBuyPrice:  c.MakeCoins(1, 0, 0),
		VendorSellPrice: c.MakeCoins(15, 0, 0),
	})
	addItem(&Item{
		Name:          "Recruit's Pants",
		DisplayID:     c.DisplayIDInvPants02,
		Class:         c.ItemClassArmor,
		Quality:       c.ItemQualityPoor,
		InventoryType: c.InventoryTypeLegs,
		Level:         1,

		VendorBuyPrice:  c.MakeCoins(1, 0, 0),
		VendorSellPrice: c.MakeCoins(15, 0, 0),

		Resistances: map[c.SpellSchool]int{
			c.SpellSchoolPhysical: 2,
		},
		MaxDurability: 25,
	})
}
