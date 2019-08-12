package objects

import (
	"sync"
)

const (
	// ObjectMaxUpdateDistance is the maximum distance before two objects stop
	// being updated of eachother's changes.
	// TODO(jeshua): Make this a config value.
	ObjectMaxUpdateDistance = 3000
)

var (
	nextGUID = 0
)

// ObjectManager manages in-game objects, which are accessed using their GUID.
type ObjectManager struct {
	objectsLock sync.Mutex
	Objects     map[GUID]GameObject

	updatesLock sync.Mutex
	updates     []GUID
	players     map[GUID]func(GameObject)
}

// Create will make a new object in the object manager.
func (om *ObjectManager) Create(obj GameObject) GameObject {
	om.objectsLock.Lock()
	obj.SetGUID(nextGUID)
	om.Objects[obj.GUID()] = obj
	nextGUID++
	om.objectsLock.Unlock()
	return obj
}

// NewObjectManager initialzies an ObjectManager object.
func NewObjectManager() *ObjectManager {
	om := new(ObjectManager)
	om.Objects = make(map[GUID]GameObject)
	om.players = make(map[GUID]func(GameObject))
	om.updates = make([]GUID, 0)
	return om
}

// Update marks the given GUID as updated. This should be done to ensure that all
// clients that need updates for this object receive them.
func (om *ObjectManager) Update(guid GUID) {
	om.updatesLock.Lock()
	om.updates = append(om.updates, guid)
	om.updatesLock.Unlock()
}

// Register notes that the given player is expecting to receive updates.
func (om *ObjectManager) Register(player GameObject, updateFunc func(GameObject)) {
	om.players[player.GUID()] = updateFunc
}

// Run takes control of the thread and watches for object updates and distributes
// them to all registered players (if it makes sense to do so).
func (om *ObjectManager) Run() {
	for {
		om.updatesLock.Lock()
		updates := om.updates[:]
		om.updates = make([]GUID, 0)
		om.updatesLock.Unlock()

		// For each update that has happened, go through each player and see if
		// they should be notified of an update to that object.
		for _, updatedGUID := range updates {
			updatedObj := om.Objects[updatedGUID]

			for playerGUID, updateFunc := range om.players {
				// If the updated object is a player/unit, check to see if
				// the locations are close enough.
				updatedObjLocation := updatedObj.GetLocation()
				playerLocation := om.Objects[playerGUID].GetLocation()

				if updatedObjLocation != nil && playerLocation != nil {
					distance := updatedObjLocation.Distance(playerLocation)
					if distance < ObjectMaxUpdateDistance {
						updateFunc(updatedObj)
					}
				}
			}
		}
	}
}
