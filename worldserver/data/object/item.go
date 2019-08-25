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
func (i *Item) UpdateFields() map[c.UpdateField]interface{} {
	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldItemOwner:               i.Owner.Low(),
		c.UpdateFieldItemOwner + 1:           i.Owner.High(),
		c.UpdateFieldItemContained:           i.Container.Low(),
		c.UpdateFieldItemContained + 1:       i.Container.High(),
		c.UpdateFieldItemCreator:             i.Creator.Low(),
		c.UpdateFieldItemCreator + 1:         i.Creator.High(),
		c.UpdateFieldItemGiftCreator:         i.GiftCreator.Low(),
		c.UpdateFieldItemGiftCreator + 1:     i.GiftCreator.High(),
		c.UpdateFieldItemStackCount:          i.StackCount,
		c.UpdateFieldItemDuration:            i.DurationRemaining.Second(),
		c.UpdateFieldItemSpellCharges:        i.ChargesRemaining,
		c.UpdateFieldItemFlags:               i.flags(),
		c.UpdateFieldItemEnchantmentID:       0, // TODO
		c.UpdateFieldItemEnchantmentDuration: 0, // TODO
		c.UpdateFieldItemEnchantmentCharges:  0, // TODO
		c.UpdateFieldItemPropertySeed:        0, // TODO
		c.UpdateFieldItemRandomPropertiesID:  0, // TODO
		c.UpdateFieldItemItemTextID:          0, // TODO
		c.UpdateFieldItemDurability:          i.Durability,
		c.UpdateFieldItemMaxDurability:       i.Template().MaxDurability,
	}

	baseFields := i.GameObject.UpdateFields()
	for k, v := range baseFields {
		fields[k] = v
	}

	return fields
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
