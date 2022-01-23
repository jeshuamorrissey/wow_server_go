package static

import "time"

// Misc constants that don't belong to an enum.
const (
	MinPlayerNameLength = 2
	MaxPlayerNameLength = 12
	NumBagSlots         = 4
	RegenTimeout        = 500 * time.Millisecond
)

// Coins represents the coinage type within the game.
type Coins int

// MakeCoins makes a coins object for a given distribute of money.
func MakeCoins(copper, silver, gold int) Coins {
	return Coins(copper + 100*silver + 100000*gold)
}

// Copper returns the number of copper coins in the given collection.
func (c Coins) Copper() int {
	// Number of copper left after removing the gold + silver.
	return int(c % 100)
}

// Silver returns the number of silver coins in the given collection.
func (c Coins) Silver() int {
	// Number of silver left after removing the gold.
	return int(c/100.0) % 100
}

// Gold returns the number of gold coins in the given collection.
func (c Coins) Gold() int {
	return int(c / 100000.0)
}

// SpellSchool information.
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
type SpellCategory uint8

// SpellCategory values.
const (
	SpellCategoryHealthManaPotions SpellCategory = 4
	SpellCategoryDevourMagic       SpellCategory = 12
)

// Gender information.
type Gender uint8

// Gender values.
const (
	GenderMale   Gender = 0
	GenderFemale Gender = 1
	GenderNone   Gender = 2
)

// TypeID is a enum value which maps to a object's type ID.
type TypeID int

// TypeID values.
const (
	TypeIDObject        TypeID = 0
	TypeIDItem          TypeID = 1
	TypeIDContainer     TypeID = 2
	TypeIDUnit          TypeID = 3
	TypeIDPlayer        TypeID = 4
	TypeIDGameObject    TypeID = 5
	TypeIDDynamicObject TypeID = 6
	TypeIDCorpse        TypeID = 7
)

// TypeMask is a enum value which maps to a object's type ID.
type TypeMask int

// TypeMask values.
const (
	TypeMaskObject        TypeMask = 0x0001
	TypeMaskItem          TypeMask = 0x0002
	TypeMaskContainer     TypeMask = 0x0004
	TypeMaskUnit          TypeMask = 0x0008
	TypeMaskPlayer        TypeMask = 0x0010
	TypeMaskGameObject    TypeMask = 0x0020
	TypeMaskDynamicObject TypeMask = 0x0040
	TypeMaskCorpse        TypeMask = 0x0080
)

// UpdateFlags is a enum value which maps to a object's type ID.
type UpdateFlags int

// UpdateFlags values.
const (
	UpdateFlagsNone        UpdateFlags = 0x0000
	UpdateFlagsSelf        UpdateFlags = 0x0001
	UpdateFlagsTransport   UpdateFlags = 0x0002
	UpdateFlagsFullGUID    UpdateFlags = 0x0004
	UpdateFlagsHighGUID    UpdateFlags = 0x0008
	UpdateFlagsAll         UpdateFlags = 0x0010
	UpdateFlagsLiving      UpdateFlags = 0x0020
	UpdateFlagsHasPosition UpdateFlags = 0x0040
)

// HighGUID is a enum value which maps to a object's type ID.
type HighGUID uint32

// HighGUID values.
const (
	HighGUIDItem          HighGUID = 0x40000000
	HighGUIDContainer     HighGUID = 0x40000000
	HighGUIDPlayer        HighGUID = 0x00000000
	HighGUIDGameobject    HighGUID = 0xF1100000
	HighGUIDTransport     HighGUID = 0xF1200000
	HighGUIDUnit          HighGUID = 0xF1300000
	HighGUIDPet           HighGUID = 0xF1400000
	HighGUIDDynamicObject HighGUID = 0xF1000000
	HighGUIDCorpse        HighGUID = 0xF1010000
	HighGUIDMoTransport   HighGUID = 0x1FC00000
)

// EquipmentSlot is a enum value which maps to a object's type ID.
type EquipmentSlot uint32

