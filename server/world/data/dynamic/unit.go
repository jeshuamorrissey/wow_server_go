package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/channels"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/components"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Unit represents an instance of an in-game monster.
type Unit struct {
	GameObject

	// Basic information.
	components.Unit
	components.Movement
	components.HealthPower
	components.Combat

	// Various timers.
	RespawnTimeMS time.Duration // Time until the unit respawns. time.Duration(0) if it doesn't respawn.
}

func (u *Unit) StartUpdateLoop() {
	if u.UpdateChannel() != nil {
		return
	}

	u.CreateUpdateChannel()
	go func() {
		for {
			for _, update := range <-u.UpdateChannel() {
				switch updateTyped := update.(type) {
				case *messages.UnitModHealth:
					u.ModHealth(updateTyped.Amount, u.maxHealth())
					if u.CurrentHealth == 0 {
						u.SendUpdates([]interface{}{
							&messages.UnitStopAttack{},
						})

						for attacker := range u.Attackers {
							GetObjectManager().Get(attacker).SendUpdates([]interface{}{
								&messages.UnitDeregisterAttacker{Attacker: u.GUID()},
								&messages.UnitDied{DeadUnit: u.GUID()},
							})
						}
					}
				case *messages.UnitModPower:
					u.ModPower(updateTyped.Amount, u.maxPower())
				case *messages.UnitRegisterAttack:
					if !u.IsInCombat() {
						u.HandleAttack(updateTyped.Attacker)
					}
					u.RegisterAttacker(updateTyped.Attacker)
				case *messages.UnitDeregisterAttacker:
					u.DeregisterAttacker(updateTyped.Attacker)
				case *messages.UnitAttack:
					target := GetObjectManager().Get(updateTyped.Target)
					target.SendUpdates([]interface{}{
						&messages.UnitRegisterAttack{Attacker: u.GUID()},
					})

					u.Attack(u, target, 1600*time.Millisecond, func() *components.Damage {
						return &components.Damage{
							Base: map[static.SpellSchool]int{
								static.SpellSchoolPhysical: 5,
							},
						}
					})

				case *messages.UnitStopAttack:
					u.StopAttack()
				}
			}

			channels.ObjectUpdates <- u.GUID()
		}
	}()
}

func InitializeUnit(unit *Unit) *Unit {
	unit.CurrentHealth = unit.maxHealth()
	unit.CurrentPower = unit.maxPower()

	return unit
}

// Object interface methods.
func (u *Unit) GUID() interfaces.GUID             { return u.GameObject.GUID() }
func (u *Unit) SetGUID(guid interfaces.GUID)      { u.GameObject.SetGUID(guid) }
func (u *Unit) GetLocation() *interfaces.Location { return &u.MovementInfo.Location }

