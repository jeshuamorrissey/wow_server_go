package object

import (
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Object represents a generic object within the game. All objects
// should implement this interface.
type Object interface {
	// Manager should return the manager associated with this object.
	Manager() *Manager

	// SetManager should update the manager to the given value.
	SetManager(*Manager)

	// GUID should return the full GUID of the object.
	GUID() GUID

	// SetGUID should set the GUID of the object.
	SetGUID(GUID)

	// Location should return the location within the game world. If the
	// object has no actual location, it should use the location of
	// it's container.
	Location() *Location

	// MovementUpdate should return the full movement update for the
	// object.
	MovementUpdate() []byte

	// UpdateFields should return the update fields for the object.
	UpdateFields() map[c.UpdateField]interface{}
}