// EquipmentSlot values.
const (
	EquipmentSlotHead      EquipmentSlot = 0
	EquipmentSlotNeck      EquipmentSlot = 1
	EquipmentSlotShoulders EquipmentSlot = 2
	EquipmentSlotBody      EquipmentSlot = 3
	EquipmentSlotChest     EquipmentSlot = 4
	EquipmentSlotWaist     EquipmentSlot = 5
	EquipmentSlotLegs      EquipmentSlot = 6
	EquipmentSlotFeet      EquipmentSlot = 7
	EquipmentSlotWrists    EquipmentSlot = 8
	EquipmentSlotHands     EquipmentSlot = 9
	EquipmentSlotFinger1   EquipmentSlot = 10
	EquipmentSlotFinger2   EquipmentSlot = 11
	EquipmentSlotTrinket1  EquipmentSlot = 12
	EquipmentSlotTrinket2  EquipmentSlot = 13
	EquipmentSlotBack      EquipmentSlot = 14
	EquipmentSlotMainHand  EquipmentSlot = 15
	EquipmentSlotOffHand   EquipmentSlot = 16
	EquipmentSlotRanged    EquipmentSlot = 17
	EquipmentSlotTabard    EquipmentSlot = 18
)

// CharacterFlag is a enum value which maps to a object's type ID.
type CharacterFlag uint32

// CharacterFlag values.
const (
	CharacterFlagNone              CharacterFlag = 0x00000000
	CharacterFlagLockedForTransfer CharacterFlag = 0x00000004
	CharacterFlagHideHelm          CharacterFlag = 0x00000400
	CharacterFlagHideCloak         CharacterFlag = 0x00000800
	CharacterFlagGhost             CharacterFlag = 0x00002000
	CharacterFlagRename            CharacterFlag = 0x00004000
	CharacterFlagLockedByBilling   CharacterFlag = 0x01000000
	CharacterFlagDeclined          CharacterFlag = 0x02000000
)

// PlayerFlag is a enum value which maps to a object's type ID.
type PlayerFlag uint32

// PlayerFlag values.
const (
	PlayerFlagsGroupLeader     PlayerFlag = 0x00000001
	PlayerFlagsAFK             PlayerFlag = 0x00000002
	PlayerFlagsDND             PlayerFlag = 0x00000004
	PlayerFlagsGM              PlayerFlag = 0x00000008
	PlayerFlagsGhost           PlayerFlag = 0x00000010
	PlayerFlagsResting         PlayerFlag = 0x00000020
	PlayerFlagsFFAPVP          PlayerFlag = 0x00000080
	PlayerFlagsContestedPVP    PlayerFlag = 0x00000100
	PlayerFlagsInPVP           PlayerFlag = 0x00000200
	PlayerFlagsHideHelm        PlayerFlag = 0x00000400
	PlayerFlagsHideCloak       PlayerFlag = 0x00000800
	PlayerFlagsPartialPlayTime PlayerFlag = 0x00001000
	PlayerFlagsNoPlayTime      PlayerFlag = 0x00002000
	PlayerFlagsSanctuary       PlayerFlag = 0x00010000
	PlayerFlagsTaxiBenchmark   PlayerFlag = 0x00020000
	PlayerFlagsPVPTimer        PlayerFlag = 0x00040000
)

// PlayerBytes is a enum value which maps to a object's type ID.
type PlayerBytes uint32

// PlayerBytes values.
const (
	PlayerBytesTrackStealthed  PlayerBytes = 0x02
	PlayerBytesReleaseTimer    PlayerBytes = 0x08
	PlayerBytesNoReleaseWindow PlayerBytes = 0x10
)

// UpdateField is a enum value which maps to a object's type ID.
type UpdateField uint32

