package object

import (
	"encoding/json"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/sirupsen/logrus"
)

// Player represents an instance of an in-game player.
type Player struct {
	Unit

	SkinColor uint8
	Face      uint8
	HairStyle uint8
	HairColor uint8
	Feature   uint8

	ZoneID int
	MapID  int

	Equipment map[c.EquipmentSlot]GUID
	Inventory map[int]GUID
	Bags      map[int]GUID

	DrunkValue int
	XP         int
	Money      int

	Tutorials [256]bool

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
	IsLoggedIn        bool

	IsTrackStealthed           bool
	ShowAutoReleaseSpiritTimer bool
	HideReleaseSpirit          bool
}

// Manager returns the manager associated with this object.
func (p *Player) Manager() *Manager { return p.GameObject.Manager() }

// SetManager updates the manager associated with this object.
func (p *Player) SetManager(manager *Manager) { p.GameObject.SetManager(manager) }

// GUID returns the globally-unique ID of the object.
func (p *Player) GUID() GUID { return p.GameObject.GUID() }

// SetGUID updates this object's GUID to the given value.
func (p *Player) SetGUID(guid GUID) { p.GameObject.SetGUID(guid) }

// Location returns the location of the object.
func (p *Player) Location() *Location { return p.Unit.Location() }

func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Player) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, p)
}

// MovementUpdate calculates and returns the movement update for the
// object.
func (p *Player) MovementUpdate() []byte { return p.Unit.MovementUpdate() }

