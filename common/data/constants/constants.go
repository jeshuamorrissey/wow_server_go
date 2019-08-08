package constants

// SpellSchool information.
//go:generate stringer -type=SpellSchool -trimprefix=SpellSchool
type SpellSchool uint8

// SpellSchool values.
const (
	SpellSchoolPhysical SpellSchool = 0
	SpellSchoolHoly     SpellSchool = 1
	SpellSchoolFire     SpellSchool = 2
	SpellSchoolNature   SpellSchool = 3
	SpellSchoolFrost    SpellSchool = 4
	SpellSchoolShadow   SpellSchool = 5
	SpellSchoolArcane   SpellSchool = 6
)

// BagFamily information.
//go:generate stringer -type=BagFamily -trimprefix=BagFamily
type BagFamily uint8

// BagFamily values.
const (
	BagFamilyNone               BagFamily = 1
	BagFamilyArrows             BagFamily = 1
	BagFamilyBullets            BagFamily = 2
	BagFamilySoulShards         BagFamily = 3
	BagFamilyUnknown1           BagFamily = 4
	BagFamilyUnknown2           BagFamily = 5
	BagFamilyHerbs              BagFamily = 6
	BagFamilyEnchantingSupport  BagFamily = 7
	BagFamilyEngineeringSupport BagFamily = 8
	BagFamilyKeys               BagFamily = 9
)

// Bonding information.
//go:generate stringer -type=Bonding -trimprefix=Bonding
type Bonding uint8

// Bonding values.
const (
	BondingNone         Bonding = 0
	BondingWhenPickedUp Bonding = 1
	BondingWhenEquipped Bonding = 2
	BondingOnUse        Bonding = 3
	BondingQuestItem    Bonding = 4
)

// ItemClass information.
//go:generate stringer -type=ItemClass -trimprefix=ItemClass
type ItemClass uint8

// ItemClass values.
const (
	ItemClassConsumable ItemClass = 0
	ItemClassContainer  ItemClass = 1
	ItemClassWeapon     ItemClass = 2
	ItemClassArmor      ItemClass = 4
	ItemClassReagent    ItemClass = 5
	ItemClassProjectile ItemClass = 6
	ItemClassTradeGoods ItemClass = 7
	ItemClassRecipe     ItemClass = 9
	ItemClassQuiver     ItemClass = 11
	ItemClassQuest      ItemClass = 12
	ItemClassKey        ItemClass = 13
	ItemClassMisc       ItemClass = 15
)

// ItemSubClass information.
//go:generate stringer -type=ItemSubClass -trimprefix=ItemSubClass
type ItemSubClass uint8

