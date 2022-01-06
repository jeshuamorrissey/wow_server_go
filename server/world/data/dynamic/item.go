package dynamic

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
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
	Owner       interfaces.GUID
	Container   interfaces.GUID
	Creator     interfaces.GUID
	GiftCreator interfaces.GUID
}

// Object interface methods.
func (i *Item) GUID() interfaces.GUID        { return i.GameObject.GUID() }
func (i *Item) SetGUID(guid interfaces.GUID) { i.GameObject.SetGUID(guid) }

func (i *Item) GetLocation() *interfaces.Location {
	if container := GetObjectManager().Get(i.Container); container != nil {
		return container.GetLocation()
	}

	return nil
}

func (i *Item) UpdateFields() interfaces.UpdateFieldsMap {
	fields := interfaces.UpdateFieldsMap{
		static.UpdateFieldItemOwner:           uint32(i.Owner.Low()),
		static.UpdateFieldItemOwner + 1:       uint32(i.Owner.High()),
		static.UpdateFieldItemContained:       uint32(i.Container.Low()),
		static.UpdateFieldItemContained + 1:   uint32(i.Container.High()),
		static.UpdateFieldItemCreator:         uint32(i.Creator.Low()),
		static.UpdateFieldItemCreator + 1:     uint32(i.Creator.High()),
		static.UpdateFieldItemGiftCreator:     uint32(i.GiftCreator.Low()),
		static.UpdateFieldItemGiftCreator + 1: uint32(i.GiftCreator.High()),
		static.UpdateFieldItemStackCount:      uint32(i.StackCount),
		static.UpdateFieldItemSpellCharges:    uint32(i.ChargesRemaining),
		static.UpdateFieldItemFlags:           uint32(i.flags()),
		// static.UpdateFieldItemEnchantmentID:       uint32(0), // TODO
		// static.UpdateFieldItemEnchantmentDuration: uint32(0), // TODO
		// static.UpdateFieldItemEnchantmentCharges:  uint32(0), // TODO
		// static.UpdateFieldItemPropertySeed:        uint32(0), // TODO
		// static.UpdateFieldItemRandomPropertiesID:  uint32(0), // TODO
		// static.UpdateFieldItemItemTextID:          uint32(0), // TODO
		static.UpdateFieldItemDurability:    uint32(i.Durability),
		static.UpdateFieldItemMaxDurability: uint32(i.GetTemplate().MaxDurability),
	}

	if i.DurationRemaining == nil {
		fields[static.UpdateFieldItemDuration] = uint32(0)
	} else {
		fields[static.UpdateFieldItemDuration] = uint32(i.DurationRemaining.Second())
	}

	mergedFields := i.GameObject.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[static.UpdateFieldType] = uint32(TypeMask(i))

	return mergedFields
}

// Item interface methods.
func (i *Item) GetTemplate() *static.Item     { return static.Items[int(i.Entry)] }
func (i *Item) GetContainer() interfaces.GUID { return i.Container }

// Utility methods.
func (i *Item) flags() int {
	var flags int
	if i.IsBound {
		flags |= int(static.ItemFlagBound)
	}

	if i.IsUnlocked {
		flags |= int(static.ItemFlagUnlocked)
	}

	if i.IsWrapped {
		flags |= int(static.ItemFlagWrapped)
	}

	if i.IsReadable {
		flags |= int(static.ItemFlagReadable)
	}

	return flags
}
