package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/channels"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// Container represents an instance of an in-game container.
type Container struct {
	Item

	NumSlots int
	Items    map[int]interfaces.GUID
}

// Object interface methods.
func (cn *Container) GUID() interfaces.GUID             { return cn.Item.GUID() }
func (cn *Container) SetGUID(guid interfaces.GUID)      { cn.Item.SetGUID(guid) }
func (cn *Container) GetLocation() *interfaces.Location { return cn.Item.GetLocation() }

func (cn *Container) UpdateFields() interfaces.UpdateFieldsMap {
	fields := interfaces.UpdateFieldsMap{
		static.UpdateFieldContainerNumSlots: uint32(cn.NumSlots),
	}

	for slot, itemGUID := range cn.Items {
		fieldStart := static.UpdateField(int(static.UpdateFieldContainerSlot1) + (slot * 2))
		fields[fieldStart] = uint32(itemGUID.Low())
		fields[fieldStart+1] = uint32(itemGUID.High())
	}

	mergedFields := cn.Item.UpdateFields()
	for k, v := range fields {
		mergedFields[k] = v
	}

	mergedFields[static.UpdateFieldType] = uint32(TypeMask(cn))

	return mergedFields
}

func (cn *Container) StartUpdateLoop() {
	if cn.UpdateChannel() != nil {
		return
	}

	cn.CreateUpdateChannel()
	go func() {
		for {
			for _, update := range <-cn.UpdateChannel() {
				switch update.(type) {
				default:
				}
			}

			channels.ObjectUpdates <- cn.GUID()
		}
	}()
}

// Item interface methods.
func (cn *Container) GetTemplate() *static.Item     { return cn.Item.GetTemplate() }
func (cn *Container) GetContainer() interfaces.GUID { return cn.Item.GetContainer() }
