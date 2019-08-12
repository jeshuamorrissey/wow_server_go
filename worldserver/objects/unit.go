package objects

import (
	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

type Unit struct {
	BaseGameObject

	// Basic information.
	Location         Location
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

	// Display items (virtual).
	VirtualItems []*data.Item

	// Flags
	CanDetectAmore0     bool
	CanDetectAmore1     bool
	CanDetectAmore2     bool
	CanDetectAmore3     bool
	IsStealth           bool
	HasInvisibilityGlow bool

	// Relationships.
	Charm      GUID
	CharmedBy  GUID
	Summon     GUID
	SummonedBy GUID
	CreatedBy  GUID
	Target     GUID
	Persuaded  GUID
}

// GUID returns the guid of the object.
func (o *Unit) GUID() GUID { return o.BaseGameObject.GUID() }

// SetGUID updates the GUID value of the object.
func (o *Unit) SetGUID(guid int) { o.guid = GUID(uint64(c.HighGUIDUnit)<<32 | uint64(guid)) }

// HighGUID returns the high GUID component for an object.
func (o *Unit) HighGUID() c.HighGUID { return c.HighGUIDUnit }

// GetLocation returns the location of the object.
func (o *Unit) GetLocation() *Location { return &o.Location }

// Fields returns the update fields of the object.
func (o *Unit) Fields() map[c.UpdateField]interface{} {
	tmpl := o.template()
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
		c.UpdateFieldUnitBytes1:                                       int(o.Byte1Flags)<<24 | o.FreeTalentPoints<<16 | int(o.StandState),
		c.UpdateFieldUnitPetnumber:                                    0, // TODO
		c.UpdateFieldUnitPetNameTimestamp:                             0, // TODO
		c.UpdateFieldUnitPetexperience:                                0, // TODO
		c.UpdateFieldUnitPetnextlevelexp:                              0, // TODO
		c.UpdateFieldUnitDynamicFlags:                                 tmpl.DynamicFlags,
		c.UpdateFieldUnitChannelSpell:                                 0, // TODO
		c.UpdateFieldUnitModCastSpeed:                                 1.0,
		c.UpdateFieldUnitCreatedBySpell:                               0, // TODO
		c.UpdateFieldUnitNpcFlags:                                     tmpl.Flags(),
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
		c.UpdateFieldUnitBytes2:                                       o.bytes2(),
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

	if o.Team == c.TeamAlliance {
		fields[c.UpdateFieldUnitFactiontemplate] = o.template().FactionAlliance
	} else {
		fields[c.UpdateFieldUnitFactiontemplate] = o.template().FactionHorde
	}

	for i, item := range o.VirtualItems {
		displayField := c.UpdateFieldUnitVirtualItemDisplay + c.UpdateField(i)
		fields[displayField] = item.DisplayID

		infoField := c.UpdateFieldUnitVirtualItemInfo + c.UpdateField(i*2)
		fields[infoField] = (int(item.Class)<<24 | int(item.SubClass)<<16 | int(item.Material)<<8 | int(item.InventoryType))
		fields[infoField+1] = item.SheathType
	}

	return mergeUpdateFields(fields, o.BaseGameObject.Fields())
}

func (o *Unit) template() *data.Unit {
	return data.Units[o.Entry]
}

func (o *Unit) bytes2() int {
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
