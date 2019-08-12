package objects

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// Item represents the game object for an item.
type Item struct {
	BaseGameObject

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
	Owner       GameObject
	Container   GameObject
	Creator     GameObject
	GiftCreator GameObject
}

// GUID returns the guid of the object.
func (o *Item) GUID() GUID { return o.BaseGameObject.GUID() }

// SetGUID updates the GUID value of the object.
func (o *Item) SetGUID(guid int) { o.guid = GUID(int(c.HighGUIDItem)<<32 | guid) }

// HighGUID returns the high GUID component for an object.
func (o *Item) HighGUID() c.HighGUID { return c.HighGUIDItem }

// GetLocation returns the location of the object.
func (o *Item) GetLocation() *Location { return o.Container.GetLocation() }

// Fields returns the update fields of the object.
func (o *Item) Fields() map[c.UpdateField]interface{} {
	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldItemOwner:               o.Owner.GUID().Low(),
		c.UpdateFieldItemOwner + 1:           o.Owner.GUID().High(),
		c.UpdateFieldItemContained:           o.Container.GUID().Low(),
		c.UpdateFieldItemContained + 1:       o.Container.GUID().High(),
		c.UpdateFieldItemCreator:             o.Creator.GUID().Low(),
		c.UpdateFieldItemCreator + 1:         o.Creator.GUID().High(),
		c.UpdateFieldItemGiftCreator:         o.GiftCreator.GUID().Low(),
		c.UpdateFieldItemGiftCreator + 1:     o.GiftCreator.GUID().High(),
		c.UpdateFieldItemStackCount:          o.StackCount,
		c.UpdateFieldItemDuration:            o.DurationRemaining.Second(),
		c.UpdateFieldItemSpellCharges:        o.ChargesRemaining,
		c.UpdateFieldItemFlags:               o.flags(),
		c.UpdateFieldItemEnchantmentID:       0, // TODO
		c.UpdateFieldItemEnchantmentDuration: 0, // TODO
		c.UpdateFieldItemEnchantmentCharges:  0, // TODO
		c.UpdateFieldItemPropertySeed:        0, // TODO
		c.UpdateFieldItemRandomPropertiesID:  0, // TODO
		c.UpdateFieldItemItemTextID:          0, // TODO
		c.UpdateFieldItemDurability:          o.Durability,
		c.UpdateFieldItemMaxDurability:       o.Template().MaxDurability,
	}

	return mergeUpdateFields(fields, o.BaseGameObject.Fields())
}

// Template returns the item template this object is based on.
func (o *Item) Template() *data.Item {
	return data.Items[o.Entry]
}

func (o *Item) flags() int {
	var flags int
	if o.IsBound {
		flags |= int(c.ItemFlagBound)
	}

	if o.IsUnlocked {
		flags |= int(c.ItemFlagUnlocked)
	}

	if o.IsWrapped {
		flags |= int(c.ItemFlagWrapped)
	}

	if o.IsReadable {
		flags |= int(c.ItemFlagReadable)
	}

	return flags
}
