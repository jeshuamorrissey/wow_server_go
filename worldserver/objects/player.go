package objects

import (
	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// Player represents the game object for a player-controlled character.
type Player struct {
	Unit

	SkinColor uint8
	Face      uint8
	HairStyle uint8
	HairColor uint8
	Feature   uint8

	ZoneID int
	MapID  int

	Equipment map[c.EquipmentSlot]*Item
	Inventory map[int]*Item
	Bags      map[int]*Container

	DrunkValue int
	XP         int
	Money      int

	// Flags.
	IsGroupLeader     bool
	IsAFK             bool
	IsDND             bool
	IsGM              bool
	IsGhost           bool
	IsResting         bool
	IsFFAPVP          bool
	IsContestedPVP    bool
	IsInPVP           bool
	HideHelm          bool
	HideCloak         bool
	IsPartialPlayTime bool
	IsNoPlayTime      bool
	IsInSanctuary     bool
	IsTaxiBenchmark   bool
	IsPVPTimer        bool

	IsTrackStealthed           bool
	ShowAutoReleaseSpiritTimer bool
	HideReleaseSpirit          bool
}

// GUID returns the guid of the object.
func (o *Player) GUID() GUID { return o.Unit.GUID() }

// SetGUID updates the GUID value of the object.
func (o *Player) SetGUID(guid int) { o.guid = GUID(int(c.HighGUIDPlayer)<<32 | guid) }

// HighGUID returns the high GUID component for an object.
func (o *Player) HighGUID() c.HighGUID { return c.HighGUIDPlayer }

// GetLocation returns the location of the object.
func (o *Player) GetLocation() *Location { return &o.Location }

// MovementUpdate returns a bytes representation of a movement update.
func (o *Player) MovementUpdate() []byte { return o.Unit.MovementUpdate() }

// NumFields returns the number of fields available for this object.
func (o *Player) NumFields() int { return 1282 }

// Fields returns the update fields of the object.
func (o *Player) Fields() map[c.UpdateField]interface{} {
	modelInfo := data.GetPlayerModelInfo(o.Race, o.Gender)

	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldUnitCharmLow:                                     uint32(o.Charm.Low()),
		c.UpdateFieldUnitCharmHigh:                                    uint32(o.Charm.High()),
		c.UpdateFieldUnitSummonLow:                                    uint32(o.Summon.Low()),
		c.UpdateFieldUnitSummonHigh:                                   uint32(o.Summon.High()),
		c.UpdateFieldUnitCharmedbyLow:                                 uint32(o.CharmedBy.Low()),
		c.UpdateFieldUnitCharmedbyHigh:                                uint32(o.CharmedBy.High()),
		c.UpdateFieldUnitSummonedbyLow:                                uint32(o.SummonedBy.Low()),
		c.UpdateFieldUnitSummonedbyHigh:                               uint32(o.SummonedBy.High()),
		c.UpdateFieldUnitCreatedbyLow:                                 uint32(o.CreatedBy.Low()),
		c.UpdateFieldUnitCreatedbyHigh:                                uint32(o.CreatedBy.High()),
		c.UpdateFieldUnitTargetLow:                                    uint32(o.Target.Low()),
		c.UpdateFieldUnitTargetHigh:                                   uint32(o.Target.High()),
		c.UpdateFieldUnitPersuadedLow:                                 uint32(o.Persuaded.Low()),
		c.UpdateFieldUnitPersuadedHigh:                                uint32(o.Persuaded.High()),
		c.UpdateFieldUnitChannelObjectLow:                             0, // TODO
		c.UpdateFieldUnitChannelObjectHigh:                            0, // TODO
		c.UpdateFieldUnitHealth:                                       uint32(o.Health),
		c.UpdateFieldUnitPowerStart + c.UpdateField(o.powerType()):    uint32(o.Power),
		c.UpdateFieldUnitMaxHealth:                                    uint32(o.maxHealth()),
		c.UpdateFieldUnitMaxPowerStart + c.UpdateField(o.powerType()): uint32(o.maxPower()),
		c.UpdateFieldUnitLevel:                                        uint32(o.Level),
		c.UpdateFieldUnitBytes0:                                       uint32(o.Race) | uint32(o.Class)<<8 | uint32(o.Gender)<<16,
		c.UpdateFieldUnitAura:                                         0, // TODO
		c.UpdateFieldUnitAuraLast:                                     0, // TODO
		c.UpdateFieldUnitAuraflags:                                    0, // TODO
		c.UpdateFieldUnitAuraflags01:                                  0, // TODO
		c.UpdateFieldUnitAuraflags02:                                  0, // TODO
		c.UpdateFieldUnitAuraflags03:                                  0, // TODO
		c.UpdateFieldUnitAuraflags04:                                  0, // TODO
		c.UpdateFieldUnitAuraflags05:                                  0, // TODO
		c.UpdateFieldUnitAuralevels:                                   0, // TODO
		c.UpdateFieldUnitAuralevelsLast:                               0, // TODO
		c.UpdateFieldUnitAuraapplications:                             0, // TODO
		c.UpdateFieldUnitAuraapplicationsLast:                         0, // TODO
		c.UpdateFieldUnitAurastate:                                    0, // TODO
		c.UpdateFieldUnitBaseattacktime:                               uint32(1000),
		c.UpdateFieldUnitBoundingradius:                               uint32(modelInfo.BoundingRadius),
		c.UpdateFieldUnitCombatreach:                                  uint32(modelInfo.CombatReach),
		c.UpdateFieldUnitDisplayid:                                    uint32(modelInfo.ID),
		c.UpdateFieldUnitNativedisplayid:                              uint32(modelInfo.ID),
		c.UpdateFieldUnitMountdisplayid:                               0, // TODO
		c.UpdateFieldUnitBytes1:                                       uint32(o.Byte1Flags)<<24 | uint32(o.FreeTalentPoints)<<16 | uint32(o.StandState),
		c.UpdateFieldUnitPetnumber:                                    0, // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             0, // TODO
		c.UpdateFieldUnitPetexperience:                                0, // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              0, // TODO
		c.UpdateFieldUnitDynamicFlags:                                 0, // TODO
		c.UpdateFieldUnitChannelSpell:                                 0, // TODO
		c.UpdateFieldUnitModCastSpeed:                                 float32(1.0),
		c.UpdateFieldUnitCreatedBySpell:                               0, // TODO
		c.UpdateFieldUnitNpcFlags:                                     0, // TODO
		c.UpdateFieldUnitNpcEmotestate:                                uint32(o.EmoteState),
		c.UpdateFieldUnitTrainingPoints:                               uint32(o.TrainingPoints),
		c.UpdateFieldUnitStrength:                                     uint32(o.Strength),
		c.UpdateFieldUnitAgility:                                      uint32(o.Agility),
		c.UpdateFieldUnitStamina:                                      uint32(o.Stamina),
		c.UpdateFieldUnitIntellect:                                    uint32(o.Intellect),
		c.UpdateFieldUnitSpirit:                                       uint32(o.Spirit),
		c.UpdateFieldUnitArmor:                                        uint32(o.Resistances[c.SpellSchoolPhysical]),
		c.UpdateFieldUnitHolyResist:                                   uint32(o.Resistances[c.SpellSchoolHoly]),
		c.UpdateFieldUnitFireResist:                                   uint32(o.Resistances[c.SpellSchoolFire]),
		c.UpdateFieldUnitNatureResist:                                 uint32(o.Resistances[c.SpellSchoolNature]),
		c.UpdateFieldUnitFrostResist:                                  uint32(o.Resistances[c.SpellSchoolFrost]),
		c.UpdateFieldUnitShadowResist:                                 uint32(o.Resistances[c.SpellSchoolShadow]),
		c.UpdateFieldUnitArcaneResist:                                 uint32(o.Resistances[c.SpellSchoolArcane]),
		c.UpdateFieldUnitBaseMana:                                     uint32(o.BasePower),
		c.UpdateFieldUnitBaseHealth:                                   uint32(o.BaseHealth),
		c.UpdateFieldUnitBytes2:                                       0, // TODO
		c.UpdateFieldUnitAttackPower:                                  0, // TODO
		c.UpdateFieldUnitAttackPowerMods:                              0, // TODO
		c.UpdateFieldUnitAttackPowerMultiplier:                        0, // TODO
		c.UpdateFieldUnitRangedAttackPower:                            0, // TODO
		c.UpdateFieldUnitRangedAttackPowerMods:                        0, // TODO
		c.UpdateFieldUnitRangedAttackPowerMultiplier:                  0, // TODO
		c.UpdateFieldUnitMinrangeddamage:                              0, // TODO
		c.UpdateFieldUnitMaxrangeddamage:                              0, // TODO
		c.UpdateFieldUnitPowerCostModifier:                            0, // TODO
		c.UpdateFieldUnitPowerCostModifier01:                          0, // TODO
		c.UpdateFieldUnitPowerCostModifier02:                          0, // TODO
		c.UpdateFieldUnitPowerCostModifier03:                          0, // TODO
		c.UpdateFieldUnitPowerCostModifier04:                          0, // TODO
		c.UpdateFieldUnitPowerCostModifier05:                          0, // TODO
		c.UpdateFieldUnitPowerCostModifier06:                          0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier:                          0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier01:                        0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier02:                        0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier03:                        0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier04:                        0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier05:                        0, // TODO
		c.UpdateFieldUnitPowerCostMultiplier06:                        0, // TODO

		c.UpdateFieldPlayerDuelArbiter:                0, // TODO
		c.UpdateFieldPlayerFlags:                      uint32(o.flags()),
		c.UpdateFieldPlayerGuildid:                    0, // TODO
		c.UpdateFieldPlayerGuildrank:                  0, // TODO
		c.UpdateFieldPlayerBytes:                      uint32(o.SkinColor) | uint32(o.Face)<<8 | uint32(o.HairStyle)<<16 | uint32(o.HairColor)<<24,
		c.UpdateFieldPlayerBytes2:                     uint32(o.Feature),
		c.UpdateFieldPlayerBytes3:                     uint32(o.Gender) | uint32(o.DrunkValue)&0xFFFE,
		c.UpdateFieldPlayerDuelTeam:                   0, // TODO
		c.UpdateFieldPlayerGuildTimestamp:             0, // TODO
		c.UpdateFieldPlayerQuestStart:                 0, // TODO
		c.UpdateFieldPlayerBankSlot1:                  0, // TODO
		c.UpdateFieldPlayerBankSlotLast:               0, // TODO
		c.UpdateFieldPlayerBankbagSlot1:               0, // TODO
		c.UpdateFieldPlayerBankbagSlotLast:            0, // TODO
		c.UpdateFieldPlayerVendorbuybackSlot1:         0, // TODO
		c.UpdateFieldPlayerVendorbuybackSlotLast:      0, // TODO
		c.UpdateFieldPlayerKeyringSlot1:               0, // TODO
		c.UpdateFieldPlayerKeyringSlotLast:            0, // TODO
		c.UpdateFieldPlayerFarsight:                   0, // TODO
		c.UpdateFieldPlayerComboTarget:                0, // TODO
		c.UpdateFieldPlayerXp:                         uint32(o.XP),
		c.UpdateFieldPlayerNextLevelXp:                0, // TODO
		c.UpdateFieldPlayerSkillInfo11:                0, // TODO
		c.UpdateFieldPlayerCharacterPoints1:           0, // TODO
		c.UpdateFieldPlayerCharacterPoints2:           0, // TODO
		c.UpdateFieldPlayerTrackCreatures:             0, // TODO
		c.UpdateFieldPlayerTrackResources:             0, // TODO
		c.UpdateFieldPlayerBlockPercentage:            0, // TODO
		c.UpdateFieldPlayerDodgePercentage:            0, // TODO
		c.UpdateFieldPlayerParryPercentage:            0, // TODO
		c.UpdateFieldPlayerCritPercentage:             0, // TODO
		c.UpdateFieldPlayerRangedCritPercentage:       0, // TODO
		c.UpdateFieldPlayerExploredZones1:             0, // TODO
		c.UpdateFieldPlayerRestStateExperience:        0, // TODO
		c.UpdateFieldPlayerCoinage:                    uint32(o.Money),
		c.UpdateFieldPlayerPosstat0:                   0, // TODO
		c.UpdateFieldPlayerPosstat1:                   0, // TODO
		c.UpdateFieldPlayerPosstat2:                   0, // TODO
		c.UpdateFieldPlayerPosstat3:                   0, // TODO
		c.UpdateFieldPlayerPosstat4:                   0, // TODO
		c.UpdateFieldPlayerNegstat0:                   0, // TODO
		c.UpdateFieldPlayerNegstat1:                   0, // TODO
		c.UpdateFieldPlayerNegstat2:                   0, // TODO
		c.UpdateFieldPlayerNegstat3:                   0, // TODO
		c.UpdateFieldPlayerNegstat4:                   0, // TODO
		c.UpdateFieldPlayerResistancebuffmodspositive: 0, // TODO
		c.UpdateFieldPlayerResistancebuffmodsnegative: 0, // TODO
		c.UpdateFieldPlayerModDamageDonePos:           0, // TODO
		c.UpdateFieldPlayerModDamageDoneNeg:           0, // TODO
		c.UpdateFieldPlayerModDamageDonePct:           0, // TODO
		c.UpdateFieldPlayerFieldBytes:                 0, // TODO
		c.UpdateFieldPlayerAmmoID:                     0, // TODO
		c.UpdateFieldPlayerSelfResSpell:               0, // TODO
		c.UpdateFieldPlayerPvpMedals:                  0, // TODO
		c.UpdateFieldPlayerBuybackPrice1:              0, // TODO
		c.UpdateFieldPlayerBuybackPriceLast:           0, // TODO
		c.UpdateFieldPlayerBuybackTimestamp1:          0, // TODO
		c.UpdateFieldPlayerBuybackTimestampLast:       0, // TODO
		c.UpdateFieldPlayerSessionKills:               0, // TODO
		c.UpdateFieldPlayerYesterdayKills:             0, // TODO
		c.UpdateFieldPlayerLastWeekKills:              0, // TODO
		c.UpdateFieldPlayerThisWeekKills:              0, // TODO
		c.UpdateFieldPlayerThisWeekContribution:       0, // TODO
		c.UpdateFieldPlayerLifetimeHonorableKills:     0, // TODO
		c.UpdateFieldPlayerLifetimeDishonorableKills:  0, // TODO
		c.UpdateFieldPlayerYesterdayContribution:      0, // TODO
		c.UpdateFieldPlayerLastWeekContribution:       0, // TODO
		c.UpdateFieldPlayerLastWeekRank:               0, // TODO
		c.UpdateFieldPlayerBytes2b:                    0, // TODO
		c.UpdateFieldPlayerWatchedFactionIndex:        0, // TODO
		c.UpdateFieldPlayerCombatRating1:              0, // TODO
	}

	for slot, item := range o.Equipment {
		slotField := c.UpdateFieldPlayerInventoryStart + c.UpdateField(slot*2)
		fields[slotField] = uint32(item.GUID().Low())
		fields[slotField+1] = uint32(item.GUID().High())

		visibleItemSlot := c.UpdateField(slot * 12)
		fields[c.UpdateFieldPlayerVisibleItemEntryStart+visibleItemSlot] = uint32(item.Template().Entry)

		if item.Creator != nil {
			fields[c.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot] = uint32(item.Creator.GUID().Low())
			fields[c.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot+1] = uint32(item.Creator.GUID().High())
		}

		if slot == c.EquipmentSlotMainHand {
			fields[c.UpdateFieldUnitBaseattacktime] = uint32(item.Template().AttackRate)
			fields[c.UpdateFieldUnitMindamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxdamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Max)
		} else if slot == c.EquipmentSlotOffHand {
			fields[c.UpdateFieldUnitOffhandattacktime] = uint32(item.Template().AttackRate)
			fields[c.UpdateFieldUnitMinoffhanddamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxoffhanddamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Max)
		} else if slot == c.EquipmentSlotRanged {
			fields[c.UpdateFieldUnitRangedattacktime] = uint32(item.Template().AttackRate)
			fields[c.UpdateFieldUnitMinrangeddamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxrangeddamage] = uint32(item.Template().Damage[c.SpellSchoolPhysical].Max)
		}
	}

	for slot, bag := range o.Bags {
		slotField := c.UpdateFieldPlayerBagStart + c.UpdateField(slot*2)
		fields[slotField] = uint32(bag.GUID().Low())
		fields[slotField+1] = uint32(bag.GUID().High())
	}

	for slot, item := range o.Inventory {
		slotField := c.UpdateFieldPlayerPackSlot1 + c.UpdateField(slot*2)
		fields[slotField] = uint32(item.GUID().Low())
		fields[slotField+1] = uint32(item.GUID().High())
	}

	return mergeUpdateFields(fields, o.BaseGameObject.Fields())

}

