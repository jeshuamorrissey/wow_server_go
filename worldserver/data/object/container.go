package object

import (
	"encoding/json"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Container represents an instance of an in-game container.
type Container struct {
	Item

	NumSlots int
	Items    map[int]GUID
}

// Manager returns the manager associated with this object.
func (cn *Container) Manager() *Manager { return cn.GameObject.Manager() }

// SetManager updates the manager associated with this object.
func (cn *Container) SetManager(manager *Manager) { cn.GameObject.SetManager(manager) }

// GUID returns the globally-unique ID of the object.
func (cn *Container) GUID() GUID { return cn.GameObject.GUID() }

// SetGUID updates this object's GUID to the given value.
func (cn *Container) SetGUID(guid GUID) { cn.GameObject.SetGUID(guid) }

// Location returns the location of the object.
func (cn *Container) Location() *Location {
	if !cn.Manager().Exists(cn.Container) {
		return nil
	}

	return cn.Manager().Get(cn.Container).Location()
}

func (cn *Container) MarshalJSON() ([]byte, error) {
	return json.Marshal(cn)
}

func (cn *Container) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, cn)
}

// MovementUpdate calculates and returns the movement update for the
// object.
func (cn *Container) MovementUpdate() []byte { return nil }

// UpdateFields populates and returns the updated fields for the
// object.
func (cn *Container) UpdateFields() UpdateFieldsMap {
	fields := UpdateFieldsMap{
		c.UpdateFieldContainerNumSlots: uint32(cn.NumSlots),
	}

	for slot, itemGUID := range cn.Items {
		fieldStart := c.UpdateField(int(c.UpdateFieldContainerSlot1) + (slot * 2))
		fields[fieldStart] = uint32(itemGUID.Low())
		fields[fieldStart+1] = uint32(itemGUID.High())
	}

	mergedFields := cn.Item.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[c.UpdateFieldType] = uint32(TypeMask(cn))

	return mergedFields
}
