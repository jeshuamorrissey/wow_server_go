package dynamic

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
)

// Unit represents an instance of an in-game monster.
type Unit struct {
	GameObject

	// Basic information.
	Level            int
	Race             *static.Race
	Class            *static.Class
	Gender           static.Gender
	Team             static.Team
	StandState       static.StandState
	FreeTalentPoints int
	Byte1Flags       static.Byte1Flags
	EmoteState       int
	TrainingPoints   int

	// Movement related information.
	MovementInfo      interfaces.MovementInfo
	SpeedWalk         float32
	SpeedRun          float32
	SpeedRunBackward  float32
	SpeedSwim         float32
	SpeedSwimBackward float32
	SpeedTurn         float32

	// Stats.
	BaseHealth    int
	HealthPercent float32
	BasePower     int
	PowerPercent  float32

	Strength  int
	Agility   int
	Stamina   int
	Intellect int
	Spirit    int

	// Display items (virtual). List of entries.
	VirtualItems []int

	// Flags
	CanDetectAmore0     bool
	CanDetectAmore1     bool
	CanDetectAmore2     bool
	CanDetectAmore3     bool
	IsStealth           bool
	HasInvisibilityGlow bool

	IsSwimming bool
	IsFalling  bool

	Transport interfaces.GUID

	// Relationships.
	Charm      interfaces.GUID
	CharmedBy  interfaces.GUID
	Summon     interfaces.GUID
	SummonedBy interfaces.GUID
	CreatedBy  interfaces.GUID
	Target     interfaces.GUID
	Persuaded  interfaces.GUID
}

// Object interface methods.
func (u *Unit) GUID() interfaces.GUID             { return u.GameObject.GUID() }
func (u *Unit) SetGUID(guid interfaces.GUID)      { u.GameObject.SetGUID(guid) }
func (u *Unit) GetLocation() *interfaces.Location { return &u.MovementInfo.Location }

