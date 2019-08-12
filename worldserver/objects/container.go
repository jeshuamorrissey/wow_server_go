package objects

import c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"

// Container represents the game object for an item container.
type Container struct {
	Item

	NumSlots int
	Items    map[int]*Item
}

// GUID returns the guid of the object.
func (o *Container) GUID() GUID { return o.Item.GUID() }

// SetGUID updates the GUID value of the object.
func (o *Container) SetGUID(guid int) { o.guid = GUID(int(c.HighGUIDContainer)<<32 | guid) }

// HighGUID returns the high GUID component for an object.
func (o *Container) HighGUID() c.HighGUID { return c.HighGUIDContainer }

// Location returns the location of the object.
func (o *Container) Location() *Location { return o.Container.GetLocation() }

// Fields returns the update fields of the object.
func (o *Container) Fields() map[c.UpdateField]interface{} {
	fields := map[c.UpdateField]interface{}{
		c.UpdateFieldContainerNumSlots: o.NumSlots,
	}

	for slot, item := range o.Items {
		fieldStart := c.UpdateField(int(c.UpdateFieldContainerSlot1) + (slot * 2))
		fields[fieldStart] = item.GUID().Low()
		fields[fieldStart+1] = item.GUID().High()
	}

	return mergeUpdateFields(fields, o.Item.Fields())
}
