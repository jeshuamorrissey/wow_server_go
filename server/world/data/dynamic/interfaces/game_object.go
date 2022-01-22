package interfaces

// Object represents a generic object within the game. All objects
// should implement this interface.
type Object interface {
	// GUID should return the full GUID of the object.
	GUID() GUID

	// SetGUID should set the GUID of the object.
	SetGUID(GUID)

	// GetLocation should return the location of the object.
	GetLocation() *Location

	// UpdateFields should return the update fields for the object.
	UpdateFields() UpdateFieldsMap

	StartUpdateLoop()
	SendUpdates([]interface{})
}
