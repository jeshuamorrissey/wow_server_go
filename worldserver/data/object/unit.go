package object

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

type AttackInfo struct {
	Damage int
}

type UnitInterface interface {
	Object

	// MeleeAttackRate should return the time between melee attacks.
	MeleeAttackRate() time.Duration

	// Attack should calculate what the result of this unit attacking another unit would be.
	Attack(target UnitInterface) AttackInfo
}

// Unit represents an instance of an in-game monster.
type Unit struct {
	GameObject

	// Basic information.
	Level            int
	Race             *dbc.Race
	Class            *dbc.Class
	Gender           c.Gender
	Team             c.Team
	StandState       c.StandState
	FreeTalentPoints int
	Byte1Flags       c.Byte1Flags
	EmoteState       int
	TrainingPoints   int

	// Movement related information.
	MovementInfo      MovementInfo
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

	Transport GUID

	// Relationships.
	Charm      GUID
	CharmedBy  GUID
	Summon     GUID
	SummonedBy GUID
	CreatedBy  GUID
	Target     GUID
	Persuaded  GUID
}

// Manager returns the manager associated with this object.
func (u *Unit) Manager() *Manager { return u.GameObject.Manager() }

// SetManager updates the manager associated with this object.
func (u *Unit) SetManager(manager *Manager) { u.GameObject.SetManager(manager) }

// GUID returns the globally-unique ID of the object.
func (u *Unit) GUID() GUID { return u.GameObject.GUID() }

// SetGUID updates this object's GUID to the given value.
func (u *Unit) SetGUID(guid GUID) { u.GameObject.SetGUID(guid) }

// Location returns the location of the object.
func (u *Unit) Location() *Location { return &u.MovementInfo.Location }

func (u *Unit) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Unit) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

