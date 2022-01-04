package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
)

// GameObject represents a generic game object.
type GameObject struct {
	guid   interfaces.GUID
	Entry  uint32
	ScaleX float32
}

// Object interface methods.
func (o *GameObject) GUID() interfaces.GUID             { return o.guid }
func (o *GameObject) SetGUID(guid interfaces.GUID)      { o.guid = guid }
func (o *GameObject) GetLocation() *interfaces.Location { return nil }

func (o *GameObject) UpdateFields() interfaces.UpdateFieldsMap {
	return interfaces.UpdateFieldsMap{
		static.UpdateFieldGUIDLow:  uint32(o.GUID().Low()),
		static.UpdateFieldGUIDHigh: uint32(o.GUID().High()),
		static.UpdateFieldEntry:    uint32(o.Entry),
		static.UpdateFieldScaleX:   float32(o.ScaleX),
	}
}