func (u *Unit) UpdateFields() interfaces.UpdateFieldsMap {
	tmpl := u.Template()
	fields := interfaces.UpdateFieldsMap{
		static.UpdateFieldUnitCharmLow:                                          uint32(u.Charm.Low()),
		static.UpdateFieldUnitCharmHigh:                                         uint32(u.Charm.High()),
		static.UpdateFieldUnitSummonLow:                                         uint32(u.Summon.Low()),
		static.UpdateFieldUnitSummonHigh:                                        uint32(u.Summon.High()),
		static.UpdateFieldUnitCharmedbyLow:                                      uint32(u.CharmedBy.Low()),
		static.UpdateFieldUnitCharmedbyHigh:                                     uint32(u.CharmedBy.High()),
		static.UpdateFieldUnitSummonedbyLow:                                     uint32(u.SummonedBy.Low()),
		static.UpdateFieldUnitSummonedbyHigh:                                    uint32(u.SummonedBy.High()),
		static.UpdateFieldUnitCreatedbyLow:                                      uint32(u.CreatedBy.Low()),
		static.UpdateFieldUnitCreatedbyHigh:                                     uint32(u.CreatedBy.High()),
		static.UpdateFieldUnitTargetLow:                                         uint32(u.Target.Low()),
		static.UpdateFieldUnitTargetHigh:                                        uint32(u.Target.High()),
		static.UpdateFieldUnitPersuadedLow:                                      uint32(u.Persuaded.Low()),
		static.UpdateFieldUnitPersuadedHigh:                                     uint32(u.Persuaded.High()),
		static.UpdateFieldUnitChannelObjectLow:                                  uint32(0), // TODO
		static.UpdateFieldUnitChannelObjectHigh:                                 uint32(0), // TODO
		static.UpdateFieldUnitHealth:                                            uint32(float32(tmpl.MaxHealth) * u.HealthPercent),
		static.UpdateFieldUnitPowerStart + static.UpdateField(u.powerType()):    uint32(float32(tmpl.MaxPower) * u.PowerPercent),
		static.UpdateFieldUnitMaxHealth:                                         uint32(tmpl.MaxHealth),
		static.UpdateFieldUnitMaxPowerStart + static.UpdateField(u.powerType()): uint32(tmpl.MaxPower),
		static.UpdateFieldUnitLevel:                                             uint32(u.Level),
		static.UpdateFieldUnitBytes0:                                            uint32(u.Race.ID) | uint32(u.Class.ID)<<8 | uint32(u.Gender)<<16,
		static.UpdateFieldUnitFlags:                                             uint32(tmpl.Flags()),
		static.UpdateFieldUnitAura:                                              uint32(0), // TODO
		static.UpdateFieldUnitAuraLast:                                          uint32(0), // TODO
		static.UpdateFieldUnitAuraflags:                                         uint32(0), // TODO
		static.UpdateFieldUnitAuraflags01:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuraflags02:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuraflags03:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuraflags04:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuraflags05:                                       uint32(0), // TODO
		static.UpdateFieldUnitAuralevels:                                        uint32(0), // TODO
		static.UpdateFieldUnitAuralevelsLast:                                    uint32(0), // TODO
		static.UpdateFieldUnitAuraapplications:                                  uint32(0), // TODO
		static.UpdateFieldUnitAuraapplicationsLast:                              uint32(0), // TODO
		static.UpdateFieldUnitAurastate:                                         uint32(0), // TODO
		static.UpdateFieldUnitBaseattacktime:                                    uint32(0), // TODO
		static.UpdateFieldUnitOffhandattacktime:                                 uint32(0), // TODO
		static.UpdateFieldUnitRangedattacktime:                                  uint32(0), // TODO
		static.UpdateFieldUnitBoundingradius:                                    uint32(tmpl.BoundingRadius),
		static.UpdateFieldUnitCombatreach:                                       uint32(tmpl.CombatReach),
		static.UpdateFieldUnitDisplayid:                                         uint32(tmpl.DisplayID),
		static.UpdateFieldUnitNativedisplayid:                                   uint32(0), // TODO
		static.UpdateFieldUnitMountdisplayid:                                    uint32(0), // TODO
		static.UpdateFieldUnitMindamage:                                         uint32(0), // TODO
		static.UpdateFieldUnitMaxdamage:                                         uint32(0), // TODO
		static.UpdateFieldUnitMinoffhanddamage:                                  uint32(0), // TODO
		static.UpdateFieldUnitMaxoffhanddamage:                                  uint32(0), // TODO
		static.UpdateFieldUnitBytes1:                                            uint32(0), // TODO
		static.UpdateFieldUnitPetnumber:                                         uint32(0), // TODO
		static.UpdateFieldUnitPetNameTimestamp:                                  uint32(0), // TODO
		static.UpdateFieldUnitPetexperience:                                     uint32(0), // TODO
		static.UpdateFieldUnitPetnextlevelexp:                                   uint32(0), // TODO
		static.UpdateFieldUnitDynamicFlags:                                      uint32(0), // TODO
		static.UpdateFieldUnitChannelSpell:                                      uint32(0), // TODO
		static.UpdateFieldUnitModCastSpeed:                                      uint32(0), // TODO
		static.UpdateFieldUnitCreatedBySpell:                                    uint32(0), // TODO
		static.UpdateFieldUnitNpcFlags:                                          uint32(tmpl.Flags()),
		static.UpdateFieldUnitNpcEmotestate:                                     uint32(0),  // TODO
		static.UpdateFieldUnitTrainingPoints:                                    uint32(0),  // TODO
		static.UpdateFieldUnitStrength:                                          uint32(0),  // TODO
		static.UpdateFieldUnitAgility:                                           uint32(0),  // TODO
		static.UpdateFieldUnitStamina:                                           uint32(0),  // TODO
		static.UpdateFieldUnitIntellect:                                         uint32(0),  // TODO
		static.UpdateFieldUnitSpirit:                                            uint32(0),  // TODO
		static.UpdateFieldUnitArmor:                                             uint32(0),  // TODO
		static.UpdateFieldUnitHolyResist:                                        uint32(0),  // TODO
		static.UpdateFieldUnitFireResist:                                        uint32(0),  // TODO
		static.UpdateFieldUnitNatureResist:                                      uint32(0),  // TODO
		static.UpdateFieldUnitFrostResist:                                       uint32(0),  // TODO
		static.UpdateFieldUnitShadowResist:                                      uint32(0),  // TODO
		static.UpdateFieldUnitArcaneResist:                                      uint32(0),  // TODO
		static.UpdateFieldUnitBaseMana:                                          uint32(0),  // TODO
		static.UpdateFieldUnitBaseHealth:                                        uint32(0),  // TODO
		static.UpdateFieldUnitBytes2:                                            uint32(0),  // TODO
		static.UpdateFieldUnitAttackPower:                                       uint32(0),  // TODO
		static.UpdateFieldUnitAttackPowerMods:                                   uint32(0),  // TODO
		static.UpdateFieldUnitAttackPowerMultiplier:                             uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPower:                                 uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPowerMods:                             uint32(0),  // TODO
		static.UpdateFieldUnitRangedAttackPowerMultiplier:                       uint32(0),  // TODO
		static.UpdateFieldUnitMinrangeddamage:                                   uint32(0),  // TODO
		static.UpdateFieldUnitMaxrangeddamage:                                   uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier:                                 uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier01:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier02:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier03:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier04:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier05:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostModifier06:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier:                               uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier01:                             uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier02:                             uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier03:                             uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier04:                             uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier05:                             uint32(0),  // TODO
		static.UpdateFieldUnitPowerCostMultiplier06:                             uint32(0),  // TODO
		static.UpdateFieldUnitFactiontemplate:                                   uint32(17), // TODO
	}

	mergedFields := u.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[static.UpdateFieldType] = uint32(TypeMask(u))

	return mergedFields
}