// UpdateField values.
const (
	UpdateFieldGUIDLow  UpdateField = 0
	UpdateFieldGUIDHigh UpdateField = 1
	UpdateFieldType     UpdateField = 2
	UpdateFieldEntry    UpdateField = 3
	UpdateFieldScaleX   UpdateField = 4

	UpdateFieldItemOwner               UpdateField = 6
	UpdateFieldItemContained           UpdateField = 8
	UpdateFieldItemCreator             UpdateField = 10
	UpdateFieldItemGiftCreator         UpdateField = 12
	UpdateFieldItemStackCount          UpdateField = 14
	UpdateFieldItemDuration            UpdateField = 15
	UpdateFieldItemSpellCharges        UpdateField = 16
	UpdateFieldItemFlags               UpdateField = 21
	UpdateFieldItemEnchantmentID       UpdateField = 22
	UpdateFieldItemEnchantmentDuration UpdateField = 23
	UpdateFieldItemEnchantmentCharges  UpdateField = 24
	UpdateFieldItemPropertySeed        UpdateField = 43
	UpdateFieldItemRandomPropertiesID  UpdateField = 44
	UpdateFieldItemItemTextID          UpdateField = 45
	UpdateFieldItemDurability          UpdateField = 46
	UpdateFieldItemMaxDurability       UpdateField = 47

	UpdateFieldContainerNumSlots UpdateField = 48
	UpdateFieldContainerSlot1    UpdateField = 50
	UpdateFieldContainerSlotLast UpdateField = 104

	UpdateFieldUnitCharmLow                    UpdateField = 6
	UpdateFieldUnitCharmHigh                   UpdateField = 7
	UpdateFieldUnitSummonLow                   UpdateField = 8
	UpdateFieldUnitSummonHigh                  UpdateField = 9
	UpdateFieldUnitCharmedbyLow                UpdateField = 10
	UpdateFieldUnitCharmedbyHigh               UpdateField = 11
	UpdateFieldUnitSummonedbyLow               UpdateField = 12
	UpdateFieldUnitSummonedbyHigh              UpdateField = 13
	UpdateFieldUnitCreatedbyLow                UpdateField = 14
	UpdateFieldUnitCreatedbyHigh               UpdateField = 15
	UpdateFieldUnitTargetLow                   UpdateField = 16
	UpdateFieldUnitTargetHigh                  UpdateField = 17
	UpdateFieldUnitPersuadedLow                UpdateField = 18
	UpdateFieldUnitPersuadedHigh               UpdateField = 19
	UpdateFieldUnitChannelObjectLow            UpdateField = 20
	UpdateFieldUnitChannelObjectHigh           UpdateField = 21
	UpdateFieldUnitHealth                      UpdateField = 22
	UpdateFieldUnitPowerStart                  UpdateField = 23
	UpdateFieldUnitMaxHealth                   UpdateField = 28
	UpdateFieldUnitMaxPowerStart               UpdateField = 29
	UpdateFieldUnitLevel                       UpdateField = 34
	UpdateFieldUnitFactiontemplate             UpdateField = 35
	UpdateFieldUnitBytes0                      UpdateField = 36
	UpdateFieldUnitVirtualItemDisplay          UpdateField = 37
	UpdateFieldUnitVirtualItemInfo             UpdateField = 40
	UpdateFieldUnitFlags                       UpdateField = 46
	UpdateFieldUnitAura                        UpdateField = 47
	UpdateFieldUnitAuraLast                    UpdateField = 94
	UpdateFieldUnitAuraflags                   UpdateField = 95
	UpdateFieldUnitAuraflags01                 UpdateField = 96
	UpdateFieldUnitAuraflags02                 UpdateField = 97
	UpdateFieldUnitAuraflags03                 UpdateField = 98
	UpdateFieldUnitAuraflags04                 UpdateField = 99
	UpdateFieldUnitAuraflags05                 UpdateField = 100
	UpdateFieldUnitAuralevels                  UpdateField = 101
	UpdateFieldUnitAuralevelsLast              UpdateField = 112
	UpdateFieldUnitAuraapplications            UpdateField = 113
	UpdateFieldUnitAuraapplicationsLast        UpdateField = 124
	UpdateFieldUnitAurastate                   UpdateField = 125
	UpdateFieldUnitBaseattacktime              UpdateField = 126
	UpdateFieldUnitOffhandattacktime           UpdateField = 127
	UpdateFieldUnitRangedattacktime            UpdateField = 128
	UpdateFieldUnitBoundingradius              UpdateField = 129
	UpdateFieldUnitCombatreach                 UpdateField = 130
	UpdateFieldUnitDisplayid                   UpdateField = 131
	UpdateFieldUnitNativedisplayid             UpdateField = 132
	UpdateFieldUnitMountdisplayid              UpdateField = 133
	UpdateFieldUnitMindamage                   UpdateField = 134
	UpdateFieldUnitMaxdamage                   UpdateField = 135
	UpdateFieldUnitMinoffhanddamage            UpdateField = 136
	UpdateFieldUnitMaxoffhanddamage            UpdateField = 137
	UpdateFieldUnitBytes1                      UpdateField = 138
	UpdateFieldUnitPetnumber                   UpdateField = 139
	UpdateFieldUnitPetNameTimestamp            UpdateField = 140
	UpdateFieldUnitPetexperience               UpdateField = 141
	UpdateFieldUnitPetnextlevelexp             UpdateField = 142
	UpdateFieldUnitDynamicFlags                UpdateField = 143
	UpdateFieldUnitChannelSpell                UpdateField = 144
	UpdateFieldUnitModCastSpeed                UpdateField = 145
	UpdateFieldUnitCreatedBySpell              UpdateField = 146
	UpdateFieldUnitNpcFlags                    UpdateField = 147
	UpdateFieldUnitNpcEmotestate               UpdateField = 148
	UpdateFieldUnitTrainingPoints              UpdateField = 149
	UpdateFieldUnitStrength                    UpdateField = 150
	UpdateFieldUnitAgility                     UpdateField = 151
	UpdateFieldUnitStamina                     UpdateField = 152
	UpdateFieldUnitIntellect                   UpdateField = 153
	UpdateFieldUnitSpirit                      UpdateField = 154
	UpdateFieldUnitArmor                       UpdateField = 155
	UpdateFieldUnitHolyResist                  UpdateField = 156
	UpdateFieldUnitFireResist                  UpdateField = 157
	UpdateFieldUnitNatureResist                UpdateField = 158
	UpdateFieldUnitFrostResist                 UpdateField = 159
	UpdateFieldUnitShadowResist                UpdateField = 160
	UpdateFieldUnitArcaneResist                UpdateField = 161
	UpdateFieldUnitBaseMana                    UpdateField = 162
	UpdateFieldUnitBaseHealth                  UpdateField = 163
	UpdateFieldUnitBytes2                      UpdateField = 164
	UpdateFieldUnitAttackPower                 UpdateField = 165
	UpdateFieldUnitAttackPowerMods             UpdateField = 166
	UpdateFieldUnitAttackPowerMultiplier       UpdateField = 167
	UpdateFieldUnitRangedAttackPower           UpdateField = 168
	UpdateFieldUnitRangedAttackPowerMods       UpdateField = 169
	UpdateFieldUnitRangedAttackPowerMultiplier UpdateField = 170
	UpdateFieldUnitMinrangeddamage             UpdateField = 171
	UpdateFieldUnitMaxrangeddamage             UpdateField = 172
	UpdateFieldUnitPowerCostModifier           UpdateField = 173
	UpdateFieldUnitPowerCostModifier01         UpdateField = 174
	UpdateFieldUnitPowerCostModifier02         UpdateField = 175
	UpdateFieldUnitPowerCostModifier03         UpdateField = 176
	UpdateFieldUnitPowerCostModifier04         UpdateField = 177
	UpdateFieldUnitPowerCostModifier05         UpdateField = 178
	UpdateFieldUnitPowerCostModifier06         UpdateField = 179
	UpdateFieldUnitPowerCostMultiplier         UpdateField = 180
	UpdateFieldUnitPowerCostMultiplier01       UpdateField = 181
	UpdateFieldUnitPowerCostMultiplier02       UpdateField = 182
	UpdateFieldUnitPowerCostMultiplier03       UpdateField = 183
	UpdateFieldUnitPowerCostMultiplier04       UpdateField = 184
	UpdateFieldUnitPowerCostMultiplier05       UpdateField = 185
	UpdateFieldUnitPowerCostMultiplier06       UpdateField = 186
	UpdateFieldUnitPadding                     UpdateField = 187

	UpdateFieldPlayerDuelArbiter                UpdateField = 188
	UpdateFieldPlayerFlags                      UpdateField = 190
	UpdateFieldPlayerGuildid                    UpdateField = 191
	UpdateFieldPlayerGuildrank                  UpdateField = 192
	UpdateFieldPlayerBytes                      UpdateField = 193
	UpdateFieldPlayerBytes2                     UpdateField = 194
	UpdateFieldPlayerBytes3                     UpdateField = 195
	UpdateFieldPlayerDuelTeam                   UpdateField = 196
	UpdateFieldPlayerGuildTimestamp             UpdateField = 197
	UpdateFieldPlayerQuestStart                 UpdateField = 198
	UpdateFieldPlayerQuestLog11                 UpdateField = 198
	UpdateFieldPlayerQuestLog12                 UpdateField = 199
	UpdateFieldPlayerQuestLog13                 UpdateField = 200
	UpdateFieldPlayerQuestLogLast1              UpdateField = 255
	UpdateFieldPlayerQuestLogLast2              UpdateField = 256
	UpdateFieldPlayerQuestLogLast3              UpdateField = 257
	UpdateFieldPlayerVisibleItem1Creator        UpdateField = 258
	UpdateFieldPlayerVisibleItemEntryStart      UpdateField = 260
	UpdateFieldPlayerVisibleItem1Ench           UpdateField = 261
	UpdateFieldPlayerVisibleItem1Properties     UpdateField = 268
	UpdateFieldPlayerVisibleItem1Pad            UpdateField = 269
	UpdateFieldPlayerVisibleItemLastCreator     UpdateField = 474
	UpdateFieldPlayerVisibleItemLast0           UpdateField = 476
	UpdateFieldPlayerVisibleItemLastProperties  UpdateField = 484
	UpdateFieldPlayerVisibleItemLastPad         UpdateField = 485
	UpdateFieldPlayerInventoryStart             UpdateField = 486
	UpdateFieldPlayerBagStart                   UpdateField = 524
	UpdateFieldPlayerPackSlot1                  UpdateField = 532
	UpdateFieldPlayerPackSlotLast               UpdateField = 562
	UpdateFieldPlayerBankSlot1                  UpdateField = 564
	UpdateFieldPlayerBankSlotLast               UpdateField = 610
	UpdateFieldPlayerBankbagSlot1               UpdateField = 612
	UpdateFieldPlayerBankbagSlotLast            UpdateField = 622
	UpdateFieldPlayerVendorbuybackSlot1         UpdateField = 624
	UpdateFieldPlayerVendorbuybackSlotLast      UpdateField = 646
	UpdateFieldPlayerKeyringSlot1               UpdateField = 648
	UpdateFieldPlayerKeyringSlotLast            UpdateField = 710
	UpdateFieldPlayerFarsight                   UpdateField = 712
	UpdateFieldPlayerComboTarget                UpdateField = 714
	UpdateFieldPlayerXp                         UpdateField = 716
	UpdateFieldPlayerNextLevelXp                UpdateField = 717
	UpdateFieldPlayerSkillInfo11                UpdateField = 718
	UpdateFieldPlayerCharacterPoints1           UpdateField = 1102
	UpdateFieldPlayerCharacterPoints2           UpdateField = 1103
	UpdateFieldPlayerTrackCreatures             UpdateField = 1104
	UpdateFieldPlayerTrackResources             UpdateField = 1105
	UpdateFieldPlayerBlockPercentage            UpdateField = 1106
	UpdateFieldPlayerDodgePercentage            UpdateField = 1107
	UpdateFieldPlayerParryPercentage            UpdateField = 1108
	UpdateFieldPlayerCritPercentage             UpdateField = 1109
	UpdateFieldPlayerRangedCritPercentage       UpdateField = 1110
	UpdateFieldPlayerExploredZones1             UpdateField = 1111
	UpdateFieldPlayerRestStateExperience        UpdateField = 1175
	UpdateFieldPlayerCoinage                    UpdateField = 1176
	UpdateFieldPlayerPosstat0                   UpdateField = 1177
	UpdateFieldPlayerPosstat1                   UpdateField = 1178
	UpdateFieldPlayerPosstat2                   UpdateField = 1179
	UpdateFieldPlayerPosstat3                   UpdateField = 1180
	UpdateFieldPlayerPosstat4                   UpdateField = 1181
	UpdateFieldPlayerNegstat0                   UpdateField = 1182
	UpdateFieldPlayerNegstat1                   UpdateField = 1183
	UpdateFieldPlayerNegstat2                   UpdateField = 1184
	UpdateFieldPlayerNegstat3                   UpdateField = 1185
	UpdateFieldPlayerNegstat4                   UpdateField = 1186
	UpdateFieldPlayerResistancebuffmodspositive UpdateField = 1187
	UpdateFieldPlayerResistancebuffmodsnegative UpdateField = 1194
	UpdateFieldPlayerModDamageDonePos           UpdateField = 1201
	UpdateFieldPlayerModDamageDoneNeg           UpdateField = 1208
	UpdateFieldPlayerModDamageDonePct           UpdateField = 1215
	UpdateFieldPlayerFieldBytes                 UpdateField = 1222
	UpdateFieldPlayerAmmoID                     UpdateField = 1223
	UpdateFieldPlayerSelfResSpell               UpdateField = 1224
	UpdateFieldPlayerPvpMedals                  UpdateField = 1225
	UpdateFieldPlayerBuybackPrice1              UpdateField = 1226
	UpdateFieldPlayerBuybackPriceLast           UpdateField = 1237
	UpdateFieldPlayerBuybackTimestamp1          UpdateField = 1238
	UpdateFieldPlayerBuybackTimestampLast       UpdateField = 1249
	UpdateFieldPlayerSessionKills               UpdateField = 1250
	UpdateFieldPlayerYesterdayKills             UpdateField = 1251
	UpdateFieldPlayerLastWeekKills              UpdateField = 1252
	UpdateFieldPlayerThisWeekKills              UpdateField = 1253
	UpdateFieldPlayerThisWeekContribution       UpdateField = 1254
	UpdateFieldPlayerLifetimeHonorableKills     UpdateField = 1255
	UpdateFieldPlayerLifetimeDishonorableKills  UpdateField = 1256
	UpdateFieldPlayerYesterdayContribution      UpdateField = 1257
	UpdateFieldPlayerLastWeekContribution       UpdateField = 1258
	UpdateFieldPlayerLastWeekRank               UpdateField = 1259
	UpdateFieldPlayerBytes2b                    UpdateField = 1260
	UpdateFieldPlayerWatchedFactionIndex        UpdateField = 1261
	UpdateFieldPlayerCombatRating1              UpdateField = 1262
)

