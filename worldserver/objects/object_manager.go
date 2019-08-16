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
	objects     map[GUID]GameObject

	updatesLock               sync.Mutex
	updatesAvailableMutex     sync.Mutex
	updatesAvailableCondition *sync.Cond
	updates                   []GUID

	playersLock sync.Mutex
	players     map[GUID]func([]GameObject)
}

// Create will make a new object in the object manager.
func (om *ObjectManager) Create(obj GameObject) GameObject {
	om.objectsLock.Lock()
	obj.SetGUID(nextGUID)
	om.objects[obj.GUID()] = obj
	nextGUID++
	om.objectsLock.Unlock()
	return obj
}

// NewObjectManager initialzies an ObjectManager object.
func NewObjectManager() *ObjectManager {
	om := new(ObjectManager)
	om.objects = make(map[GUID]GameObject)
	om.players = make(map[GUID]func([]GameObject))
	om.updates = make([]GUID, 0)
	om.updatesAvailableCondition = sync.NewCond(&om.updatesAvailableMutex)
	return om
}

// Update marks the given GUID as updated. This should be done to ensure that all
// clients that need updates for this object receive them.
func (om *ObjectManager) Update(guid GUID) {
	om.updatesLock.Lock()
	om.updates = append(om.updates, guid)
	om.updatesAvailableCondition.Signal()
	om.updatesLock.Unlock()
}

// Get will fetch the game object with the given GUID and return it. `nil` will be
// returned if no object with that GUID exists.
func (om *ObjectManager) Get(guid GUID) GameObject {
	if obj, ok := om.objects[guid]; ok {
		return obj
	}

	return nil
}

// Exists will return true iff and item with the given GUID exists.
func (om *ObjectManager) Exists(guid GUID) bool {
	_, ok := om.objects[guid]
	return ok
}

// Register notes that the given player is expecting to receive updates.
func (om *ObjectManager) Register(player GameObject, updateFunc func([]GameObject)) {
	om.playersLock.Lock()
	om.players[player.GUID()] = updateFunc
	om.playersLock.Unlock()
}

// Run takes control of the thread and watches for object updates and distributes
// them to all registered players (if it makes sense to do so).
func (om *ObjectManager) Run() {
	for {
		om.updatesAvailableMutex.Lock()
		om.updatesAvailableCondition.Wait()
		om.updatesAvailableMutex.Unlock()

		om.updatesLock.Lock()
		updates := om.updates[:]
		om.updates = make([]GUID, 0)
		om.updatesLock.Unlock()

		for playerGUID, updateFunc := range om.players {
			playerLocation := om.Get(playerGUID).GetLocation()
			updatesToSend := make([]GameObject, len(updates))

			for _, updatedGUID := range updates {
				updatedObj := om.Get(updatedGUID)

				// If the updated object is a player/unit, check to see if
				// the locations are close enough.
				updatedObjLocation := updatedObj.GetLocation()

				if updatedObjLocation != nil && playerLocation != nil {
					distance := updatedObjLocation.Distance(playerLocation)
					if distance < ObjectMaxUpdateDistance {
						updatesToSend = append(updatesToSend, updatedObj)
					}
				}
			}

			updateFunc(updatesToSend)
		}
	}
}
