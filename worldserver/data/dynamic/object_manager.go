package dynamic

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/sirupsen/logrus"
)

var (
	instance *ObjectManager
)

// ObjectManager is a structure which manages a set of game objects. Each type of object is stored
// in a separate map to make for easy serialization/deserialization.
type ObjectManager struct {
	log         *logrus.Entry
	objectsLock sync.Mutex

	// Mappings from GUID --> Object.
	Containers  map[interfaces.GUID]*Container
	GameObjects map[interfaces.GUID]*GameObject
	Items       map[interfaces.GUID]*Item
	Players     map[interfaces.GUID]*Player
	Units       map[interfaces.GUID]*Unit

	// A set of all active GUIDs, for easy existence checking.
	ActiveIDs map[interfaces.GUID]bool

	// Keep track of the next free ID for each HighGUID type. Also keep track of IDs which
	// have been freed so they can be re-used.
	NextID  map[static.HighGUID]uint32
	FreeIDs map[static.HighGUID][]uint32
}

// GetObjectManager constructs a new ObjectManager and returns it.
func GetObjectManager() *ObjectManager {
	if instance == nil {
		instance = &ObjectManager{
			log: logrus.WithField("system", "ObjectObjectManager"),

			Containers:  make(map[interfaces.GUID]*Container, 0),
			GameObjects: make(map[interfaces.GUID]*GameObject, 0),
			Items:       make(map[interfaces.GUID]*Item, 0),
			Players:     make(map[interfaces.GUID]*Player, 0),
			Units:       make(map[interfaces.GUID]*Unit, 0),

			ActiveIDs: make(map[interfaces.GUID]bool, 0),

			NextID: map[static.HighGUID]uint32{
				static.HighGUIDItem:       1,
				static.HighGUIDPlayer:     1,
				static.HighGUIDGameobject: 1,
				static.HighGUIDUnit:       1,
			},

			FreeIDs: map[static.HighGUID][]uint32{
				static.HighGUIDItem:       make([]uint32, 0),
				static.HighGUIDPlayer:     make([]uint32, 0),
				static.HighGUIDGameobject: make([]uint32, 0),
				static.HighGUIDUnit:       make([]uint32, 0),
			},
		}
	}

	return instance
}

func (m *ObjectManager) getNextID(highGUID static.HighGUID) interfaces.GUID {
	var id interfaces.GUID
	if len(m.FreeIDs[highGUID]) > 0 {
		id = interfaces.MakeGUID(m.FreeIDs[highGUID][0], highGUID)
		m.FreeIDs[highGUID] = m.FreeIDs[highGUID][1:]
	} else {
		id = interfaces.MakeGUID(m.NextID[highGUID], highGUID)
		m.NextID[highGUID]++
	}

	m.ActiveIDs[id] = true
	return id
}

// Add registers a new object with the ObjectManager, assigning it a GUID in the process.
func (m *ObjectManager) Add(object interfaces.Object) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	switch objectTyped := object.(type) {
	case *Container:
		if objectTyped.GUID() == 0 {
			objectTyped.SetGUID(m.getNextID(static.HighGUIDContainer))
		}

		m.Containers[objectTyped.GUID()] = objectTyped
	case *GameObject:
		if objectTyped.GUID() == 0 {
			objectTyped.SetGUID(m.getNextID(static.HighGUIDGameobject))
		}

		m.GameObjects[objectTyped.GUID()] = objectTyped
	case *Item:
		if objectTyped.GUID() == 0 {
			objectTyped.SetGUID(m.getNextID(static.HighGUIDItem))
		}

		m.Items[objectTyped.GUID()] = objectTyped
	case *Player:
		if objectTyped.GUID() == 0 {
			objectTyped.SetGUID(m.getNextID(static.HighGUIDPlayer))
		}

		m.Players[objectTyped.GUID()] = objectTyped
	case *Unit:
		if objectTyped.GUID() == 0 {
			objectTyped.SetGUID(m.getNextID(static.HighGUIDUnit))
		}

		m.Units[objectTyped.GUID()] = objectTyped
	default:
		panic("unknown object type")
	}
}

// Get returns a generic Object with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) Get(guid interfaces.GUID) interfaces.Object {
	switch guid.High() {
	case static.HighGUIDItem:
		if item := m.GetItem(guid); item != nil {
			return item
		}

		return m.GetContainer(guid)
	case static.HighGUIDPlayer:
		return m.GetPlayer(guid)
	case static.HighGUIDGameobject:
		return m.GetGameObject(guid)
	case static.HighGUIDUnit:
		return m.GetUnit(guid)
	}

	return nil
}

// GetContainer returns an Container with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) GetContainer(guid interfaces.GUID) *Container {
	if object, ok := m.Containers[guid]; ok {
		return object
	}

	return nil
}

// GetGameObject returns an GameObject with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) GetGameObject(guid interfaces.GUID) *GameObject {
	if object, ok := m.GameObjects[guid]; ok {
		return object
	}

	return nil
}

// GetItem returns an Item with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) GetItem(guid interfaces.GUID) *Item {
	if object, ok := m.Items[guid]; ok {
		return object
	}

	return nil
}

// GetPlayer returns an Player with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) GetPlayer(guid interfaces.GUID) *Player {
	if object, ok := m.Players[guid]; ok {
		return object
	}

	return nil
}

// GetUnit returns an Unit with the given GUID, or nil if it doesn't exist.
func (m *ObjectManager) GetUnit(guid interfaces.GUID) *Unit {
	if object, ok := m.Units[guid]; ok {
		return object
	}

	return nil
}

// Remove removes the object with the given GUID and registers the GUID as free.
func (m *ObjectManager) Remove(guid interfaces.GUID) {
	m.objectsLock.Lock()
	defer m.objectsLock.Unlock()

	delete(m.Containers, guid)
	delete(m.GameObjects, guid)
	delete(m.Items, guid)
	delete(m.Players, guid)
	delete(m.Units, guid)
	m.FreeIDs[guid.High()] = append(m.FreeIDs[guid.High()], guid.Low())
}

// Exists will return true iff an object with the given GUID exists.
func (m *ObjectManager) Exists(guid interfaces.GUID) bool {
	if _, ok := m.ActiveIDs[guid]; ok {
		return true
	}

	return false
}

// LoadFromJSON will load data to populate the Object Manager from a JSON file.
func (m *ObjectManager) LoadFromJSON(jsonFilepath string) error {
	file, err := os.OpenFile(jsonFilepath, os.O_RDONLY, 0555)
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

// SaveToJSON exports the data within the object manager to a JSON file.
func (m *ObjectManager) SaveToJSON(filepath string) error {
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