func (o *Player) flags() uint32 {
	var flags uint32
	if o.IsGroupLeader {
		flags |= uint32(c.PlayerFlagsGroupLeader)
	}
	if o.IsAFK {
		flags |= uint32(c.PlayerFlagsAFK)
	}
	if o.IsDND {
		flags |= uint32(c.PlayerFlagsDND)
	}
	if o.IsGM {
		flags |= uint32(c.PlayerFlagsGM)
	}
	if o.IsGhost {
		flags |= uint32(c.PlayerFlagsGhost)
	}
	if o.IsResting {
		flags |= uint32(c.PlayerFlagsResting)
	}
	if o.IsFFAPVP {
		flags |= uint32(c.PlayerFlagsFFAPVP)
	}
	if o.IsContestedPVP {
		flags |= uint32(c.PlayerFlagsContestedPVP)
	}
	if o.IsInPVP {
		flags |= uint32(c.PlayerFlagsInPVP)
	}
	if o.HideHelm {
		flags |= uint32(c.PlayerFlagsHideHelm)
	}
	if o.HideCloak {
		flags |= uint32(c.PlayerFlagsHideCloak)
	}
	if o.IsPartialPlayTime {
		flags |= uint32(c.PlayerFlagsPartialPlayTime)
	}
	if o.IsNoPlayTime {
		flags |= uint32(c.PlayerFlagsNoPlayTime)
	}
	if o.IsInSanctuary {
		flags |= uint32(c.PlayerFlagsSanctuary)
	}
	if o.IsTaxiBenchmark {
		flags |= uint32(c.PlayerFlagsTaxiBenchmark)
	}
	if o.IsPVPTimer {
		flags |= uint32(c.PlayerFlagsPVPTimer)
	}

	return flags
}