/// Utility methods.
// MovementUpdate returns a buffer which represents the movement update object for this unit.
func (u *Unit) MovementUpdate() []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(u.movementFlags()))
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time

	binary.Write(buffer, binary.LittleEndian, float32(u.GetLocation().X))
	binary.Write(buffer, binary.LittleEndian, float32(u.GetLocation().Y))
	binary.Write(buffer, binary.LittleEndian, float32(u.GetLocation().Z))
	binary.Write(buffer, binary.LittleEndian, float32(u.GetLocation().O))

	// if GetObjectManager().Exists(u.Transport) {
	// TODO(jeshua): transport.
	// transportObj := GetObjectManager().GetTransport(u.Transport)
	// binary.Write(buffer, binary.LittleEndian, uint64(u.Transport))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.GetLocation().X))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.GetLocation().Y))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.GetLocation().Z))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.GetLocation().O))
	// binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time
	// }

	if u.IsSwimming {
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Pitch))
	}

	if !GetObjectManager().Exists(u.Transport) {
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // LastFallTime
	}

	if u.IsFalling {
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Jump.Velocity))
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Jump.SinAngle))
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Jump.CosAngle))
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Jump.XYSpeed))
	}

	// SplineElevation update goes HERE.

	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedWalk))
	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedRun))
	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedRunBackward))
	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedSwim))
	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedSwimBackward))
	binary.Write(buffer, binary.LittleEndian, float32(u.SpeedTurn))

	// Spline update goes HERE.

	return buffer.Bytes()
}

// Template returns the unit template this object is based on.
func (u *Unit) Template() *static.Unit {
	return static.Units[int(u.Entry)]
}

func (u *Unit) movementFlags() uint32 {
	movementFlags := uint32(0)
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagForward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagBackward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagStrafeLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagStrafeRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagTurnLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagTurnRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagPitchUp)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagPitchDown)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagWalkMode)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagLevitating)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFlying)
	// }
	if u.IsFalling {
		movementFlags |= uint32(static.MovementFlagFalling)
	}
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFallingFar)
	// }
	if u.IsSwimming {
		movementFlags |= uint32(static.MovementFlagSwimming)
	}
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSplineEnabled)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagCanFly)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFlyingOld)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagOnTransport)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSplineElevation)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagRoot)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagWaterWalking)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSafeFall)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagHover)
	// }
	return movementFlags
}
