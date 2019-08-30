package object

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/sirupsen/logrus"
)

// Unit represents an instance of an in-game monster.
type Unit struct {
	GameObject

	// Basic information.
	Loc              Location
	Pitch            float32
	Level            int
	Race             c.Race
	Class            c.Class
	Gender           c.Gender
	Team             c.Team
	StandState       c.StandState
	FreeTalentPoints int
	Byte1Flags       c.Byte1Flags
	EmoteState       int
	TrainingPoints   int

	SpeedWalk         float32
	SpeedRun          float32
	SpeedRunBackward  float32
	SpeedSwim         float32
	SpeedSwimBackward float32
	SpeedTurn         float32

	Velocity float32
	SinAngle float32
	CosAngle float32
	XYSpeed  float32

	// Stats.
	BaseHealth int
	Health     int
	BasePower  int
	Power      int

	Strength  int
	Agility   int
	Stamina   int
	Intellect int
	Spirit    int

	Resistances map[c.SpellSchool]int

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
func (u *Unit) Location() *Location { return &u.Loc }

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

	// if u.Manager().Exists(u.Transport) {
	// 	transportObj := u.Manager().Get(u.Transport)
	// 	binary.Write(buffer, binary.LittleEndian, uint64(u.Transport))
	// 	binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().X))
	// 	binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().Y))
	// 	binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().Z))
	// 	binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location().O))
	// 	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time
	// }

	// if u.IsSwimming {
	// 	binary.Write(buffer, binary.LittleEndian, float32(u.Pitch))
	// }

	// if !u.Manager().Exists(u.Transport) {
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // LastFallTime
	// }

	// if u.IsFalling {
	// 	binary.Write(buffer, binary.LittleEndian, float32(u.Velocity))
	// 	binary.Write(buffer, binary.LittleEndian, float32(u.SinAngle))
	// 	binary.Write(buffer, binary.LittleEndian, float32(u.CosAngle))
	// 	binary.Write(buffer, binary.LittleEndian, float32(u.XYSpeed))
	// }

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
		c.UpdateFieldUnitCharmLow:                                     u.Charm.Low(),
		c.UpdateFieldUnitCharmHigh:                                    u.Charm.High(),
		c.UpdateFieldUnitSummonLow:                                    u.Summon.Low(),
		c.UpdateFieldUnitSummonHigh:                                   u.Summon.High(),
		c.UpdateFieldUnitCharmedbyLow:                                 u.CharmedBy.Low(),
		c.UpdateFieldUnitCharmedbyHigh:                                u.CharmedBy.High(),
		c.UpdateFieldUnitSummonedbyLow:                                u.SummonedBy.Low(),
		c.UpdateFieldUnitSummonedbyHigh:                               u.SummonedBy.High(),
		c.UpdateFieldUnitCreatedbyLow:                                 u.CreatedBy.Low(),
		c.UpdateFieldUnitCreatedbyHigh:                                u.CreatedBy.High(),
		c.UpdateFieldUnitTargetLow:                                    u.Target.Low(),
		c.UpdateFieldUnitTargetHigh:                                   u.Target.High(),
		c.UpdateFieldUnitPersuadedLow:                                 u.Persuaded.Low(),
		c.UpdateFieldUnitPersuadedHigh:                                u.Persuaded.High(),
		c.UpdateFieldUnitChannelObjectLow:                             0, // TODO
		c.UpdateFieldUnitChannelObjectHigh:                            0, // TODO
		c.UpdateFieldUnitHealth:                                       u.Health,
		c.UpdateFieldUnitPowerStart + c.UpdateField(u.powerType()):    u.Power,
		c.UpdateFieldUnitMaxHealth:                                    u.maxHealth(),
		c.UpdateFieldUnitMaxPowerStart + c.UpdateField(u.powerType()): u.maxPower(),
		c.UpdateFieldUnitLevel:                                        u.Level,
		c.UpdateFieldUnitBytes0:                                       int(u.Race) | int(u.Class)<<8 | int(u.Gender)<<16,
		c.UpdateFieldUnitFlags:                                        tmpl.Flags(),
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
		c.UpdateFieldUnitBaseattacktime:                               tmpl.MeleeBaseAttackTime,
		c.UpdateFieldUnitOffhandattacktime:                            tmpl.MeleeBaseAttackTime,
		c.UpdateFieldUnitRangedattacktime:                             tmpl.RangedBaseAttackTime,
		c.UpdateFieldUnitBoundingradius:                               tmpl.Models[0].BoundingRadius,
		c.UpdateFieldUnitCombatreach:                                  tmpl.Models[0].CombatReach,
		c.UpdateFieldUnitDisplayid:                                    tmpl.Models[0].ID,
		c.UpdateFieldUnitNativedisplayid:                              0, // TODO
		c.UpdateFieldUnitMountdisplayid:                               0, // TODO
		c.UpdateFieldUnitMindamage:                                    tmpl.MinMeleeDmg,
		c.UpdateFieldUnitMaxdamage:                                    tmpl.MaxMeleeDmg,
		c.UpdateFieldUnitMinoffhanddamage:                             tmpl.MinMeleeDmg,
		c.UpdateFieldUnitMaxoffhanddamage:                             tmpl.MaxMeleeDmg,
		c.UpdateFieldUnitBytes1:                                       int(u.Byte1Flags)<<24 | u.FreeTalentPoints<<16 | int(u.StandState),
		c.UpdateFieldUnitPetnumber:                                    0, // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             0, // TODO
		c.UpdateFieldUnitPetexperience:                                0, // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              0, // TODO
		c.UpdateFieldUnitDynamicFlags:                                 tmpl.DynamicFlags,
		c.UpdateFieldUnitChannelSpell:                                 0, // TODO
		c.UpdateFieldUnitModCastSpeed:                                 1.0,
		c.UpdateFieldUnitCreatedBySpell:                               0, // TODO
		c.UpdateFieldUnitNpcFlags:                                     tmpl.Flags(),
		c.UpdateFieldUnitNpcEmotestate:                                u.EmoteState,
		c.UpdateFieldUnitTrainingPoints:                               u.TrainingPoints,
		c.UpdateFieldUnitStrength:                                     u.Strength,
		c.UpdateFieldUnitAgility:                                      u.Agility,
		c.UpdateFieldUnitStamina:                                      u.Stamina,
		c.UpdateFieldUnitIntellect:                                    u.Intellect,
		c.UpdateFieldUnitSpirit:                                       u.Spirit,
		c.UpdateFieldUnitArmor:                                        u.Resistances[c.SpellSchoolPhysical],
		c.UpdateFieldUnitHolyResist:                                   u.Resistances[c.SpellSchoolHoly],
		c.UpdateFieldUnitFireResist:                                   u.Resistances[c.SpellSchoolFire],
		c.UpdateFieldUnitNatureResist:                                 u.Resistances[c.SpellSchoolNature],
		c.UpdateFieldUnitFrostResist:                                  u.Resistances[c.SpellSchoolFrost],
		c.UpdateFieldUnitShadowResist:                                 u.Resistances[c.SpellSchoolShadow],
		c.UpdateFieldUnitArcaneResist:                                 u.Resistances[c.SpellSchoolArcane],
		c.UpdateFieldUnitBaseMana:                                     u.BasePower,
		c.UpdateFieldUnitBaseHealth:                                   u.BaseHealth,
		c.UpdateFieldUnitBytes2:                                       0, // TODO
		c.UpdateFieldUnitAttackPower:                                  tmpl.MeleeAttackPower,
		c.UpdateFieldUnitAttackPowerMods:                              0, // TODO
		c.UpdateFieldUnitAttackPowerMultiplier:                        tmpl.PowerMultiplier,
		c.UpdateFieldUnitRangedAttackPower:                            tmpl.RangedAttackPower,
		c.UpdateFieldUnitRangedAttackPowerMods:                        0, // TODO
		c.UpdateFieldUnitRangedAttackPowerMultiplier:                  tmpl.PowerMultiplier,
		c.UpdateFieldUnitMinrangeddamage:                              tmpl.MinRangedDmg,
		c.UpdateFieldUnitMaxrangeddamage:                              tmpl.MaxRangedDmg,
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
	}

	if u.Team == c.TeamAlliance {
		fields[c.UpdateFieldUnitFactiontemplate] = tmpl.FactionAlliance
	} else {
		fields[c.UpdateFieldUnitFactiontemplate] = tmpl.FactionHorde
	}

	for i, itemEntry := range u.VirtualItems {
		item, ok := dbc.Items[itemEntry]
		if !ok {
			u.Manager().log.WithFields(logrus.Fields{
				"unit":       u.GUID(),
				"item_entry": itemEntry,
				"item_slot":  i,
			}).Errorf("Unknown VirtualItem")
		}

		displayField := c.UpdateFieldUnitVirtualItemDisplay + c.UpdateField(i)
		fields[displayField] = item.DisplayID

		infoField := c.UpdateFieldUnitVirtualItemInfo + c.UpdateField(i*2)
		fields[infoField] = (int(item.Class)<<24 | int(item.SubClass)<<16 | int(item.Material)<<8 | int(item.InventoryType))
		fields[infoField+1] = item.SheathType
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
	var movementFlags uint32
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
