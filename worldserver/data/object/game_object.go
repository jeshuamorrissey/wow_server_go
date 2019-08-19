package object

import c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"

type gameObject struct {
	manager *Manager

	guid   GUID
	Entry  uint32
	ScaleX float32
}

// Manager returns the manager associated with this object.
func (o *gameObject) Manager() *Manager { return o.manager }

// SetManager updates the manager associated with this object.
func (o *gameObject) SetManager(manager *Manager) { o.manager = manager }

// GUID returns the globally-unique ID of the object.
func (o *gameObject) GUID() GUID { return o.guid }

// SetGUID updates this object's GUID to the given value.
func (o *gameObject) SetGUID(guid GUID) { o.guid = guid }

// Location returns the location of the object.
func (o *gameObject) Location() *Location { return nil }

// MovementUpdate calculates and returns the movement update for the
// object.
func (o *gameObject) MovementUpdate() []byte { return nil }

// UpdateFields populates and returns the updated fields for the
// object.
func (o *gameObject) UpdateFields() map[c.UpdateField]interface{} {
	return map[c.UpdateField]interface{}{
		c.UpdateFieldGUIDLow:  uint32(o.GUID().Low()),
		c.UpdateFieldGUIDHigh: uint32(o.GUID().High()),
		c.UpdateFieldType:     uint32(TypeMask(o)),
		c.UpdateFieldEntry:    uint32(o.Entry),
		c.UpdateFieldScaleX:   float32(o.ScaleX),
	}
}
