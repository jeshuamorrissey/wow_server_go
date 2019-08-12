package objects

import (
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
}

// GUID returns the guid of the object.
func (o *Player) GUID() GUID { return o.Unit.GUID() }

// SetGUID updates the GUID value of the object.
func (o *Player) SetGUID(guid int) { o.guid = GUID(int(c.HighGUIDPlayer)<<32 | guid) }

// HighGUID returns the high GUID component for an object.
func (o *Player) HighGUID() c.HighGUID { return c.HighGUIDPlayer }

// GetLocation returns the location of the object.
func (o *Player) GetLocation() *Location { return &o.Location }

// Fields returns the update fields of the object.
func (o *Player) Fields() map[c.UpdateField]interface{} {
	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldUnitCharmLow:                                     o.Charm.Low(),
		c.UpdateFieldUnitCharmHigh:                                    o.Charm.High(),
		c.UpdateFieldUnitSummonLow:                                    o.Summon.Low(),
		c.UpdateFieldUnitSummonHigh:                                   o.Summon.High(),
		c.UpdateFieldUnitCharmedbyLow:                                 o.CharmedBy.Low(),
		c.UpdateFieldUnitCharmedbyHigh:                                o.CharmedBy.High(),
		c.UpdateFieldUnitSummonedbyLow:                                o.SummonedBy.Low(),
		c.UpdateFieldUnitSummonedbyHigh:                               o.SummonedBy.High(),
		c.UpdateFieldUnitCreatedbyLow:                                 o.CreatedBy.Low(),
		c.UpdateFieldUnitCreatedbyHigh:                                o.CreatedBy.High(),
		c.UpdateFieldUnitTargetLow:                                    o.Target.Low(),
		c.UpdateFieldUnitTargetHigh:                                   o.Target.High(),
		c.UpdateFieldUnitPersuadedLow:                                 o.Persuaded.Low(),
		c.UpdateFieldUnitPersuadedHigh:                                o.Persuaded.High(),
		c.UpdateFieldUnitChannelObjectLow:                             0, // TODO
		c.UpdateFieldUnitChannelObjectHigh:                            0, // TODO
		c.UpdateFieldUnitHealth:                                       o.Health,
		c.UpdateFieldUnitPowerStart + c.UpdateField(o.powerType()):    o.Power,
		c.UpdateFieldUnitMaxHealth:                                    o.maxHealth(),
		c.UpdateFieldUnitMaxPowerStart + c.UpdateField(o.powerType()): o.maxPower(),
		c.UpdateFieldUnitLevel:                                        o.Level,
		c.UpdateFieldUnitBytes0:                                       int(o.Race) | int(o.Class)<<8 | int(o.Gender)<<16,
		c.UpdateFieldUnitFlags:                                        0, // TODO
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
		c.UpdateFieldUnitBaseattacktime:                               0, // TODO
		c.UpdateFieldUnitOffhandattacktime:                            0, // TODO
		c.UpdateFieldUnitRangedattacktime:                             0, // TODO
		c.UpdateFieldUnitBoundingradius:                               0, // TODO
		c.UpdateFieldUnitCombatreach:                                  0, // TODO
		c.UpdateFieldUnitDisplayid:                                    0, // TODO
		c.UpdateFieldUnitNativedisplayid:                              0, // TODO
		c.UpdateFieldUnitMountdisplayid:                               0, // TODO
		c.UpdateFieldUnitMindamage:                                    0, // TODO
		c.UpdateFieldUnitMaxdamage:                                    0, // TODO
		c.UpdateFieldUnitMinoffhanddamage:                             0, // TODO
		c.UpdateFieldUnitMaxoffhanddamage:                             0, // TODO
		c.UpdateFieldUnitBytes1:                                       int(o.Byte1Flags)<<24 | o.FreeTalentPoints<<16 | int(o.StandState),
		c.UpdateFieldUnitPetnumber:                                    0, // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             0, // TODO
		c.UpdateFieldUnitPetexperience:                                0, // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              0, // TODO
		c.UpdateFieldUnitDynamicFlags:                                 0, // TODO
		c.UpdateFieldUnitChannelSpell:                                 0, // TODO
		c.UpdateFieldUnitModCastSpeed:                                 1.0,
		c.UpdateFieldUnitCreatedBySpell:                               0, // TODO
		c.UpdateFieldUnitNpcFlags:                                     0, // TODO
		c.UpdateFieldUnitNpcEmotestate:                                o.EmoteState,
		c.UpdateFieldUnitTrainingPoints:                               o.TrainingPoints,
		c.UpdateFieldUnitStrength:                                     o.Strength,
		c.UpdateFieldUnitAgility:                                      o.Agility,
		c.UpdateFieldUnitStamina:                                      o.Stamina,
		c.UpdateFieldUnitIntellect:                                    o.Intellect,
		c.UpdateFieldUnitSpirit:                                       o.Spirit,
		c.UpdateFieldUnitArmor:                                        o.Resistances[c.SpellSchoolPhysical],
		c.UpdateFieldUnitHolyResist:                                   o.Resistances[c.SpellSchoolHoly],
		c.UpdateFieldUnitFireResist:                                   o.Resistances[c.SpellSchoolFire],
		c.UpdateFieldUnitNatureResist:                                 o.Resistances[c.SpellSchoolNature],
		c.UpdateFieldUnitFrostResist:                                  o.Resistances[c.SpellSchoolFrost],
		c.UpdateFieldUnitShadowResist:                                 o.Resistances[c.SpellSchoolShadow],
		c.UpdateFieldUnitArcaneResist:                                 o.Resistances[c.SpellSchoolArcane],
		c.UpdateFieldUnitBaseMana:                                     o.BasePower,
		c.UpdateFieldUnitBaseHealth:                                   o.BaseHealth,
		c.UpdateFieldUnitBytes2:                                       o.Unit.bytes2(),
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
		c.UpdateFieldPlayerFlags:                      0, // TODO
		c.UpdateFieldPlayerGuildid:                    0, // TODO
		c.UpdateFieldPlayerGuildrank:                  0, // TODO
		c.UpdateFieldPlayerBytes:                      0, // TODO
		c.UpdateFieldPlayerBytes2:                     0, // TODO
		c.UpdateFieldPlayerBytes3:                     0, // TODO
		c.UpdateFieldPlayerDuelTeam:                   0, // TODO
		c.UpdateFieldPlayerGuildTimestamp:             0, // TODO
		c.UpdateFieldPlayerQuestStart:                 0, // TODO
		c.UpdateFieldPlayerVisibleItem1Creator:        0, // TODO
		c.UpdateFieldPlayerVisibleItemEntryStart:      0, // TODO
		c.UpdateFieldPlayerVisibleItem10Ench:          0, // TODO
		c.UpdateFieldPlayerVisibleItem1Properties:     0, // TODO
		c.UpdateFieldPlayerVisibleItem1Pad:            0, // TODO
		c.UpdateFieldPlayerVisibleItemLastCreator:     0, // TODO
		c.UpdateFieldPlayerVisibleItemLast0:           0, // TODO
		c.UpdateFieldPlayerVisibleItemLastProperties:  0, // TODO
		c.UpdateFieldPlayerVisibleItemLastPad:         0, // TODO
		c.UpdateFieldPlayerInventoryStart:             0, // TODO
		c.UpdateFieldPlayerPackSlot1:                  0, // TODO
		c.UpdateFieldPlayerPackSlotLast:               0, // TODO
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
		c.UpdateFieldPlayerXp:                         0, // TODO
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
		c.UpdateFieldPlayerCoinage:                    0, // TODO
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

	return mergeUpdateFields(fields, o.BaseGameObject.Fields())

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
