package system

// const (
// 	// ObjectMaxUpdateDistance is the maximum distance before two objects stop
// 	// being updated of eachother's changes.
// 	// TODO(jeshua): Make this a config value.
// 	ObjectMaxUpdateDistance = 3000
// )

// var (
// 	nextGUID = 0
// )

// // NewCharacter makes a new character with some basic information.
// func NewCharacter(
// 	om *ObjectManager,
// 	name string,
// 	account *Account, realm *Realm,
// 	class c.Class, race c.Race, gender c.Gender,
// 	skinColor, face, hairStyle, hairColor, feature uint8) *Character {
// 	startingEquipment, startingItems := data.GetStartingItems(class, race)

// 	equipment := make(map[c.EquipmentSlot]*objects.Item)
// 	for slot, item := range startingEquipment {
// 		equipment[slot] = om.Create(&objects.Item{
// 			BaseGameObject: objects.BaseGameObject{
// 				Entry: item.Entry,
// 			},
// 		}).(*objects.Item)
// 	}

// 	inventory := make(map[int]*objects.Item)
// 	for i, item := range startingItems {
// 		inventory[i] = om.Create(&objects.Item{
// 			BaseGameObject: objects.BaseGameObject{
// 				Entry: item.Entry,
// 			},
// 		}).(*objects.Item)
// 	}

// 	startingLocation := data.GetStartingLocation(class, race)

// 	return &Character{
// 		Name: name,
// 		GUID: om.Create(&objects.Player{
// 			Unit: objects.Unit{
// 				BaseGameObject: objects.BaseGameObject{
// 					Entry:  0,
// 					ScaleX: data.GetPlayerScale(race, gender),
// 				},

// 				Location: objects.Location{
// 					X: startingLocation.X,
// 					Y: startingLocation.Y,
// 					Z: startingLocation.Z,
// 					O: startingLocation.O,
// 				},

// 				Level:  1,
// 				Race:   race,
// 				Class:  class,
// 				Gender: gender,
// 			},

// 			SkinColor: skinColor,
// 			Face:      face,
// 			HairStyle: hairStyle,
// 			HairColor: hairColor,
// 			Feature:   feature,

// 			ZoneID: startingLocation.Zone,
// 			MapID:  startingLocation.Map,

// 			Equipment: equipment,
// 			Inventory: inventory,
// 		}).GUID(),
// 		AccountID: account.ID,
// 		RealmID:   realm.ID,
// 	}
// }

// // ObjectManager manages in-game objects, which are accessed using their GUID.
// type ObjectManager struct {
// 	// Mapping of object GUID --> in-game object.
// 	data map[o.GUID]o.GameObject

// 	// A list of currently running WorldServer sessions.
// 	sessions []Session

// 	objectsLock sync.Mutex
// 	objects     map[objects.]objects.GameObject

// 	updatesLock               sync.Mutex
// 	updatesAvailableMutex     sync.Mutex
// 	updatesAvailableCondition *sync.Cond
// 	updates                   []objects.

// 	playersLock       sync.Mutex
// 	players           map[objects.]func(map[c.UpdateType][]objects.GameObject)
// 	playerUpdateCache map[objects.]map[objects.]map[c.UpdateField]interface{} // map from Player --> Object --> UpdateFields.
// }

// // Create will make a new object in the object manager.
// func (om *ObjectManager) Create(obj GameObject) GameObject {
// 	om.objectsLock.Lock()
// 	obj.SetGUID(nextGUID)
// 	om.objects[obj.GUID()] = obj
// 	nextGUID++
// 	om.objectsLock.Unlock()
// 	return obj
// }

// // NewObjectManager initialzies an ObjectManager object.
// func NewObjectManager() *ObjectManager {
// 	om := new(ObjectManager)
// 	om.objects = make(map[GUID]GameObject)
// 	om.players = make(map[GUID]func(map[c.UpdateType][]GameObject))
// 	om.playerUpdateCache = make(map[GUID]map[GUID]map[c.UpdateField]interface{})
// 	om.updates = make([]GUID, 0)
// 	om.updatesAvailableCondition = sync.NewCond(&om.updatesAvailableMutex)
// 	return om
// }

// // Update marks the given GUID as updated. This should be done to ensure that all
// // clients that need updates for this object receive them.
// func (om *ObjectManager) Update(guid GUID) {
// 	om.updatesLock.Lock()
// 	om.updates = append(om.updates, guid)
// 	om.updatesAvailableCondition.Signal()
// 	om.updatesLock.Unlock()
// }

// // Get will fetch the game object with the given GUID and return it. `nil` will be
// // returned if no object with that GUID exists.
// func (om *ObjectManager) Get(guid GUID) GameObject {
// 	if obj, ok := om.objects[guid]; ok {
// 		return obj
// 	}

// 	return nil
// }

// // Exists will return true iff and item with the given GUID exists.
// func (om *ObjectManager) Exists(guid GUID) bool {
// 	_, ok := om.objects[guid]
// 	return ok
// }

// // Register notes that the given player is expecting to receive updates.
// func (om *ObjectManager) Register(player GameObject, updateFunc func(map[c.UpdateType][]GameObject)) {
// 	om.playersLock.Lock()
// 	om.players[player.GUID()] = updateFunc
// 	om.playerUpdateCache[player.GUID()] = make(map[GUID]map[c.UpdateField]interface{})
// 	om.playersLock.Unlock()
// }

// // Run takes control of the thread and watches for object updates and distributes
// // them to all registered players (if it makes sense to do so).
// func (om *ObjectManager) Run() {
// 	for {
// 		om.updatesAvailableMutex.Lock()
// 		om.updatesAvailableCondition.Wait()
// 		om.updatesAvailableMutex.Unlock()

// 		om.updatesLock.Lock()
// 		updates := om.updates[:]
// 		om.updates = make([]GUID, 0)
// 		om.updatesLock.Unlock()

// 		for playerGUID, updateFunc := range om.players {
// 			playerLocation := om.Get(playerGUID).GetLocation()
// 			updatesToSend := map[c.UpdateType][]GameObject{
// 				c.UpdateTypeCreateObject: make([]GameObject, 0),
// 			}
// 			updateCache := om.playerUpdateCache[playerGUID]

// 			for _, updatedGUID := range updates {
// 				updatedObj := om.Get(updatedGUID)

// 				// If the updated object is a player/unit, check to see if
// 				// the locations are close enough.
// 				updatedObjLocation := updatedObj.GetLocation()

// 				if updatedObjLocation != nil && playerLocation != nil {
// 					distance := updatedObjLocation.Distance(playerLocation)
// 					if distance < ObjectMaxUpdateDistance {
// 						// This is an object we want to update. First, check the cache to see
// 						// if it exists. If it doesn't, then we are creating.
// 						if _, ok := updateCache[updatedGUID]; !ok {
// 							updatesToSend[c.UpdateTypeCreateObject] = append(updatesToSend[c.UpdateTypeCreateObject], updatedObj)
// 						} else {
// 							fmt.Printf("WARNING: Don't know how to do regular object updates!")
// 						}
// 					}
// 				}
// 			}

// 			updateFunc(updatesToSend)
// 		}
// 	}
// }
