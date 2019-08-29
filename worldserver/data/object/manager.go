package object

import (
	"fmt"
	"sync"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/sirupsen/logrus"
)

// Manager is a container which managed a list of objects. It can
// be serialized/deserialized from JSON.
type Manager struct {
	log *logrus.Entry

	objectsLock sync.Mutex
	objects     map[GUID]Object

	changeWaitersLock sync.Mutex
	changeWaiters     []func(GUID)

	nextID  map[c.HighGUID]uint32
	freeIDs map[c.HighGUID][]uint32
}

// NewManager constructs a new object manager and returns it.
func NewManager(log *logrus.Entry) *Manager {
	return &Manager{
		log:           log,
		objects:       make(map[GUID]Object, 0),
		changeWaiters: make([]func(GUID), 0),
		nextID: map[c.HighGUID]uint32{
			c.HighGUIDItem:          0,
			c.HighGUIDPlayer:        0,
			c.HighGUIDGameobject:    0,
			c.HighGUIDTransport:     0,
			c.HighGUIDUnit:          0,
			c.HighGUIDPet:           0,
			c.HighGUIDDynamicObject: 0,
			c.HighGUIDCorpse:        0,
			c.HighGUIDMoTransport:   0,
		},

		freeIDs: map[c.HighGUID][]uint32{
			c.HighGUIDItem:          make([]uint32, 0),
			c.HighGUIDPlayer:        make([]uint32, 0),
			c.HighGUIDGameobject:    make([]uint32, 0),
			c.HighGUIDTransport:     make([]uint32, 0),
			c.HighGUIDUnit:          make([]uint32, 0),
			c.HighGUIDPet:           make([]uint32, 0),
			c.HighGUIDDynamicObject: make([]uint32, 0),
			c.HighGUIDCorpse:        make([]uint32, 0),
			c.HighGUIDMoTransport:   make([]uint32, 0),
		},
	}
}

// Add adds a new Object to the object manager. A new GUID will be
// assigned to the object.
func (m *Manager) Add(obj Object) error {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	// Get the next available ID. Re-use one if available.
	var id GUID
	highGUID := HighGUID(obj)
	if len(m.freeIDs[highGUID]) > 0 {
		id = MakeGUID(m.freeIDs[highGUID][0], highGUID)
		m.freeIDs[highGUID] = m.freeIDs[highGUID][1:]
	} else {
		id = MakeGUID(m.nextID[highGUID], highGUID)
		m.nextID[highGUID]++
	}

	// Add the object to the manager and set its ID.
	obj.SetGUID(id)
	obj.SetManager(m)
	m.objects[id] = obj

	return nil
}

// Remove deletes the object with the given GUID from the system. Will
// return an error if the GUID does not exist.
func (m *Manager) Remove(guid GUID) error {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	if _, ok := m.objects[guid]; !ok {
		return fmt.Errorf("object %v does not exist", guid)
	}

	delete(m.objects, guid)
	m.freeIDs[guid.High()] = append(m.freeIDs[guid.High()], guid.Low())

	return nil
}

// Get retreives the given object from the manager. Will return an error
// if the given GUID is not registered in the manager.
func (m *Manager) Get(guid GUID) Object {
	return m.objects[guid]
}

// Exists checks for the existance of a given object.
func (m *Manager) Exists(guid GUID) bool {
	_, ok := m.objects[guid]
	return ok
}

// Update marks the object as updated (which will trigger all OnChange
// events).
func (m *Manager) Update(guid GUID) {
	for _, onChange := range m.changeWaiters {
		onChange(guid)
	}
}

// AwaitChange will wait for any change to any object in the manager and
// call the appropriate function with the changed object.
func (m *Manager) AwaitChange(onChange func(GUID)) {
	m.changeWaitersLock.Lock()
	defer m.changeWaitersLock.Unlock()

	m.changeWaiters = append(m.changeWaiters, onChange)
}

// Objects returns a reference to the full object map.
// TODO(jeshua): Make this more efficient (e.g. only return objects within
// a certain distance, ...).
func (m *Manager) Objects() map[GUID]Object {
	return m.objects
}
