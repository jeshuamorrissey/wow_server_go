package system

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

const (
	// MaxObjectDistance is the distance at which objects will be sent to
	// the client. Anything with a distance larger than this is considered
	// "out of range".
	// TODO(jeshua): Make this reasonable.
	MaxObjectDistance = 500
)

// MakeUpdateObjectPacketFn is a function which takes as input some
// update information and returns a ServerPacket that can be sent.
type MakeUpdateObjectPacketFn func(OutOfRangeUpdate, []ObjectUpdate) ServerPacket

// MakeAttackerStateUpdatePackerFn is a function which takes as input some
// combat information and returns a ServerPacket that can be sent.
type MakeAttackerStateUpdatePackerFn func(interfaces.GUID, interfaces.GUID, interfaces.AttackInfo) ServerPacket

type updateCache struct {
	UpdateFields   interfaces.UpdateFieldsMap
	MovementUpdate []byte
}

// OutOfRangeUpdate represents a list of GUIDs which are no longer in range.
type OutOfRangeUpdate struct {
	GUIDS []interfaces.GUID
}

// ObjectUpdate represents an update to an individual object.
type ObjectUpdate struct {
	GUID        interfaces.GUID
	UpdateType  static.UpdateType
	UpdateFlags static.UpdateFlags
	IsSelf      bool
	TypeID      static.TypeID

	MovementUpdate  []byte
	NumUpdateFields int
	UpdateFields    interfaces.UpdateFieldsMap

	VictimGUID interfaces.GUID
	WorldTime  uint32
}

type loginData struct {
	Session     *Session
	UpdateCache map[interfaces.GUID]*updateCache
}

// Updater manages sending object updates to sessions based on when objects have been changed.
type Updater struct {
	log *logrus.Entry
	om  *dynamic.ObjectManager

	sessionsLock sync.Mutex
	sessions     map[interfaces.GUID]*loginData // logged in character --> Session

	toUpdateLock sync.Mutex
	toUpdate     []interfaces.GUID

	makeUpdateObjectPacketFn        MakeUpdateObjectPacketFn
	makeAttackerStateUpdatePackerFn MakeAttackerStateUpdatePackerFn
}

// NewUpdater makes a new updater object.
func NewUpdater(log *logrus.Entry, om *dynamic.ObjectManager, makeUpdateObjectPacketFn MakeUpdateObjectPacketFn, makeAttackerStateUpdatePacker MakeAttackerStateUpdatePackerFn) *Updater {
	u := &Updater{
		log: log.WithFields(logrus.Fields{
			"system": "Updater",
		}),
		om:                              om,
		sessions:                        make(map[interfaces.GUID]*loginData),
		toUpdate:                        make([]interfaces.GUID, 0),
		makeUpdateObjectPacketFn:        makeUpdateObjectPacketFn,
		makeAttackerStateUpdatePackerFn: makeAttackerStateUpdatePacker,
	}

	return u
}

// Login registers the given player as logged in for the given session.
func (u *Updater) Login(playerGUID interfaces.GUID, session *Session) error {
	u.sessionsLock.Lock()
	defer u.sessionsLock.Unlock()

	if currPlayerGUID, ok := u.sessions[playerGUID]; ok {
		u.log.WithFields(logrus.Fields{
			"current_player": currPlayerGUID,
			"new_player":     playerGUID,
		}).Warningf("Tried to log in with multiple characters")
		return fmt.Errorf("player with GUID %v is already logged in", playerGUID)
	}

	u.sessions[playerGUID] = &loginData{
		Session:     session,
		UpdateCache: make(map[interfaces.GUID]*updateCache),
	}

	u.log.Tracef("Registered GUID=%v", playerGUID)

	// Mark the player as logged in.
	u.om.GetPlayer(playerGUID).IsLoggedIn = true

	u.toUpdateLock.Lock()
	defer u.toUpdateLock.Unlock()
	u.toUpdate = append(u.toUpdate, playerGUID)

	return nil
}

// Logout deregisters the player.
func (u *Updater) Logout(playerGUID interfaces.GUID) error {
	u.sessionsLock.Lock()
	defer u.sessionsLock.Unlock()

	if _, ok := u.sessions[playerGUID]; !ok {
		u.log.Warningf("player with GUID %v is not logged in, but we got a logout request", playerGUID)
	}

	// Mark the player as logged in.
	u.om.GetPlayer(playerGUID).IsLoggedIn = false

	delete(u.sessions, playerGUID)
	return nil
}