// UpdateFields populates and returns the updated fields for the
// object.
func (p *Player) UpdateFields() UpdateFieldsMap {
	modelInfo := dbc.GetPlayerModelInfo(p.Race, p.Gender)
	resistances := p.Resistances()

	fields := UpdateFieldsMap{
		c.UpdateFieldUnitCharmLow:                                     uint32(p.Charm.Low()),
		c.UpdateFieldUnitCharmHigh:                                    uint32(p.Charm.High()),
		c.UpdateFieldUnitSummonLow:                                    uint32(p.Summon.Low()),
		c.UpdateFieldUnitSummonHigh:                                   uint32(p.Summon.High()),
		c.UpdateFieldUnitCharmedbyLow:                                 uint32(p.CharmedBy.Low()),
		c.UpdateFieldUnitCharmedbyHigh:                                uint32(p.CharmedBy.High()),
		c.UpdateFieldUnitSummonedbyLow:                                uint32(p.SummonedBy.Low()),
		c.UpdateFieldUnitSummonedbyHigh:                               uint32(p.SummonedBy.High()),
		c.UpdateFieldUnitCreatedbyLow:                                 uint32(p.CreatedBy.Low()),
		c.UpdateFieldUnitCreatedbyHigh:                                uint32(p.CreatedBy.High()),
		c.UpdateFieldUnitTargetLow:                                    uint32(p.Target.Low()),
		c.UpdateFieldUnitTargetHigh:                                   uint32(p.Target.High()),
		c.UpdateFieldUnitPersuadedLow:                                 uint32(p.Persuaded.Low()),
		c.UpdateFieldUnitPersuadedHigh:                                uint32(p.Persuaded.High()),
		c.UpdateFieldUnitChannelObjectLow:                             uint32(0), // TODO
		c.UpdateFieldUnitChannelObjectHigh:                            uint32(0), // TODO
		c.UpdateFieldUnitHealth:                                       uint32(float32(p.maxHealth()) * p.HealthPercent),
		c.UpdateFieldUnitPowerStart + c.UpdateField(p.powerType()):    uint32(float32(p.maxPower()) * 0.5),
		c.UpdateFieldUnitMaxHealth:                                    uint32(p.maxHealth()),
		c.UpdateFieldUnitMaxPowerStart + c.UpdateField(p.powerType()): uint32(p.maxPower()),
		c.UpdateFieldUnitLevel:                                        uint32(p.Level),
		c.UpdateFieldUnitFactiontemplate:                              uint32(4),
		c.UpdateFieldUnitBytes0:                                       uint32(uint32(p.Race.ID) | uint32(p.Class.ID)<<8 | uint32(p.Gender)<<16),
		c.UpdateFieldUnitFlags:                                        uint32(0),
		c.UpdateFieldUnitAura:                                         uint32(0), // TODO
		c.UpdateFieldUnitAuraLast:                                     uint32(0), // TODO
		c.UpdateFieldUnitAuraflags:                                    uint32(0), // TODO
		c.UpdateFieldUnitAuraflags01:                                  uint32(0), // TODO
		c.UpdateFieldUnitAuraflags02:                                  uint32(0), // TODO
		c.UpdateFieldUnitAuraflags03:                                  uint32(0), // TODO
		c.UpdateFieldUnitAuraflags04:                                  uint32(0), // TODO
		c.UpdateFieldUnitAuraflags05:                                  uint32(0), // TODO
		c.UpdateFieldUnitAuralevels:                                   uint32(0), // TODO
		c.UpdateFieldUnitAuralevelsLast:                               uint32(0), // TODO
		c.UpdateFieldUnitAuraapplications:                             uint32(0), // TODO
		c.UpdateFieldUnitAuraapplicationsLast:                         uint32(0), // TODO
		c.UpdateFieldUnitAurastate:                                    uint32(0), // TODO
		c.UpdateFieldUnitBaseattacktime:                               uint32(1000),
		c.UpdateFieldUnitBoundingradius:                               float32(modelInfo.BoundingRadius),
		c.UpdateFieldUnitCombatreach:                                  float32(modelInfo.CombatReach),
		c.UpdateFieldUnitDisplayid:                                    uint32(modelInfo.ID),
		c.UpdateFieldUnitNativedisplayid:                              uint32(modelInfo.ID),
		c.UpdateFieldUnitMountdisplayid:                               uint32(0), // TODO
		c.UpdateFieldUnitBytes1:                                       uint32(p.Byte1Flags)<<24 | uint32(p.FreeTalentPoints)<<16 | uint32(p.StandState),
		c.UpdateFieldUnitPetnumber:                                    uint32(0), // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             uint32(0), // TODO
		c.UpdateFieldUnitPetexperience:                                uint32(0), // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              uint32(0), // TODO
		c.UpdateFieldUnitDynamicFlags:                                 uint32(0), // TODO
		c.UpdateFieldUnitChannelSpell:                                 uint32(0), // TODO
		c.UpdateFieldUnitModCastSpeed:                                 float32(1.0),
		c.UpdateFieldUnitCreatedBySpell:                               uint32(0), // TODO
		c.UpdateFieldUnitNpcFlags:                                     uint32(0), // TODO
		c.UpdateFieldUnitNpcEmotestate:                                uint32(p.EmoteState),
		c.UpdateFieldUnitTrainingPoints:                               uint32(p.TrainingPoints),
		c.UpdateFieldUnitStrength:                                     uint32(p.Strength),
		c.UpdateFieldUnitAgility:                                      uint32(p.Agility),
		c.UpdateFieldUnitStamina:                                      uint32(p.Stamina),
		c.UpdateFieldUnitIntellect:                                    uint32(p.Intellect),
		c.UpdateFieldUnitSpirit:                                       uint32(p.Spirit),
		c.UpdateFieldUnitArmor:                                        uint32(resistances[c.SpellSchoolPhysical]),
		c.UpdateFieldUnitHolyResist:                                   uint32(resistances[c.SpellSchoolHoly]),
		c.UpdateFieldUnitFireResist:                                   uint32(resistances[c.SpellSchoolFire]),
		c.UpdateFieldUnitNatureResist:                                 uint32(resistances[c.SpellSchoolNature]),
		c.UpdateFieldUnitFrostResist:                                  uint32(resistances[c.SpellSchoolFrost]),
		c.UpdateFieldUnitShadowResist:                                 uint32(resistances[c.SpellSchoolShadow]),
		c.UpdateFieldUnitArcaneResist:                                 uint32(resistances[c.SpellSchoolArcane]),
		c.UpdateFieldUnitBaseMana:                                     uint32(p.BasePower),
		c.UpdateFieldUnitBaseHealth:                                   uint32(p.BaseHealth),
		c.UpdateFieldUnitBytes2:                                       uint32(0), // TODO
		c.UpdateFieldUnitAttackPower:                                  uint32(p.meleeAttackPower()),
		c.UpdateFieldUnitAttackPowerMods:                              uint32(p.meleeAttackPowerMods()),
		c.UpdateFieldUnitAttackPowerMultiplier:                        uint32(0), // TODO
		c.UpdateFieldUnitRangedAttackPower:                            uint32(p.rangedAttackPower()),
		c.UpdateFieldUnitRangedAttackPowerMods:                        uint32(p.rangedAttackPowerMods()),
		c.UpdateFieldUnitRangedAttackPowerMultiplier:                  uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier:                            uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier01:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier02:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier03:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier04:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier05:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostModifier06:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier:                          uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier01:                        uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier02:                        uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier03:                        uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier04:                        uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier05:                        uint32(0), // TODO
		c.UpdateFieldUnitPowerCostMultiplier06:                        uint32(0), // TODO

		c.UpdateFieldPlayerDuelArbiter:                uint32(0), // TODO
		c.UpdateFieldPlayerFlags:                      uint32(p.flags()),
		c.UpdateFieldPlayerGuildid:                    uint32(0), // TODO
		c.UpdateFieldPlayerGuildrank:                  uint32(0), // TODO
		c.UpdateFieldPlayerBytes:                      uint32(p.SkinColor) | uint32(p.Face)<<8 | uint32(p.HairStyle)<<16 | uint32(p.HairColor)<<24,
		c.UpdateFieldPlayerBytes2:                     uint32(p.Feature),
		c.UpdateFieldPlayerBytes3:                     uint32(p.Gender) | uint32(p.DrunkValue)&0xFFFE,
		c.UpdateFieldPlayerDuelTeam:                   uint32(0), // TODO
		c.UpdateFieldPlayerGuildTimestamp:             uint32(0), // TODO
		c.UpdateFieldPlayerQuestStart:                 uint32(0), // TODO
		c.UpdateFieldPlayerBankSlot1:                  uint32(0), // TODO
		c.UpdateFieldPlayerBankSlotLast:               uint32(0), // TODO
		c.UpdateFieldPlayerBankbagSlot1:               uint32(0), // TODO
		c.UpdateFieldPlayerBankbagSlotLast:            uint32(0), // TODO
		c.UpdateFieldPlayerVendorbuybackSlot1:         uint32(0), // TODO
		c.UpdateFieldPlayerVendorbuybackSlotLast:      uint32(0), // TODO
		c.UpdateFieldPlayerKeyringSlot1:               uint32(0), // TODO
		c.UpdateFieldPlayerKeyringSlotLast:            uint32(0), // TODO
		c.UpdateFieldPlayerFarsight:                   uint32(0), // TODO
		c.UpdateFieldPlayerComboTarget:                uint32(0), // TODO
		c.UpdateFieldPlayerXp:                         uint32(p.XP),
		c.UpdateFieldPlayerNextLevelXp:                uint32(1), // TODO
		c.UpdateFieldPlayerSkillInfo11:                uint32(0), // TODO
		c.UpdateFieldPlayerCharacterPoints1:           uint32(0), // TODO
		c.UpdateFieldPlayerCharacterPoints2:           uint32(0), // TODO
		c.UpdateFieldPlayerTrackCreatures:             uint32(0), // TODO
		c.UpdateFieldPlayerTrackResources:             uint32(0), // TODO
		c.UpdateFieldPlayerBlockPercentage:            uint32(0), // TODO
		c.UpdateFieldPlayerDodgePercentage:            uint32(0), // TODO
		c.UpdateFieldPlayerParryPercentage:            uint32(0), // TODO
		c.UpdateFieldPlayerCritPercentage:             uint32(0), // TODO
		c.UpdateFieldPlayerRangedCritPercentage:       uint32(0), // TODO
		c.UpdateFieldPlayerExploredZones1:             uint32(0), // TODO
		c.UpdateFieldPlayerRestStateExperience:        uint32(0), // TODO
		c.UpdateFieldPlayerCoinage:                    uint32(p.Money),
		c.UpdateFieldPlayerPosstat0:                   uint32(0), // TODO
		c.UpdateFieldPlayerPosstat1:                   uint32(0), // TODO
		c.UpdateFieldPlayerPosstat2:                   uint32(0), // TODO
		c.UpdateFieldPlayerPosstat3:                   uint32(0), // TODO
		c.UpdateFieldPlayerPosstat4:                   uint32(0), // TODO
		c.UpdateFieldPlayerNegstat0:                   uint32(0), // TODO
		c.UpdateFieldPlayerNegstat1:                   uint32(0), // TODO
		c.UpdateFieldPlayerNegstat2:                   uint32(0), // TODO
		c.UpdateFieldPlayerNegstat3:                   uint32(0), // TODO
		c.UpdateFieldPlayerNegstat4:                   uint32(0), // TODO
		c.UpdateFieldPlayerResistancebuffmodspositive: uint32(0), // TODO
		c.UpdateFieldPlayerResistancebuffmodsnegative: uint32(0), // TODO
		c.UpdateFieldPlayerModDamageDonePos:           uint32(0), // TODO
		c.UpdateFieldPlayerModDamageDoneNeg:           uint32(0), // TODO
		c.UpdateFieldPlayerModDamageDonePct:           float32(p.damageModPercentage()),
		c.UpdateFieldPlayerFieldBytes:                 uint32(0), // TODO
		c.UpdateFieldPlayerAmmoID:                     uint32(0), // TODO
		c.UpdateFieldPlayerSelfResSpell:               uint32(0), // TODO
		c.UpdateFieldPlayerPvpMedals:                  uint32(0), // TODO
		c.UpdateFieldPlayerBuybackPrice1:              uint32(0), // TODO
		c.UpdateFieldPlayerBuybackPriceLast:           uint32(0), // TODO
		c.UpdateFieldPlayerBuybackTimestamp1:          uint32(0), // TODO
		c.UpdateFieldPlayerBuybackTimestampLast:       uint32(0), // TODO
		c.UpdateFieldPlayerSessionKills:               uint32(0), // TODO
		c.UpdateFieldPlayerYesterdayKills:             uint32(0), // TODO
		c.UpdateFieldPlayerLastWeekKills:              uint32(0), // TODO
		c.UpdateFieldPlayerThisWeekKills:              uint32(0), // TODO
		c.UpdateFieldPlayerThisWeekContribution:       uint32(0), // TODO
		c.UpdateFieldPlayerLifetimeHonorableKills:     uint32(0), // TODO
		c.UpdateFieldPlayerLifetimeDishonorableKills:  uint32(0), // TODO
		c.UpdateFieldPlayerYesterdayContribution:      uint32(0), // TODO
		c.UpdateFieldPlayerLastWeekContribution:       uint32(0), // TODO
		c.UpdateFieldPlayerLastWeekRank:               uint32(0), // TODO
		c.UpdateFieldPlayerBytes2b:                    uint32(0), // TODO
		c.UpdateFieldPlayerWatchedFactionIndex:        uint32(0), // TODO
		c.UpdateFieldPlayerCombatRating1:              uint32(0), // TODO
	}

	for slot, itemGUID := range p.Equipment {
		if !p.Manager().Exists(itemGUID) {
			p.Manager().log.WithFields(logrus.Fields{
				"player":    p.GUID(),
				"slot":      slot.String(),
				"item_guid": itemGUID,
			}).Errorf("Unknown equipped item")
			continue
		}

		item := p.Manager().Get(itemGUID).(*Item)

		slotField := c.UpdateFieldPlayerInventoryStart + c.UpdateField(slot*2)
		fields[slotField] = uint32(item.GUID().Low())
		fields[slotField+1] = uint32(item.GUID().High())

		visibleItemSlot := c.UpdateField(slot * 12)
		fields[c.UpdateFieldPlayerVisibleItemEntryStart+visibleItemSlot] = uint32(item.Template().Entry)

		if p.Manager().Exists(item.Creator) {
			fields[c.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot] = uint32(item.Creator.Low())
			fields[c.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot+1] = uint32(item.Creator.High())
		}

		if slot == c.EquipmentSlotMainHand {
			fields[c.UpdateFieldUnitBaseattacktime] = uint32(item.Template().AttackRate.Milliseconds())
			fields[c.UpdateFieldUnitMindamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxdamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Max)
		} else if slot == c.EquipmentSlotOffHand {
			fields[c.UpdateFieldUnitOffhandattacktime] = uint32(item.Template().AttackRate.Milliseconds())
			fields[c.UpdateFieldUnitMinoffhanddamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxoffhanddamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Max)
		} else if slot == c.EquipmentSlotRanged {
			fields[c.UpdateFieldUnitRangedattacktime] = uint32(item.Template().AttackRate.Milliseconds())
			fields[c.UpdateFieldUnitMinrangeddamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Min)
			fields[c.UpdateFieldUnitMaxrangeddamage] = float32(item.Template().Damages[c.SpellSchoolPhysical].Max)
		}
	}

	for slot, bag := range p.Bags {
		slotField := c.UpdateFieldPlayerBagStart + c.UpdateField(slot*2)
		fields[slotField] = uint32(bag.Low())
		fields[slotField+1] = uint32(bag.High())
	}

	for slot, item := range p.Inventory {
		slotField := c.UpdateFieldPlayerPackSlot1 + c.UpdateField(slot*2)
		fields[slotField] = uint32(item.Low())
		fields[slotField+1] = uint32(item.High())
	}

	mergedFields := p.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[c.UpdateFieldType] = uint32(TypeMask(p))
	delete(mergedFields, c.UpdateFieldEntry)

	return mergedFields
}

