package dynamic

import (
	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/channels"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/components"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Player represents an instance of an in-game player.
type Player struct {
	GameObject

	components.BasicStats
	components.Combat
	components.HealthPower
	components.Movement
	components.Player
	components.PlayerFeatures
	components.Unit

	ZoneID int
	MapID  int

	Equipment map[static.EquipmentSlot]interfaces.GUID
	Inventory map[int]interfaces.GUID
	Bags      map[int]interfaces.GUID

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

func (p *Player) StartUpdateLoop() {
	if p.UpdateChannel() != nil {
		return
	}

	p.CreateUpdateChannel()
	go func() {
		for {
			for _, msgRaw := range <-p.UpdateChannel() {
				switch msg := msgRaw.(type) {
				case *messages.UnitAttack:
					p.HandleAttack(msg)
				case *messages.UnitDeregisterAttacker:
					p.HandleDeregisterAttacker(msg)
				case *messages.ModHealth:
					p.HandleModHealth(msg)
				case *messages.ModPower:
					p.HandleModPower(msg)
				case *messages.UnitRegisterAttack:
					p.HandleRegisterAttacker(msg)
				case *messages.UnitStopAttack:
					p.HandleAttackStop(msg)
				}
			}
			channels.ObjectUpdates <- p.GUID()
		}
	}()
}

func InitializePlayer(player *Player) *Player {
	player.CurrentHealth = player.MaxHealth()
	player.CurrentPower = player.MaxPower()

	return player
}

// Object interface methods.
func (p *Player) GUID() interfaces.GUID             { return p.GameObject.GUID() }
func (p *Player) SetGUID(guid interfaces.GUID)      { p.GameObject.SetGUID(guid) }
func (p *Player) GetLocation() *interfaces.Location { return &p.MovementInfo.Location }

func (p *Player) UpdateFields() interfaces.UpdateFieldsMap {
	modelInfo := static.GetPlayerModelInfo(p.Race, p.Gender)
	resistances := p.resistances()

	fields := interfaces.UpdateFieldsMap{
		static.UpdateFieldUnitCharmLow:                                             uint32(0),
		static.UpdateFieldUnitCharmHigh:                                            uint32(0),
		static.UpdateFieldUnitSummonLow:                                            uint32(0),
		static.UpdateFieldUnitSummonHigh:                                           uint32(0),
		static.UpdateFieldUnitCharmedbyLow:                                         uint32(0),
		static.UpdateFieldUnitCharmedbyHigh:                                        uint32(0),
		static.UpdateFieldUnitSummonedbyLow:                                        uint32(0),
		static.UpdateFieldUnitSummonedbyHigh:                                       uint32(0),
		static.UpdateFieldUnitCreatedbyLow:                                         uint32(0),
		static.UpdateFieldUnitCreatedbyHigh:                                        uint32(0),
		static.UpdateFieldUnitTargetLow:                                            uint32(p.Target.Low()),
		static.UpdateFieldUnitTargetHigh:                                           uint32(p.Target.High()),
		static.UpdateFieldUnitPersuadedLow:                                         uint32(0),
		static.UpdateFieldUnitPersuadedHigh:                                        uint32(0),
		static.UpdateFieldUnitChannelObjectLow:                                     uint32(0), // TODO
		static.UpdateFieldUnitChannelObjectHigh:                                    uint32(0), // TODO
		static.UpdateFieldUnitHealth:                                               uint32(p.CurrentHealth),
		static.UpdateFieldUnitPowerStart + static.UpdateField(static.PowerMana):    uint32(p.CurrentPower),
		static.UpdateFieldUnitMaxHealth:                                            uint32(p.MaxHealth()),
		static.UpdateFieldUnitMaxPowerStart + static.UpdateField(static.PowerMana): uint32(p.MaxPower()),
		static.UpdateFieldUnitLevel:                                                uint32(p.Level),
		static.UpdateFieldUnitFactiontemplate:                                      uint32(4),
		static.UpdateFieldUnitBytes0:                                               uint32(uint32(p.Race.ID) | uint32(p.Class.ID)<<8 | uint32(p.Gender)<<16),
		static.UpdateFieldUnitFlags:                                                uint32(0),
		static.UpdateFieldUnitAura:                                                 uint32(0), // TODO
		static.UpdateFieldUnitAuraLast:                                             uint32(0), // TODO
		static.UpdateFieldUnitAuraflags:                                            uint32(0), // TODO
		static.UpdateFieldUnitAuraflags01:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuraflags02:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuraflags03:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuraflags04:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuraflags05:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuralevels:                                           uint32(0), // TODO
		static.UpdateFieldUnitAuralevelsLast:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuraapplications:                                     uint32(0), // TODO
		static.UpdateFieldUnitAuraapplicationsLast:                                 uint32(0), // TODO
		static.UpdateFieldUnitAurastate:                                            uint32(0), // TODO
		static.UpdateFieldUnitBaseattacktime:                                       uint32(1000),
		static.UpdateFieldUnitBoundingradius:                                       float32(modelInfo.BoundingRadius),
		static.UpdateFieldUnitCombatreach:                                          float32(modelInfo.CombatReach),
		static.UpdateFieldUnitDisplayid:                                            uint32(modelInfo.ID),
		static.UpdateFieldUnitNativedisplayid:                                      uint32(modelInfo.ID),
		static.UpdateFieldUnitMountdisplayid:                                       uint32(0), // TODO
		static.UpdateFieldUnitBytes1:                                               uint32(p.FreeTalentPoints)<<16 | uint32(p.StandState),
		static.UpdateFieldUnitPetnumber:                                            uint32(0), // TODO
		static.UpdateFieldUnitPetNameTimestamp:                                     uint32(0), // TODO
		static.UpdateFieldUnitPetexperience:                                        uint32(0), // TODO
		static.UpdateFieldUnitPetnextlevelexp:                                      uint32(0), // TODO
		static.UpdateFieldUnitDynamicFlags:                                         uint32(0), // TODO
		static.UpdateFieldUnitChannelSpell:                                         uint32(0), // TODO
		static.UpdateFieldUnitModCastSpeed:                                         float32(1.0),
		static.UpdateFieldUnitCreatedBySpell:                                       uint32(0), // TODO
		static.UpdateFieldUnitNpcFlags:                                             uint32(0), // TODO
		static.UpdateFieldUnitNpcEmotestate:                                        uint32(0),
		static.UpdateFieldUnitTrainingPoints:                                       uint32(0),
		static.UpdateFieldUnitStrength:                                             uint32(p.Strength),
		static.UpdateFieldUnitAgility:                                              uint32(p.Agility),
		static.UpdateFieldUnitStamina:                                              uint32(p.Stamina),
		static.UpdateFieldUnitIntellect:                                            uint32(p.Intellect),
		static.UpdateFieldUnitSpirit:                                               uint32(p.Spirit),
		static.UpdateFieldUnitArmor:                                                uint32(resistances[static.SpellSchoolPhysical]),
		static.UpdateFieldUnitHolyResist:                                           uint32(resistances[static.SpellSchoolHoly]),
		static.UpdateFieldUnitFireResist:                                           uint32(resistances[static.SpellSchoolFire]),
		static.UpdateFieldUnitNatureResist:                                         uint32(resistances[static.SpellSchoolNature]),
		static.UpdateFieldUnitFrostResist:                                          uint32(resistances[static.SpellSchoolFrost]),
		static.UpdateFieldUnitShadowResist:                                         uint32(resistances[static.SpellSchoolShadow]),
		static.UpdateFieldUnitArcaneResist:                                         uint32(resistances[static.SpellSchoolArcane]),
		static.UpdateFieldUnitBaseMana:                                             uint32(0),
		static.UpdateFieldUnitBaseHealth:                                           uint32(0),
		static.UpdateFieldUnitBytes2:                                               uint32(0), // TODO
		static.UpdateFieldUnitAttackPower:                                          uint32(p.MeleeAttackPower(p.Class)),
		static.UpdateFieldUnitAttackPowerMods:                                      uint32(0),
		static.UpdateFieldUnitAttackPowerMultiplier:                                uint32(0), // TODO
		static.UpdateFieldUnitRangedAttackPower:                                    uint32(p.RangedAttackPower(p.Class)),
		static.UpdateFieldUnitRangedAttackPowerMods:                                uint32(0),
		static.UpdateFieldUnitRangedAttackPowerMultiplier:                          uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier:                                    uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier01:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier02:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier03:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier04:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier05:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostModifier06:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier:                                  uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier01:                                uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier02:                                uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier03:                                uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier04:                                uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier05:                                uint32(0), // TODO
		static.UpdateFieldUnitPowerCostMultiplier06:                                uint32(0), // TODO

		static.UpdateFieldPlayerDuelArbiter:                uint32(0), // TODO
		static.UpdateFieldPlayerFlags:                      uint32(p.flags()),
		static.UpdateFieldPlayerGuildid:                    uint32(0), // TODO
		static.UpdateFieldPlayerGuildrank:                  uint32(0), // TODO
		static.UpdateFieldPlayerBytes:                      p.Bytes(),
		static.UpdateFieldPlayerBytes2:                     p.Bytes2(),
		static.UpdateFieldPlayerBytes3:                     uint32(p.Gender) | uint32(p.DrunkValue)&0xFFFE,
		static.UpdateFieldPlayerDuelTeam:                   uint32(0), // TODO
		static.UpdateFieldPlayerGuildTimestamp:             uint32(0), // TODO
		static.UpdateFieldPlayerQuestStart:                 uint32(0), // TODO
		static.UpdateFieldPlayerBankSlot1:                  uint32(0), // TODO
		static.UpdateFieldPlayerBankSlotLast:               uint32(0), // TODO
		static.UpdateFieldPlayerBankbagSlot1:               uint32(0), // TODO
		static.UpdateFieldPlayerBankbagSlotLast:            uint32(0), // TODO
		static.UpdateFieldPlayerVendorbuybackSlot1:         uint32(0), // TODO
		static.UpdateFieldPlayerVendorbuybackSlotLast:      uint32(0), // TODO
		static.UpdateFieldPlayerKeyringSlot1:               uint32(0), // TODO
		static.UpdateFieldPlayerKeyringSlotLast:            uint32(0), // TODO
		static.UpdateFieldPlayerFarsight:                   uint32(0), // TODO
		static.UpdateFieldPlayerComboTarget:                uint32(0), // TODO
		static.UpdateFieldPlayerXp:                         uint32(p.XP),
		static.UpdateFieldPlayerNextLevelXp:                uint32(1), // TODO
		static.UpdateFieldPlayerSkillInfo11:                uint32(0), // TODO
		static.UpdateFieldPlayerCharacterPoints1:           uint32(0), // TODO
		static.UpdateFieldPlayerCharacterPoints2:           uint32(0), // TODO
		static.UpdateFieldPlayerTrackCreatures:             uint32(0), // TODO
		static.UpdateFieldPlayerTrackResources:             uint32(0), // TODO
		static.UpdateFieldPlayerBlockPercentage:            uint32(0), // TODO
		static.UpdateFieldPlayerDodgePercentage:            uint32(0), // TODO
		static.UpdateFieldPlayerParryPercentage:            uint32(0), // TODO
		static.UpdateFieldPlayerCritPercentage:             uint32(0), // TODO
		static.UpdateFieldPlayerRangedCritPercentage:       uint32(0), // TODO
		static.UpdateFieldPlayerExploredZones1:             uint32(0), // TODO
		static.UpdateFieldPlayerRestStateExperience:        uint32(0), // TODO
		static.UpdateFieldPlayerCoinage:                    uint32(p.Money),
		static.UpdateFieldPlayerPosstat0:                   uint32(0), // TODO
		static.UpdateFieldPlayerPosstat1:                   uint32(0), // TODO
		static.UpdateFieldPlayerPosstat2:                   uint32(0), // TODO
		static.UpdateFieldPlayerPosstat3:                   uint32(0), // TODO
		static.UpdateFieldPlayerPosstat4:                   uint32(0), // TODO
		static.UpdateFieldPlayerNegstat0:                   uint32(0), // TODO
		static.UpdateFieldPlayerNegstat1:                   uint32(0), // TODO
		static.UpdateFieldPlayerNegstat2:                   uint32(0), // TODO
		static.UpdateFieldPlayerNegstat3:                   uint32(0), // TODO
		static.UpdateFieldPlayerNegstat4:                   uint32(0), // TODO
		static.UpdateFieldPlayerResistancebuffmodspositive: uint32(0), // TODO
		static.UpdateFieldPlayerResistancebuffmodsnegative: uint32(0), // TODO
		static.UpdateFieldPlayerModDamageDonePos:           uint32(0), // TODO
		static.UpdateFieldPlayerModDamageDoneNeg:           uint32(0), // TODO
		static.UpdateFieldPlayerModDamageDonePct:           float32(1.0),
		static.UpdateFieldPlayerFieldBytes:                 uint32(0), // TODO
		static.UpdateFieldPlayerAmmoID:                     uint32(0), // TODO
		static.UpdateFieldPlayerSelfResSpell:               uint32(0), // TODO
		static.UpdateFieldPlayerPvpMedals:                  uint32(0), // TODO
		static.UpdateFieldPlayerBuybackPrice1:              uint32(0), // TODO
		static.UpdateFieldPlayerBuybackPriceLast:           uint32(0), // TODO
		static.UpdateFieldPlayerBuybackTimestamp1:          uint32(0), // TODO
		static.UpdateFieldPlayerBuybackTimestampLast:       uint32(0), // TODO
		static.UpdateFieldPlayerSessionKills:               uint32(0), // TODO
		static.UpdateFieldPlayerYesterdayKills:             uint32(0), // TODO
		static.UpdateFieldPlayerLastWeekKills:              uint32(0), // TODO
		static.UpdateFieldPlayerThisWeekKills:              uint32(0), // TODO
		static.UpdateFieldPlayerThisWeekContribution:       uint32(0), // TODO
		static.UpdateFieldPlayerLifetimeHonorableKills:     uint32(0), // TODO
		static.UpdateFieldPlayerLifetimeDishonorableKills:  uint32(0), // TODO
		static.UpdateFieldPlayerYesterdayContribution:      uint32(0), // TODO
		static.UpdateFieldPlayerLastWeekContribution:       uint32(0), // TODO
		static.UpdateFieldPlayerLastWeekRank:               uint32(0), // TODO
		static.UpdateFieldPlayerBytes2b:                    uint32(0), // TODO
		static.UpdateFieldPlayerWatchedFactionIndex:        uint32(0), // TODO
		static.UpdateFieldPlayerCombatRating1:              uint32(0), // TODO
	}

	for slot, itemGUID := range p.Equipment {
		if !GetObjectManager().Exists(itemGUID) {
			GetObjectManager().log.WithFields(logrus.Fields{
				"player":    p.GUID(),
				"slot":      slot,
				"item_guid": itemGUID,
			}).Errorf("Unknown equipped item")
			continue
		}

		item := GetObjectManager().GetItem(itemGUID)

		slotField := static.UpdateFieldPlayerInventoryStart + static.UpdateField(slot*2)
		fields[slotField] = uint32(item.GUID().Low())
		fields[slotField+1] = uint32(item.GUID().High())

		visibleItemSlot := static.UpdateField(slot * 12)
		fields[static.UpdateFieldPlayerVisibleItemEntryStart+visibleItemSlot] = uint32(item.GetTemplate().Entry)

		if GetObjectManager().Exists(item.Creator) {
			fields[static.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot] = uint32(item.Creator.Low())
			fields[static.UpdateFieldPlayerVisibleItem1Creator+visibleItemSlot+1] = uint32(item.Creator.High())
		}

		if slot == static.EquipmentSlotMainHand {
			fields[static.UpdateFieldUnitBaseattacktime] = uint32(item.GetTemplate().AttackRate.Milliseconds())
			fields[static.UpdateFieldUnitMindamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Min)
			fields[static.UpdateFieldUnitMaxdamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Max)
		} else if slot == static.EquipmentSlotOffHand {
			fields[static.UpdateFieldUnitOffhandattacktime] = uint32(item.GetTemplate().AttackRate.Milliseconds())
			fields[static.UpdateFieldUnitMinoffhanddamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Min)
			fields[static.UpdateFieldUnitMaxoffhanddamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Max)
		} else if slot == static.EquipmentSlotRanged {
			fields[static.UpdateFieldUnitRangedattacktime] = uint32(item.GetTemplate().AttackRate.Milliseconds())
			fields[static.UpdateFieldUnitMinrangeddamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Min)
			fields[static.UpdateFieldUnitMaxrangeddamage] = float32(item.GetTemplate().Damages[static.SpellSchoolPhysical].Max)
		}
	}

	for slot, bag := range p.Bags {
		slotField := static.UpdateFieldPlayerBagStart + static.UpdateField(slot*2)
		fields[slotField] = uint32(bag.Low())
		fields[slotField+1] = uint32(bag.High())
	}

	for slot, item := range p.Inventory {
		slotField := static.UpdateFieldPlayerPackSlot1 + static.UpdateField(slot*2)
		fields[slotField] = uint32(item.Low())
		fields[slotField+1] = uint32(item.High())
	}

	mergedFields := p.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[static.UpdateFieldType] = uint32(TypeMask(p))
	delete(mergedFields, static.UpdateFieldEntry)

	return mergedFields
}

// Utility methods.
func (p *Player) flags() uint32 {
	var flags uint32
	if p.IsGroupLeader {
		flags |= uint32(static.PlayerFlagsGroupLeader)
	}
	if p.IsAFK {
		flags |= uint32(static.PlayerFlagsAFK)
	}
	if p.IsDND {
		flags |= uint32(static.PlayerFlagsDND)
	}
	if p.IsGM {
		flags |= uint32(static.PlayerFlagsGM)
	}
	if p.IsGhost {
		flags |= uint32(static.PlayerFlagsGhost)
	}
	if p.IsResting {
		flags |= uint32(static.PlayerFlagsResting)
	}
	if p.IsFFAPVP {
		flags |= uint32(static.PlayerFlagsFFAPVP)
	}
	if p.IsContestedPVP {
		flags |= uint32(static.PlayerFlagsContestedPVP)
	}
	if p.IsInPVP {
		flags |= uint32(static.PlayerFlagsInPVP)
	}
	if p.HideHelm {
		flags |= uint32(static.PlayerFlagsHideHelm)
	}
	if p.HideCloak {
		flags |= uint32(static.PlayerFlagsHideCloak)
	}
	if p.IsPartialPlayTime {
		flags |= uint32(static.PlayerFlagsPartialPlayTime)
	}
	if p.IsNoPlayTime {
		flags |= uint32(static.PlayerFlagsNoPlayTime)
	}
	if p.IsInSanctuary {
		flags |= uint32(static.PlayerFlagsSanctuary)
	}
	if p.IsTaxiBenchmark {
		flags |= uint32(static.PlayerFlagsTaxiBenchmark)
	}
	if p.IsPVPTimer {
		flags |= uint32(static.PlayerFlagsPVPTimer)
	}

	return flags
}

// FirstBag returns the first bag the player has equipped, or nil if there are no bags.
func (p *Player) FirstBag() *Container {
	for i := 0; i < 4; i++ {
		bagGUID, ok := p.Bags[i]
		if ok {
			return GetObjectManager().GetContainer(bagGUID)
		}
	}

	return nil
}