func (o *Player) bytes() uint32 {
	var bytes uint32
	if o.IsTrackStealthed {
		bytes |= uint32(c.PlayerBytesTrackStealthed)
	}
	if o.ShowAutoReleaseSpiritTimer {
		bytes |= uint32(c.PlayerBytesReleaseTimer)
	}
	if o.HideReleaseSpirit {
		bytes |= uint32(c.PlayerBytesNoReleaseWindow)
	}

	return bytes
}

func (o *Player) bytes2() int {
	var flags int
	if o.CanDetectAmore0 {
		flags |= int(c.Byte2FlagsDetectAmore0)
	}
	if o.CanDetectAmore1 {
		flags |= int(c.Byte2FlagsDetectAmore1)
	}
	if o.CanDetectAmore2 {
		flags |= int(c.Byte2FlagsDetectAmore2)
	}
	if o.CanDetectAmore3 {
		flags |= int(c.Byte2FlagsDetectAmore3)
	}
	if o.IsStealth {
		flags |= int(c.Byte2FlagsStealth)
	}
	if o.HasInvisibilityGlow {
		flags |= int(c.Byte2FlagsInvisibilityGlow)
	}

	return flags
}

// FirstBag returns the first bag in the player's inventory, or
// nil if there are no bags.
func (o *Player) FirstBag() *Container {
	for i := 0; i < c.NumBagSlots; i++ {
		if bag, ok := o.Bags[i]; ok {
			return bag
		}
	}

	return nil
}