// ItemSubClass values.
const (
	ItemSubClassConsumable                ItemSubClass = 0
	ItemSubClassConsumablePotion          ItemSubClass = 1
	ItemSubClassConsumableElixir          ItemSubClass = 2
	ItemSubClassConsumableFlask           ItemSubClass = 3
	ItemSubClassConsumableScroll          ItemSubClass = 4
	ItemSubClassConsumableFood            ItemSubClass = 5
	ItemSubClassConsumableItemEnhancement ItemSubClass = 6
	ItemSubClassConsumableBandage         ItemSubClass = 7
	ItemSubClassConsumableOther           ItemSubClass = 8

	ItemSubClassContainer            ItemSubClass = 0
	ItemSubClassContainerSoul        ItemSubClass = 1
	ItemSubClassContainerHerb        ItemSubClass = 2
	ItemSubClassContainerEnchanting  ItemSubClass = 3
	ItemSubClassContainerEngineering ItemSubClass = 4

	ItemSubClassWeaponAxe         ItemSubClass = 0
	ItemSubClassWeaponAxe2        ItemSubClass = 1
	ItemSubClassWeaponBow         ItemSubClass = 2
	ItemSubClassWeaponGun         ItemSubClass = 3
	ItemSubClassWeaponMace        ItemSubClass = 4
	ItemSubClassWeaponMace2       ItemSubClass = 5
	ItemSubClassWeaponPolearm     ItemSubClass = 6
	ItemSubClassWeaponSword       ItemSubClass = 7
	ItemSubClassWeaponSword2      ItemSubClass = 8
	ItemSubClassWeaponStaff       ItemSubClass = 10
	ItemSubClassWeaponFist        ItemSubClass = 13
	ItemSubClassWeaponMisc        ItemSubClass = 14
	ItemSubClassWeaponDagger      ItemSubClass = 15
	ItemSubClassWeaponThrown      ItemSubClass = 16
	ItemSubClassWeaponSpear       ItemSubClass = 17
	ItemSubClassWeaponCrossbow    ItemSubClass = 18
	ItemSubClassWeaponWand        ItemSubClass = 19
	ItemSubClassWeaponFishingPole ItemSubClass = 20

	ItemSubClassArmorMisc    ItemSubClass = 0
	ItemSubClassArmorCloth   ItemSubClass = 1
	ItemSubClassArmorLeather ItemSubClass = 2
	ItemSubClassArmorMail    ItemSubClass = 3
	ItemSubClassArmorPlate   ItemSubClass = 4
	ItemSubClassArmorBuckler ItemSubClass = 5
	ItemSubClassArmorShield  ItemSubClass = 6
	ItemSubClassArmorLibram  ItemSubClass = 7
	ItemSubClassArmorIdol    ItemSubClass = 8
	ItemSubClassArmorTotem   ItemSubClass = 9

	ItemSubClassReagent ItemSubClass = 0

	ItemSubClassProjectileArrow  ItemSubClass = 2
	ItemSubClassProjectileBullet ItemSubClass = 3

	ItemSubClassTradeGoods           ItemSubClass = 0
	ItemSubClassTradeGoodsParts      ItemSubClass = 1
	ItemSubClassTradeGoodsExplosives ItemSubClass = 2
	ItemSubClassTradeGoodsDevices    ItemSubClass = 3

	ItemSubClassRecipeBook                  ItemSubClass = 0
	ItemSubClassRecipeLeatherworkingPattern ItemSubClass = 1
	ItemSubClassRecipeTailoringPattern      ItemSubClass = 2
	ItemSubClassRecipeEngineeringSchematic  ItemSubClass = 3
	ItemSubClassRecipeBlacksmithing         ItemSubClass = 4
	ItemSubClassRecipeCookingRecipe         ItemSubClass = 5
	ItemSubClassRecipeAlchemyRecipe         ItemSubClass = 6
	ItemSubClassRecipeFirstAidManual        ItemSubClass = 7
	ItemSubClassRecipeEnchantingFormula     ItemSubClass = 8
	ItemSubClassRecipeFishingManual         ItemSubClass = 9

	ItemSubClassQuiver    ItemSubClass = 2
	ItemSubClassAmmoPouch ItemSubClass = 3

	ItemSubClassQuest ItemSubClass = 0

	ItemSubClassKey         ItemSubClass = 0
	ItemSubClassKeyLockpick ItemSubClass = 1

	ItemSubClassPermanent ItemSubClass = 0

	ItemSUbClassJunk ItemSubClass = 0
)

// FoodType information.
//go:generate stringer -type=FoodType -trimprefix=FoodType
type FoodType uint8

// FoodType values.
const (
	FoodTypeMeat    FoodType = 1
	FoodTypeFish    FoodType = 2
	FoodTypeCheese  FoodType = 3
	FoodTypeBread   FoodType = 4
	FoodTypeFungas  FoodType = 5
	FoodTypeFruit   FoodType = 6
	FoodTypeRawMeat FoodType = 7
	FoodTypeRawFish FoodType = 8
)

// InventoryType information.
//go:generate stringer -type=InventoryType -trimprefix=InventoryType
type InventoryType uint8

