package objects

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GameObject is a common interface to access in-game objects through.
type GameObject interface {
	GUID() GUID
	SetGUID(int)
	HighGUID() c.HighGUID
	GetLocation() *Location
	Fields() map[c.UpdateField]interface{}
}

// BaseGameObject represents the common fields of all game objects.
type BaseGameObject struct {
	guid   GUID
	Entry  int
	ScaleX float32
}

// GUID returns the guid of the object.
func (o *BaseGameObject) GUID() GUID { return o.guid }

// SetGUID updates the GUID value of the object.
func (o *BaseGameObject) SetGUID(guid int) { o.guid = GUID(guid) }

// HighGUID returns the high GUID component for an object.
func (o *BaseGameObject) HighGUID() c.HighGUID { return 0 }

// GetLocation returns the location of the object.
func (o *BaseGameObject) GetLocation() *Location { return nil }

// Fields returns the update fields of the object.
func (o *BaseGameObject) Fields() map[c.UpdateField]interface{} {
	return map[c.UpdateField]interface{}{
		c.UpdateFieldGUIDLow:  o.guid.Low(),
		c.UpdateFieldGUIDHigh: o.guid.High(),
		c.UpdateFieldType:     TypeMask(o),
		c.UpdateFieldEntry:    o.Entry,
		c.UpdateFieldScaleX:   o.ScaleX,
	}
}