// ItemFlag is a enum value which maps to a object's type ID.
type ItemFlag uint32

// ItemFlag values.
const (
	ItemFlagNone     ItemFlag = 0x00000000
	ItemFlagBound    ItemFlag = 0x00000001
	ItemFlagUnlocked ItemFlag = 0x00000004
	ItemFlagWrapped  ItemFlag = 0x00000008
	ItemFlagReadable ItemFlag = 0x00000200
)

// ItemPrototypeFlag is a enum value which maps to a object's type ID.
type ItemPrototypeFlag uint32

// ItemPrototypeFlag values.
const (
	ItemPrototypeFlagConjured        ItemPrototypeFlag = 0x00000002
	ItemPrototypeFlagLootable        ItemPrototypeFlag = 0x00000004
	ItemPrototypeFlagWrapped         ItemPrototypeFlag = 0x00000008
	ItemPrototypeFlagDeprecated      ItemPrototypeFlag = 0x00000010
	ItemPrototypeFlagIndestructible  ItemPrototypeFlag = 0x00000020
	ItemPrototypeFlagUsable          ItemPrototypeFlag = 0x00000040
	ItemPrototypeFlagNoEquipCooldown ItemPrototypeFlag = 0x00000080
	ItemPrototypeFlagWrapper         ItemPrototypeFlag = 0x00000200
	ItemPrototypeFlagStackable       ItemPrototypeFlag = 0x00000400
	ItemPrototypeFlagPartyLoot       ItemPrototypeFlag = 0x00000800
	ItemPrototypeFlagCharter         ItemPrototypeFlag = 0x00002000
	ItemPrototypeFlagLetter          ItemPrototypeFlag = 0x00004000
	ItemPrototypeFlagPVPReward       ItemPrototypeFlag = 0x00008000
	ItemPrototypeFlagUniqueEquipped  ItemPrototypeFlag = 0x00080000
)