// InventoryType values.
const (
	InventoryTypeNonEquip       InventoryType = 0
	InventoryTypeHead           InventoryType = 1
	InventoryTypeNeck           InventoryType = 2
	InventoryTypeShoulders      InventoryType = 3
	InventoryTypeBody           InventoryType = 4
	InventoryTypeChest          InventoryType = 5
	InventoryTypeWaist          InventoryType = 6
	InventoryTypeLegs           InventoryType = 7
	InventoryTypeFeet           InventoryType = 8
	InventoryTypeWrists         InventoryType = 9
	InventoryTypeHands          InventoryType = 10
	InventoryTypeFinger         InventoryType = 11
	InventoryTypeTrinket        InventoryType = 12
	InventoryTypeWeapon         InventoryType = 13
	InventoryTypeShield         InventoryType = 14
	InventoryTypeRanged         InventoryType = 15
	InventoryTypeCloak          InventoryType = 16
	InventoryType2HWeapon       InventoryType = 17
	InventoryTypeBag            InventoryType = 18
	InventoryTypeTabard         InventoryType = 19
	InventoryTypeRobe           InventoryType = 20
	InventoryTypeWeaponMainHand InventoryType = 21
	InventoryTypeWeaponOffHand  InventoryType = 22
	InventoryTypeHoldable       InventoryType = 23
	InventoryTypeAmmo           InventoryType = 24
	InventoryTypeThrown         InventoryType = 25
	InventoryTypeRangedRight    InventoryType = 26
	InventoryTypeQuiver         InventoryType = 27
	InventoryTypeRelic          InventoryType = 28
)

// Language information.
//go:generate stringer -type=Language -trimprefix=Language
type Language uint8

// Language values.
const (
	LanguageUniversal   Language = 0
	LanguageOrcish      Language = 1
	LanguageDarnassian  Language = 2
	LanguageTaurahe     Language = 3
	LanguageDwarvish    Language = 6
	LanguageCommon      Language = 7
	LanguageDemonic     Language = 8
	LanguageTitan       Language = 9
	LanguageThalassian  Language = 10
	LanguageDraconic    Language = 11
	LanguageKalimag     Language = 12
	LanguageGnomish     Language = 13
	LanguageTroll       Language = 14
	LanguageGutterspeak Language = 33
)

// ItemQuality information.
//go:generate stringer -type=ItemQuality -trimprefix=ItemQuality
type ItemQuality uint8

// ItemQuality values.
const (
	ItemQualityPoor      ItemQuality = 0
	ItemQualityNormal    ItemQuality = 1
	ItemQualityUncommon  ItemQuality = 2
	ItemQualityRare      ItemQuality = 3
	ItemQualityEpic      ItemQuality = 4
	ItemQualityLegendary ItemQuality = 5
	ItemQualityArtifact  ItemQuality = 6
)

// SheathType information.
//go:generate stringer -type=SheathType -trimprefix=SheathType
type SheathType uint8

// SheathType values.
const (
	SheathTypeNone             SheathType = 0
	SheathTypeMainHand         SheathType = 1
	SheathTypeOffHand          SheathType = 2
	SheathTypeLargeWeaponLeft  SheathType = 3
	SheathTypeLargeWeaponRight SheathType = 4
	SheathTypeHipWeaponLeft    SheathType = 5
	SheathTypeHipWeaponRight   SheathType = 6
	SheathTypeShield           SheathType = 7
)

// Stat information.
//go:generate stringer -type=Stat -trimprefix=Stat
type Stat uint8

// Stat values.
const (
	StatStrength  Stat = 0
	StatAgility   Stat = 1
	StatStamina   Stat = 2
	StatIntellect Stat = 3
	StatSpirit    Stat = 4
)

// SpellCategory information.
//go:generate stringer -type=SpellCategory -trimprefix=SpellCategory
type SpellCategory uint8

// SpellCategory values.
const (
	SpellCategoryHealthManaPotions SpellCategory = 4
	SpellCategoryDevourMagic       SpellCategory = 12
)
