package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// GameObject represents a generic game object.
type GameObject struct {
	ID     interfaces.GUID
	Entry  uint32
	ScaleX float32
}

// Object interface methods.
func (o *GameObject) GUID() interfaces.GUID             { return o.ID }
func (o *GameObject) SetGUID(guid interfaces.GUID)      { o.ID = guid }
func (o *GameObject) GetLocation() *interfaces.Location { return nil }

func (o *GameObject) UpdateFields() interfaces.UpdateFieldsMap {
	return interfaces.UpdateFieldsMap{
		static.UpdateFieldGUIDLow:  uint32(o.GUID().Low()),
		static.UpdateFieldGUIDHigh: uint32(o.GUID().High()),
		static.UpdateFieldEntry:    uint32(o.Entry),
		static.UpdateFieldScaleX:   float32(o.ScaleX),
	}
}