// Power is a enum value which maps to a object's type ID.
type Power uint32

// Power values.
const (
	PowerMana      Power = 0
	PowerRage      Power = 1
	PowerFocus     Power = 2
	PowerEnergy    Power = 3
	PowerHappiness Power = 4
)

// Team is a enum value which maps to a object's type ID.
type Team uint32

// Team values.
const (
	TeamAlliance Team = 67
	TeamHorde    Team = 469
)

// StandState is a enum value which maps to a object's type ID.
type StandState uint32

// StandState values.
const (
	StandStateStand          StandState = 0
	StandStateSit            StandState = 1
	StandStateSitChair       StandState = 2
	StandStateSleep          StandState = 3
	StandStateSitLowChair    StandState = 4
	StandStateSitMediumChair StandState = 5
	StandStateSitHighChair   StandState = 6
	StandStateDead           StandState = 7
	StandStateKneel          StandState = 8
)

// Byte1Flags is a enum value which maps to a object's type ID.
type Byte1Flags uint32

// Byte1Flags values.
const (
	Byte1FlagsAlwaysStand Byte1Flags = 0x01
	Byte1FlagsCreep       Byte1Flags = 0x02
	Byte1FlagsUntrackable Byte1Flags = 0x04
	Byte1FlagsAll         Byte1Flags = 0xFF
)