func (p *Player) flags() uint32 {
	var flags uint32
	if p.IsGroupLeader {
		flags |= uint32(c.PlayerFlagsGroupLeader)
	}
	if p.IsAFK {
		flags |= uint32(c.PlayerFlagsAFK)
	}
	if p.IsDND {
		flags |= uint32(c.PlayerFlagsDND)
	}
	if p.IsGM {
		flags |= uint32(c.PlayerFlagsGM)
	}
	if p.IsGhost {
		flags |= uint32(c.PlayerFlagsGhost)
	}
	if p.IsResting {
		flags |= uint32(c.PlayerFlagsResting)
	}
	if p.IsFFAPVP {
		flags |= uint32(c.PlayerFlagsFFAPVP)
	}
	if p.IsContestedPVP {
		flags |= uint32(c.PlayerFlagsContestedPVP)
	}
	if p.IsInPVP {
		flags |= uint32(c.PlayerFlagsInPVP)
	}
	if p.HideHelm {
		flags |= uint32(c.PlayerFlagsHideHelm)
	}
	if p.HideCloak {
		flags |= uint32(c.PlayerFlagsHideCloak)
	}
	if p.IsPartialPlayTime {
		flags |= uint32(c.PlayerFlagsPartialPlayTime)
	}
	if p.IsNoPlayTime {
		flags |= uint32(c.PlayerFlagsNoPlayTime)
	}
	if p.IsInSanctuary {
		flags |= uint32(c.PlayerFlagsSanctuary)
	}
	if p.IsTaxiBenchmark {
		flags |= uint32(c.PlayerFlagsTaxiBenchmark)
	}
	if p.IsPVPTimer {
		flags |= uint32(c.PlayerFlagsPVPTimer)
	}

	return flags
}

// FirstBag returns the first bag the player has equipped, or nil if there are no bags.
func (p *Player) FirstBag() *Container {
	for i := 0; i < 4; i++ {
		bagGUID, ok := p.Bags[i]
		if ok {
			return p.Manager().Get(bagGUID).(*Container)
		}
	}

	return nil
}
