package object

import (
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Container represents an instance of an in-game container.
type Container struct {
	Item

	NumSlots int
	Items    map[int]GUID
}

// Manager returns the manager associated with this object.
func (cn *Container) Manager() *Manager { return cn.gameObject.Manager() }

// SetManager updates the manager associated with this object.
func (cn *Container) SetManager(manager *Manager) { cn.gameObject.SetManager(manager) }

// GUID returns the globally-unique ID of the object.
func (cn *Container) GUID() GUID { return cn.gameObject.GUID() }

// SetGUID updates this object's GUID to the given value.
func (cn *Container) SetGUID(guid GUID) { cn.gameObject.SetGUID(guid) }

// Location returns the location of the object.
func (cn *Container) Location() *Location {
	obj, err := cn.Manager().Get(cn.Container)
	if err != nil {
		return nil
	}

	return obj.Location()
}

// MovementUpdate calculates and returns the movement update for the
// object.
func (cn *Container) MovementUpdate() []byte { return nil }

// UpdateFields populates and returns the updated fields for the
// object.
func (cn *Container) UpdateFields() map[c.UpdateField]interface{} {
	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldContainerNumSlots: cn.NumSlots,
	}

	for slot, itemGUID := range cn.Items {
		fieldStart := c.UpdateField(int(c.UpdateFieldContainerSlot1) + (slot * 2))
		fields[fieldStart] = itemGUID.Low()
		fields[fieldStart+1] = itemGUID.High()
	}

	baseFields := cn.Item.UpdateFields()
	for k, v := range baseFields {
		fields[k] = v
	}

	return fields
}
