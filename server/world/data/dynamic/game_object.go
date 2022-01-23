package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/channels"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// GameObject represents a generic game object.
type GameObject struct {
	ID     interfaces.GUID
	Entry  uint32
	ScaleX float32

	// The channel to receive updates on.
	update chan []interface{}
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

func (o *GameObject) UpdateChannel() chan []interface{} {
	return o.update
}

func (o *GameObject) SendUpdates(updates []interface{}) {
	o.UpdateChannel() <- updates
}

func (o *GameObject) CreateUpdateChannel() {
	if o.update == nil {
		o.update = make(chan []interface{}, 100)
	}
}

func (o *GameObject) StartUpdateLoop() {
	if o.UpdateChannel() != nil {
		return
	}

	o.CreateUpdateChannel()
	go func() {
		for {
			for _, update := range <-o.UpdateChannel() {
				switch update.(type) {
				default:
				}
			}

			channels.ObjectUpdates <- o.GUID()
		}
	}()
}
