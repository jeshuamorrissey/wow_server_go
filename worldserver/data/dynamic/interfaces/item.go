package interfaces

import "github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"

type Item interface {
	// GetTemplate retreives the item template this object is based on.
	GetTemplate() *static.Item

	// GetContainer retrives the container that the item is within, or nil if it isn't inside
	// a container.
	GetContainer() GUID
}