// Byte2Flags is a enum value which maps to a object's type ID.
type Byte2Flags uint32

// Byte2Flags values.
const (
	Byte2FlagsNone             Byte2Flags = 0x00
	Byte2FlagsDetectAmore0     Byte2Flags = 0x02
	Byte2FlagsDetectAmore1     Byte2Flags = 0x04
	Byte2FlagsDetectAmore2     Byte2Flags = 0x08
	Byte2FlagsDetectAmore3     Byte2Flags = 0x10
	Byte2FlagsStealth          Byte2Flags = 0x20
	Byte2FlagsInvisibilityGlow Byte2Flags = 0x40
)

// UpdateType is a enum value which maps to a object's type ID.
type UpdateType uint8

// Byte2Flags values.
const (
	UpdateTypeValues            UpdateType = 0
	UpdateTypeMovement          UpdateType = 1
	UpdateTypeCreateObject      UpdateType = 2
	UpdateTypeCreateObject2     UpdateType = 3
	UpdateTypeOutOfRangeObjects UpdateType = 4
	UpdateTypeNearObjects       UpdateType = 5
)

// MovementFlag is a enum value which maps to a object's type ID.
type MovementFlag uint32

// Byte2Flags values.
const (
	MovementFlagForward         MovementFlag = 0x00000001
	MovementFlagBackward        MovementFlag = 0x00000002
	MovementFlagStrafeLeft      MovementFlag = 0x00000004
	MovementFlagStrafeRight     MovementFlag = 0x00000008
	MovementFlagTurnLeft        MovementFlag = 0x00000010
	MovementFlagTurnRight       MovementFlag = 0x00000020
	MovementFlagPitchUp         MovementFlag = 0x00000040
	MovementFlagPitchDown       MovementFlag = 0x00000080
	MovementFlagWalkMode        MovementFlag = 0x00000100
	MovementFlagLevitating      MovementFlag = 0x00000400
	MovementFlagFlying          MovementFlag = 0x00000800
	MovementFlagFalling         MovementFlag = 0x00002000
	MovementFlagFallingFar      MovementFlag = 0x00004000
	MovementFlagSwimming        MovementFlag = 0x00200000
	MovementFlagSplineEnabled   MovementFlag = 0x00400000
	MovementFlagCanFly          MovementFlag = 0x00800000
	MovementFlagFlyingOld       MovementFlag = 0x01000000
	MovementFlagOnTransport     MovementFlag = 0x02000000
	MovementFlagSplineElevation MovementFlag = 0x04000000
	MovementFlagRoot            MovementFlag = 0x08000000
	MovementFlagWaterWalking    MovementFlag = 0x10000000
	MovementFlagSafeFall        MovementFlag = 0x20000000
	MovementFlagHover           MovementFlag = 0x40000000
)

