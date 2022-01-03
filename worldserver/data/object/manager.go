package object

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/sirupsen/logrus"
)

// Manager is a container which managed a list of Objects. It can
// be serialized/deserialized from JSON.
type Manager struct {
	log *logrus.Entry

	objectsLock sync.Mutex

	Containers  map[GUID]*Container
	GameObjects map[GUID]*GameObject
	Items       map[GUID]*Item
	Players     map[GUID]*Player
	Units       map[GUID]*Unit

	// Transports     map[GUID]*Transport

	// Pets     map[GUID]*Pet
	// DynamicObjects     map[GUID]*DynamicObjects
	// Corpses     map[GUID]*Corpse

	changeWaitersLock sync.Mutex
	changeWaiters     []func(GUID)

	NextID  map[c.HighGUID]uint32
	FreeIDs map[c.HighGUID][]uint32
}

// NewManager constructs a new object manager and returns it.
func NewManager(log *logrus.Entry) *Manager {
	return &Manager{
		log:           log,
		Containers:    make(map[GUID]*Container, 0),
		GameObjects:   make(map[GUID]*GameObject, 0),
		Items:         make(map[GUID]*Item, 0),
		Players:       make(map[GUID]*Player, 0),
		Units:         make(map[GUID]*Unit, 0),
		changeWaiters: make([]func(GUID), 0),
		NextID: map[c.HighGUID]uint32{
			c.HighGUIDItem:          1,
			c.HighGUIDPlayer:        1,
			c.HighGUIDGameobject:    1,
			c.HighGUIDTransport:     1,
			c.HighGUIDUnit:          1,
			c.HighGUIDPet:           1,
			c.HighGUIDDynamicObject: 1,
			c.HighGUIDCorpse:        1,
			c.HighGUIDMoTransport:   1,
		},

		FreeIDs: map[c.HighGUID][]uint32{
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

// NewManagerFrom creates a new Manager and then loads data from a file.
func NewManagerFrom(log *logrus.Entry, filepath string) *Manager {
	manager := NewManager(log)
	manager.LoadFrom(filepath)
	return manager
}

func (m *Manager) getNextID(highGUID c.HighGUID) GUID {
	var id GUID
	if len(m.FreeIDs[highGUID]) > 0 {
		id = MakeGUID(m.FreeIDs[highGUID][0], highGUID)
		m.FreeIDs[highGUID] = m.FreeIDs[highGUID][1:]
	} else {
		id = MakeGUID(m.NextID[highGUID], highGUID)
		m.NextID[highGUID]++
	}

	return id
}

func (m *Manager) AddPlayer(player *Player) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	player.SetGUID(m.getNextID(c.HighGUIDPlayer))
	player.SetManager(m)
	m.Players[player.GUID()] = player
}

func (m *Manager) AddUnit(unit *Unit) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	unit.SetGUID(m.getNextID(c.HighGUIDUnit))
	unit.SetManager(m)
	m.Units[unit.GUID()] = unit
}

func (m *Manager) AddContainer(container *Container) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	container.SetGUID(m.getNextID(c.HighGUIDContainer))
	container.SetManager(m)
	m.Containers[container.GUID()] = container
}

func (m *Manager) AddItem(item *Item) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	item.SetGUID(m.getNextID(c.HighGUIDItem))
	item.SetManager(m)
	m.Items[item.GUID()] = item
}

func (m *Manager) AddGameObject(gameObject *GameObject) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	gameObject.SetGUID(m.getNextID(c.HighGUIDGameobject))
	gameObject.SetManager(m)
	m.GameObjects[gameObject.GUID()] = gameObject
}

// Remove deletes the object with the given GUID from the system. Will
// return an error if the GUID does not exist.
func (m *Manager) Remove(guid GUID) error {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	if m.Get(guid) == nil {
		return fmt.Errorf("object %v does not exist", guid)
	}

	switch guid.High() {
	case c.HighGUIDItem: // same GUID as containers
		delete(m.Items, guid)
		delete(m.Containers, guid)
	case c.HighGUIDPlayer:
		delete(m.Players, guid)
	case c.HighGUIDUnit:
		delete(m.Units, guid)
	default:
	}

	m.FreeIDs[guid.High()] = append(m.FreeIDs[guid.High()], guid.Low())

	return nil
}

// Get retreives the given object from the manager. Will return an error
// if the given GUID is not registered in the manager.
func (m *Manager) Get(guid GUID) Object {
	switch guid.High() {
	case c.HighGUIDItem: // same GUID as containers
		if item, ok := m.Items[guid]; ok {
			return item
		} else if container, ok := m.Containers[guid]; ok {
			return container
		}
	case c.HighGUIDUnit:
		if unit, ok := m.Units[guid]; ok {
			return unit
		}
	case c.HighGUIDPlayer:
		if player, ok := m.Players[guid]; ok {
			return player
		}
	default:
	}

	return nil
}

// Exists checks for the existance of a given object.
func (m *Manager) Exists(guid GUID) bool {
	return m.Get(guid) != nil
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

func (m *Manager) LoadFrom(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0555)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) SaveTo(filepath string) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0555)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