// MovementUpdate calculates and returns the movement update for the
// object.
func (u *Unit) MovementUpdate() []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(u.movementFlags()))
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time

	binary.Write(buffer, binary.LittleEndian, float32(u.Location().X))
	binary.Write(buffer, binary.LittleEndian, float32(u.Location().Y))
	binary.Write(buffer, binary.LittleEndian, float32(u.Location().Z))
	binary.Write(buffer, binary.LittleEndian, float32(u.Location().O))

	if u.Manager().Exists(u.Transport) {
		transportObj := u.Manager().Get(u.Transport)
		binary.Write(buffer, binary.LittleEndian, uint64(u.Transport))
		binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().X))
		binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().Y))
		binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().Z))
		binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().O))
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time
	}

	if u.IsSwimming {
		binary.Write(buffer, binary.LittleEndian, float32(u.MovementInfo.Pitch))
	}

	if !u.Manager().Exists(u.Transport) {
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

// UpdateFields populates and returns the updated fields for the
// object.
func (u *Unit) UpdateFields() UpdateFieldsMap {
	tmpl := u.Template()
	fields := UpdateFieldsMap{
		c.UpdateFieldUnitCharmLow:                                     uint32(u.Charm.Low()),
		c.UpdateFieldUnitCharmHigh:                                    uint32(u.Charm.High()),
		c.UpdateFieldUnitSummonLow:                                    uint32(u.Summon.Low()),
		c.UpdateFieldUnitSummonHigh:                                   uint32(u.Summon.High()),
		c.UpdateFieldUnitCharmedbyLow:                                 uint32(u.CharmedBy.Low()),
		c.UpdateFieldUnitCharmedbyHigh:                                uint32(u.CharmedBy.High()),
		c.UpdateFieldUnitSummonedbyLow:                                uint32(u.SummonedBy.Low()),
		c.UpdateFieldUnitSummonedbyHigh:                               uint32(u.SummonedBy.High()),
		c.UpdateFieldUnitCreatedbyLow:                                 uint32(u.CreatedBy.Low()),
		c.UpdateFieldUnitCreatedbyHigh:                                uint32(u.CreatedBy.High()),
		c.UpdateFieldUnitTargetLow:                                    uint32(u.Target.Low()),
		c.UpdateFieldUnitTargetHigh:                                   uint32(u.Target.High()),
		c.UpdateFieldUnitPersuadedLow:                                 uint32(u.Persuaded.Low()),
		c.UpdateFieldUnitPersuadedHigh:                                uint32(u.Persuaded.High()),
		c.UpdateFieldUnitChannelObjectLow:                             uint32(0), // TODO
		c.UpdateFieldUnitChannelObjectHigh:                            uint32(0), // TODO
		c.UpdateFieldUnitHealth:                                       uint32(float32(tmpl.MaxHealth) * u.HealthPercent),
		c.UpdateFieldUnitPowerStart + c.UpdateField(u.powerType()):    uint32(float32(tmpl.MaxPower) * u.PowerPercent),
		c.UpdateFieldUnitMaxHealth:                                    uint32(tmpl.MaxHealth),
		c.UpdateFieldUnitMaxPowerStart + c.UpdateField(u.powerType()): uint32(tmpl.MaxPower),
		c.UpdateFieldUnitLevel:                                        uint32(u.Level),
		c.UpdateFieldUnitBytes0:                                       uint32(u.Race.ID) | uint32(u.Class.ID)<<8 | uint32(u.Gender)<<16,
		c.UpdateFieldUnitFlags:                                        uint32(tmpl.Flags()),
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
		c.UpdateFieldUnitBaseattacktime:                               uint32(0), // TODO
		c.UpdateFieldUnitOffhandattacktime:                            uint32(0), // TODO
		c.UpdateFieldUnitRangedattacktime:                             uint32(0), // TODO
		c.UpdateFieldUnitBoundingradius:                               uint32(tmpl.BoundingRadius),
		c.UpdateFieldUnitCombatreach:                                  uint32(tmpl.CombatReach),
		c.UpdateFieldUnitDisplayid:                                    uint32(tmpl.DisplayID),
		c.UpdateFieldUnitNativedisplayid:                              uint32(0), // TODO
		c.UpdateFieldUnitMountdisplayid:                               uint32(0), // TODO
		c.UpdateFieldUnitMindamage:                                    uint32(0), // TODO
		c.UpdateFieldUnitMaxdamage:                                    uint32(0), // TODO
		c.UpdateFieldUnitMinoffhanddamage:                             uint32(0), // TODO
		c.UpdateFieldUnitMaxoffhanddamage:                             uint32(0), // TODO
		c.UpdateFieldUnitBytes1:                                       uint32(0), // TODO
		c.UpdateFieldUnitPetnumber:                                    uint32(0), // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             uint32(0), // TODO
		c.UpdateFieldUnitPetexperience:                                uint32(0), // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              uint32(0), // TODO
		c.UpdateFieldUnitDynamicFlags:                                 uint32(0), // TODO
		c.UpdateFieldUnitChannelSpell:                                 uint32(0), // TODO
		c.UpdateFieldUnitModCastSpeed:                                 uint32(0), // TODO
		c.UpdateFieldUnitCreatedBySpell:                               uint32(0), // TODO
		c.UpdateFieldUnitNpcFlags:                                     uint32(tmpl.Flags()),
		c.UpdateFieldUnitNpcEmotestate:                                uint32(0),  // TODO
		c.UpdateFieldUnitTrainingPoints:                               uint32(0),  // TODO
		c.UpdateFieldUnitStrength:                                     uint32(0),  // TODO
		c.UpdateFieldUnitAgility:                                      uint32(0),  // TODO
		c.UpdateFieldUnitStamina:                                      uint32(0),  // TODO
		c.UpdateFieldUnitIntellect:                                    uint32(0),  // TODO
		c.UpdateFieldUnitSpirit:                                       uint32(0),  // TODO
		c.UpdateFieldUnitArmor:                                        uint32(0),  // TODO
		c.UpdateFieldUnitHolyResist:                                   uint32(0),  // TODO
		c.UpdateFieldUnitFireResist:                                   uint32(0),  // TODO
		c.UpdateFieldUnitNatureResist:                                 uint32(0),  // TODO
		c.UpdateFieldUnitFrostResist:                                  uint32(0),  // TODO
		c.UpdateFieldUnitShadowResist:                                 uint32(0),  // TODO
		c.UpdateFieldUnitArcaneResist:                                 uint32(0),  // TODO
		c.UpdateFieldUnitBaseMana:                                     uint32(0),  // TODO
		c.UpdateFieldUnitBaseHealth:                                   uint32(0),  // TODO
		c.UpdateFieldUnitBytes2:                                       uint32(0),  // TODO
		c.UpdateFieldUnitAttackPower:                                  uint32(0),  // TODO
		c.UpdateFieldUnitAttackPowerMods:                              uint32(0),  // TODO
		c.UpdateFieldUnitAttackPowerMultiplier:                        uint32(0),  // TODO
		c.UpdateFieldUnitRangedAttackPower:                            uint32(0),  // TODO
		c.UpdateFieldUnitRangedAttackPowerMods:                        uint32(0),  // TODO
		c.UpdateFieldUnitRangedAttackPowerMultiplier:                  uint32(0),  // TODO
		c.UpdateFieldUnitMinrangeddamage:                              uint32(0),  // TODO
		c.UpdateFieldUnitMaxrangeddamage:                              uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier:                            uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier01:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier02:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier03:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier04:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier05:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostModifier06:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier:                          uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier01:                        uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier02:                        uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier03:                        uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier04:                        uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier05:                        uint32(0),  // TODO
		c.UpdateFieldUnitPowerCostMultiplier06:                        uint32(0),  // TODO
		c.UpdateFieldUnitFactiontemplate:                              uint32(17), // TODO
	}

	mergedFields := u.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[c.UpdateFieldType] = uint32(TypeMask(u))

	return mergedFields
}

// Template returns the item template this object is based on.
func (u *Unit) Template() *dbc.Unit {
	return dbc.Units[int(u.Entry)]
}

func (u *Unit) movementFlags() uint32 {
	movementFlags := uint32(0)
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagForward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagBackward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagStrafeLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagStrafeRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagTurnLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagTurnRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagPitchUp)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagPitchDown)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagWalkMode)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagLevitating)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagFlying)
	// }
	if u.IsFalling {
		movementFlags |= uint32(c.MovementFlagFalling)
	}
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagFallingFar)
	// }
	if u.IsSwimming {
		movementFlags |= uint32(c.MovementFlagSwimming)
	}
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagSplineEnabled)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagCanFly)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagFlyingOld)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagOnTransport)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagSplineElevation)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagRoot)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagWaterWalking)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagSafeFall)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(c.MovementFlagHover)
	// }
	return movementFlags
}