// DisplayID is a enum value which maps to a object's type ID.
type DisplayID uint32

// DisplayID values.
const (
	DisplayIDInvBoots06  DisplayID = 10141
	DisplayIDInvPants02  DisplayID = 9892
	DisplayIDInvShield09 DisplayID = 18730
	DisplayIDInvShirt05  DisplayID = 9891
	DisplayIDInvSword04  DisplayID = 1542
)

type AttackTargetState uint8

const (
	AttackTargetStateIntact    AttackTargetState = 0 // set when attacker misses
	AttackTargetStateHit       AttackTargetState = 1 // victim got clear/blocked hit
	AttackTargetStateDodge     AttackTargetState = 2
	AttackTargetStateParry     AttackTargetState = 3
	AttackTargetStateInterrupt AttackTargetState = 4
	AttackTargetStateBlocks    AttackTargetState = 5 // unused? not set when blocked, even on full block
	AttackTargetStateEvades    AttackTargetState = 6
	AttackTargetStateIsImmune  AttackTargetState = 7
	AttackTargetStateDeflects  AttackTargetState = 8
)

type HitInfo uint32

const (
	HitInfoNormalSwing     HitInfo = 0x00000000
	HitInfoUnk1            HitInfo = 0x00000001 // req correct packet structure
	HitInfoAffectsVictim   HitInfo = 0x00000002
	HitInfoOffhand         HitInfo = 0x00000004
	HitInfoUnk2            HitInfo = 0x00000008
	HitInfoMiss            HitInfo = 0x00000010
	HitInfoFullAbsorb      HitInfo = 0x00000020
	HitInfoPartialAbsorb   HitInfo = 0x00000040
	HitInfoFullResist      HitInfo = 0x00000080
	HitInfoPartialResist   HitInfo = 0x00000100
	HitInfoCriticalhit     HitInfo = 0x00000200 // critical hit
	HitInfoUnk10           HitInfo = 0x00000400
	HitInfoUnk11           HitInfo = 0x00000800
	HitInfoUnk12           HitInfo = 0x00001000
	HitInfoBlock           HitInfo = 0x00002000 // blocked damage
	HitInfoUnk14           HitInfo = 0x00004000 // set only if meleespellid is present// no world text when victim is hit for 0 dmg(HideWorldTextForNoDamage?)
	HitInfoUnk15           HitInfo = 0x00008000 // player victim?// something related to blod sprut visual (BloodSpurtInBack?)
	HitInfoGlancing        HitInfo = 0x00010000
	HitInfoCrushing        HitInfo = 0x00020000
	HitInfoNoAnimation     HitInfo = 0x00040000
	HitInfoUnk19           HitInfo = 0x00080000
	HitInfoUnk20           HitInfo = 0x00100000
	HitInfoSwingNoHitSound HitInfo = 0x00200000 // unused?
	HitInfoUnk22           HitInfo = 0x00400000
	HitInfoRageGain        HitInfo = 0x00800000
	HitInfoFakeDamage      HitInfo = 0x01000000 // enables damage animation even if no damage done, set only if no damage
)
