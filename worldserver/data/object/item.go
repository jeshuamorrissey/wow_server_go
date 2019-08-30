package object

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Item represents an instance of an in-game item.
type Item struct {
	GameObject

	// Basic information.
	Durability        int
	StackCount        int
	DurationRemaining *time.Time
	ChargesRemaining  int

	// Flags.
	IsBound    bool
	IsUnlocked bool
	IsWrapped  bool
	IsReadable bool

	// Relationships.
	Owner       GUID
	Container   GUID
	Creator     GUID
	GiftCreator GUID
}

// Manager returns the manager associated with this object.
func (i *Item) Manager() *Manager { return i.GameObject.Manager() }

// SetManager updates the manager associated with this object.
func (i *Item) SetManager(manager *Manager) { i.GameObject.SetManager(manager) }

// GUID returns the globally-unique ID of the object.
func (i *Item) GUID() GUID { return i.GameObject.GUID() }

// SetGUID updates this object's GUID to the given value.
func (i *Item) SetGUID(guid GUID) { i.GameObject.SetGUID(guid) }

// Location returns the location of the object.
func (i *Item) Location() *Location {
	if !i.Manager().Exists(i.Container) {
		return nil
	}

	return i.Manager().Get(i.Container).Location()
}

// MovementUpdate calculates and returns the movement update for the
// object.
func (i *Item) MovementUpdate() []byte { return nil }

// UpdateFields populates and returns the updated fields for the
// object.
func (i *Item) UpdateFields() UpdateFieldsMap {
	fields := UpdateFieldsMap{
		c.UpdateFieldItemOwner:               uint32(i.Owner.Low()),
		c.UpdateFieldItemOwner + 1:           uint32(i.Owner.High()),
		c.UpdateFieldItemContained:           uint32(i.Container.Low()),
		c.UpdateFieldItemContained + 1:       uint32(i.Container.High()),
		c.UpdateFieldItemCreator:             uint32(i.Creator.Low()),
		c.UpdateFieldItemCreator + 1:         uint32(i.Creator.High()),
		c.UpdateFieldItemGiftCreator:         uint32(i.GiftCreator.Low()),
		c.UpdateFieldItemGiftCreator + 1:     uint32(i.GiftCreator.High()),
		c.UpdateFieldItemStackCount:          uint32(i.StackCount),
		c.UpdateFieldItemSpellCharges:        uint32(i.ChargesRemaining),
		c.UpdateFieldItemFlags:               uint32(i.flags()),
		c.UpdateFieldItemEnchantmentID:       uint32(0), // TODO
		c.UpdateFieldItemEnchantmentDuration: uint32(0), // TODO
		c.UpdateFieldItemEnchantmentCharges:  uint32(0), // TODO
		c.UpdateFieldItemPropertySeed:        uint32(0), // TODO
		c.UpdateFieldItemRandomPropertiesID:  uint32(0), // TODO
		c.UpdateFieldItemItemTextID:          uint32(0), // TODO
		c.UpdateFieldItemDurability:          uint32(i.Durability),
		c.UpdateFieldItemMaxDurability:       uint32(i.Template().MaxDurability),
	}

	if i.DurationRemaining == nil {
		fields[c.UpdateFieldItemDuration] = uint32(0)
	} else {
		fields[c.UpdateFieldItemDuration] = uint32(i.DurationRemaining.Second())
	}

	mergedFields := i.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[c.UpdateFieldType] = uint32(TypeMask(i))

	return mergedFields
}

// Template returns the item template this object is based on.
func (i *Item) Template() *dbc.Item {
	return dbc.Items[int(i.Entry)]
}

func (i *Item) flags() int {
	var flags int
	if i.IsBound {
		flags |= int(c.ItemFlagBound)
	}

	if i.IsUnlocked {
		flags |= int(c.ItemFlagUnlocked)
	}

	if i.IsWrapped {
		flags |= int(c.ItemFlagWrapped)
	}

	if i.IsReadable {
		flags |= int(c.ItemFlagReadable)
	}

	return flags
}