func (u *Unit) UpdateFields() interfaces.UpdateFieldsMap {
	tmpl := u.Template()
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
		static.UpdateFieldUnitTargetLow:                                            uint32(u.Target.Low()),
		static.UpdateFieldUnitTargetHigh:                                           uint32(u.Target.High()),
		static.UpdateFieldUnitPersuadedLow:                                         uint32(0),
		static.UpdateFieldUnitPersuadedHigh:                                        uint32(0),
		static.UpdateFieldUnitChannelObjectLow:                                     uint32(0), // TODO
		static.UpdateFieldUnitChannelObjectHigh:                                    uint32(0), // TODO
		static.UpdateFieldUnitHealth:                                               uint32(u.CurrentHealth),
		static.UpdateFieldUnitPowerStart + static.UpdateField(static.PowerMana):    uint32(u.CurrentPower),
		static.UpdateFieldUnitMaxHealth:                                            uint32(u.maxHealth()),
		static.UpdateFieldUnitMaxPowerStart + static.UpdateField(static.PowerMana): uint32(tmpl.MaxPower),
		static.UpdateFieldUnitLevel:                                                uint32(u.Level),
		static.UpdateFieldUnitBytes0:                                               uint32(u.Race.ID) | uint32(u.Class.ID)<<8 | uint32(u.Gender)<<16,
		static.UpdateFieldUnitFlags:                                                uint32(tmpl.Flags()),
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
		static.UpdateFieldUnitBaseattacktime:                                       uint32(0), // TODO
		static.UpdateFieldUnitOffhandattacktime:                                    uint32(0), // TODO
		static.UpdateFieldUnitRangedattacktime:                                     uint32(0), // TODO
		static.UpdateFieldUnitBoundingradius:                                       uint32(tmpl.BoundingRadius),
		static.UpdateFieldUnitCombatreach:                                          uint32(tmpl.CombatReach),
		static.UpdateFieldUnitDisplayid:                                            uint32(tmpl.DisplayID),
		static.UpdateFieldUnitNativedisplayid:                                      uint32(0), // TODO
		static.UpdateFieldUnitMountdisplayid:                                       uint32(0), // TODO
		static.UpdateFieldUnitMindamage:                                            uint32(0), // TODO
		static.UpdateFieldUnitMaxdamage:                                            uint32(0), // TODO
		static.UpdateFieldUnitMinoffhanddamage:                                     uint32(0), // TODO
		static.UpdateFieldUnitMaxoffhanddamage:                                     uint32(0), // TODO
		static.UpdateFieldUnitBytes1:                                               uint32(0), // TODO
		static.UpdateFieldUnitPetnumber:                                            uint32(0), // TODO
		static.UpdateFieldUnitPetNameTimestamp:                                     uint32(0), // TODO
		static.UpdateFieldUnitPetexperience:                                        uint32(0), // TODO
		static.UpdateFieldUnitPetnextlevelexp:                                      uint32(0), // TODO
		static.UpdateFieldUnitDynamicFlags:                                         uint32(0), // TODO
		static.UpdateFieldUnitChannelSpell:                                         uint32(0), // TODO
		static.UpdateFieldUnitModCastSpeed:                                         uint32(0), // TODO
		static.UpdateFieldUnitCreatedBySpell:                                       uint32(0), // TODO
		static.UpdateFieldUnitNpcFlags:                                             uint32(tmpl.Flags()),
		static.UpdateFieldUnitNpcEmotestate:                                        uint32(0),  // TODO
		static.UpdateFieldUnitTrainingPoints:                                       uint32(0),  // TODO
		static.UpdateFieldUnitStrength:                                             uint32(0),  // TODO
		static.UpdateFieldUnitAgility:                                              uint32(0),  // TODO
		static.UpdateFieldUnitStamina:                                              uint32(0),  // TODO
		static.UpdateFieldUnitIntellect:                                            uint32(0),  // TODO
		static.UpdateFieldUnitSpirit:                                               uint32(0),  // TODO
		static.UpdateFieldUnitArmor:                                                uint32(0),  // TODO
		static.UpdateFieldUnitHolyResist:                                           uint32(0),  // TODO
		static.UpdateFieldUnitFireResist:                                           uint32(0),  // TODO
		static.UpdateFieldUnitNatureResist:                                         uint32(0),  // TODO
		static.UpdateFieldUnitFrostResist:                                          uint32(0),  // TODO
		static.UpdateFieldUnitShadowResist:                                         uint32(0),  // TODO
		static.UpdateFieldUnitArcaneResist:                                         uint32(0),  // TODO
		static.UpdateFieldUnitBaseMana:                                             uint32(0),  // TODO
		static.UpdateFieldUnitBaseHealth:                                           uint32(0),  // TODO
		static.UpdateFieldUnitBytes2:                                               uint32(0),  // TODO
		static.UpdateFieldUnitAttackPower:                                          uint32(0),  // TODO
		static.UpdateFieldUnitAttackPowerMods:                                      uint32(0),  // TODO
		static.UpdateFieldUnitAttackPowerMultiplier:                                uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPower:                                    uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPowerMods:                                uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPowerMultiplier:                          uint32(0),  // TODO
		static.UpdateFieldUnitMinrangeddamage:                                      uint32(0),  // TODO
		static.UpdateFieldUnitMaxrangeddamage:                                      uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier:                                    uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier01:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier02:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier03:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier04:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier05:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier06:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier:                                  uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier01:                                uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier02:                                uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier03:                                uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier04:                                uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier05:                                uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier06:                                uint32(0),  // TODO
		static.UpdateFieldUnitFactiontemplate:                                      uint32(14), // TODO
	}

	mergedFields := u.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[static.UpdateFieldType] = uint32(TypeMask(u))

	return mergedFields
}

/// Utility methods.
// Template returns the unit template this object is based on.
func (u *Unit) Template() *static.Unit {
	return static.Units[int(u.Entry)]
}