func (u *Updater) makeUpdates(
	loginData *loginData,
	player interfaces.Object,
	objectToUpdate interfaces.Object,
	outOfRangeUpdate *OutOfRangeUpdate,
	objectUpdates *[]ObjectUpdate) {
	if objectToUpdate.GetLocation() != nil {
		var movementUpdate []byte = nil
		switch objectToUpdateTyped := objectToUpdate.(type) {
		case *dynamic.Unit:
		case *dynamic.Player:
			movementUpdate = objectToUpdateTyped.MovementUpdate()
		}

		if objectToUpdate.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance {
			update := ObjectUpdate{
				GUID:            objectToUpdate.GUID(),
				UpdateFlags:     dynamic.UpdateFlags(objectToUpdate),
				IsSelf:          objectToUpdate.GUID() == player.GUID(),
				TypeID:          dynamic.TypeID(objectToUpdate),
				MovementUpdate:  movementUpdate,
				NumUpdateFields: dynamic.NumUpdateFields(objectToUpdate),
				UpdateFields:    objectToUpdate.UpdateFields(),
				VictimGUID:      0, // TODO: what is this?
				WorldTime:       0, // TODO: what is this?
			}

			// If we have seen the object, update it.
			lastUpdateFields, ok := loginData.UpdateCache[objectToUpdate.GUID()]
			if ok {
				update.UpdateType = static.UpdateTypeValues

				for k, v := range lastUpdateFields.UpdateFields {
					if update.UpdateFields[k] == v {
						delete(update.UpdateFields, k)
					}
				}
			} else {
				if update.MovementUpdate != nil {
					update.UpdateType = static.UpdateTypeCreateObject2
				} else {
					update.UpdateType = static.UpdateTypeCreateObject
				}

				loginData.UpdateCache[objectToUpdate.GUID()] = &updateCache{
					UpdateFields:   update.UpdateFields,
					MovementUpdate: update.MovementUpdate,
				}
			}

			*objectUpdates = append(*objectUpdates, update)
		} else {
			// If the object was in range, but no longer is, delete it.
			if _, ok := loginData.UpdateCache[objectToUpdate.GUID()]; ok {
				delete(loginData.UpdateCache, objectToUpdate.GUID())
				outOfRangeUpdate.GUIDS = append(outOfRangeUpdate.GUIDS, objectToUpdate.GUID())
			}
		}
	}
}

func (u *Updater) updatePlayer(guid interfaces.GUID, loginData *loginData) {
	outOfRangeUpdate := OutOfRangeUpdate{
		GUIDS: make([]interfaces.GUID, 0),
	}

	objectUpdates := make([]ObjectUpdate, 0)

	// Find all objects that are close to this player and make sure they
	// have been updated.
	playerObj := u.om.Get(guid)
	for _, player := range u.om.Players {
		if !player.IsLoggedIn {
			continue
		}

		for _, itemGUID := range player.Inventory {
			if u.om.Exists(itemGUID) {
				u.makeUpdates(loginData, playerObj, u.om.GetItem(itemGUID), &outOfRangeUpdate, &objectUpdates)
			}
		}

		for _, itemGUID := range player.Equipment {
			if u.om.Exists(itemGUID) {
				u.makeUpdates(loginData, playerObj, u.om.GetItem(itemGUID), &outOfRangeUpdate, &objectUpdates)
			}
		}

		for _, bagGUID := range player.Bags {
			if u.om.Exists(bagGUID) {
				bag := u.om.GetContainer(bagGUID)
				u.makeUpdates(loginData, playerObj, bag, &outOfRangeUpdate, &objectUpdates)
				for _, itemGUID := range bag.Items {
					if u.om.Exists(itemGUID) {
						u.makeUpdates(loginData, playerObj, u.om.GetItem(itemGUID), &outOfRangeUpdate, &objectUpdates)
					}
				}
			}
		}

		u.makeUpdates(loginData, playerObj, player, &outOfRangeUpdate, &objectUpdates)
	}

	pkt := u.makeUpdateObjectPacketFn(outOfRangeUpdate, objectUpdates)
	loginData.Session.Send(pkt)
}

func (u *Updater) updateOtherPlayers(guid interfaces.GUID) {
	for playerGUID, loginData := range u.sessions {
		if playerGUID == guid {
			continue
		}

		outOfRangeUpdate := OutOfRangeUpdate{
			GUIDS: make([]interfaces.GUID, 0),
		}

		objectUpdates := make([]ObjectUpdate, 0)

		player := u.om.Get(playerGUID)
		obj := u.om.Get(guid)
		u.makeUpdates(loginData, player, obj, &outOfRangeUpdate, &objectUpdates)
		loginData.Session.Send(u.makeUpdateObjectPacketFn(outOfRangeUpdate, objectUpdates))
	}
}

// Run starts the updater, which will constantly scan for object updates.
// Should be run as a goroutine.
func (u *Updater) Run() {
	for {
		time.Sleep(time.Millisecond * 30)
		u.toUpdateLock.Lock()

		// There are some object to update!
		for _, guid := range u.toUpdate {
			// First, check to see if this is a player. If it is, then we
			// have to update the _player_ with all objects around them.
			if loginData, ok := u.sessions[guid]; ok {
				u.updatePlayer(guid, loginData)
			}

			// Second, we need to notify other players of this update.
			u.updateOtherPlayers(guid)
		}

		u.toUpdate = make([]interfaces.GUID, 0)
		u.toUpdateLock.Unlock()
	}
}

func (u *Updater) SendCombatUpdate(attacker interfaces.Unit, target interfaces.Unit, attackInfo interfaces.AttackInfo) {
	// Find all players in range of either the attacker or target.
	for playerGUID, loginData := range u.sessions {
		player := u.om.Get(playerGUID)
		if player == nil {
			continue
		}

		if attacker.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance || target.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance {
			loginData.Session.Send(u.makeAttackerStateUpdatePackerFn(attacker.GUID(), target.GUID(), attackInfo))
		}
	}
}
