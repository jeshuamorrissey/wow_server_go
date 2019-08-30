package system

import (
	"fmt"
	"sync"
	"time"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/sirupsen/logrus"
)

const (
	// MaxObjectDistance is the distance at which objects will be sent to
	// the client. Anything with a distance larger than this is considered
	// "out of range".
	// TODO(jeshua): Make this reasonable.
	MaxObjectDistance = 1000
)

// MakeUpdateObjectPacketFn is a function which takes as input some
// update information and returns a ServerPacket that can be sent.
type MakeUpdateObjectPacketFn func(OutOfRangeUpdate, []ObjectUpdate) ServerPacket

type updateCache struct {
	UpdateFields   object.UpdateFieldsMap
	MovementUpdate []byte
}

// OutOfRangeUpdate represents a list of GUIDs which are no longer in range.
type OutOfRangeUpdate struct {
	GUIDS []object.GUID
}

// ObjectUpdate represents an update to an individual object.
type ObjectUpdate struct {
	GUID        object.GUID
	UpdateType  c.UpdateType
	UpdateFlags c.UpdateFlags
	IsSelf      bool
	TypeID      c.TypeID

	MovementUpdate  []byte
	NumUpdateFields int
	UpdateFields    object.UpdateFieldsMap

	VictimGUID object.GUID
	WorldTime  uint32
}

type loginData struct {
	Session     *Session
	UpdateCache map[object.GUID]*updateCache
}

// Updater manages sending object updates to sessions based on when objects have been changed.
type Updater struct {
	log *logrus.Entry
	om  *object.Manager

	sessionsLock sync.Mutex
	sessions     map[object.GUID]*loginData // logged in character --> Session

	toUpdateLock sync.Mutex
	toUpdate     []object.GUID

	makeUpdateObjectPacketFn MakeUpdateObjectPacketFn
}

// NewUpdater makes a new updater object.
func NewUpdater(log *logrus.Entry, om *object.Manager, makeUpdateObjectPacketFn MakeUpdateObjectPacketFn) *Updater {
	u := &Updater{
		log: log.WithFields(logrus.Fields{
			"system": "Updater",
		}),
		om:                       om,
		sessions:                 make(map[object.GUID]*loginData),
		toUpdate:                 make([]object.GUID, 0),
		makeUpdateObjectPacketFn: makeUpdateObjectPacketFn,
	}

	return u
}

// Login registers the given player as logged in for the given session.
func (u *Updater) Login(playerGUID object.GUID, session *Session) error {
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
		UpdateCache: make(map[object.GUID]*updateCache),
	}

	u.log.Tracef("Registered GUID=%v", playerGUID)

	u.toUpdateLock.Lock()
	defer u.toUpdateLock.Unlock()
	u.toUpdate = append(u.toUpdate, playerGUID)

	return nil
}

// Logout deregisters the player.
func (u *Updater) Logout(playerGUID object.GUID) error {
	u.sessionsLock.Lock()
	defer u.sessionsLock.Unlock()

	if _, ok := u.sessions[playerGUID]; !ok {
		u.log.Warningf("player with GUID %v is not logged in, but we got a logout request", playerGUID)
	}

	delete(u.sessions, playerGUID)
	return nil
}

func (u *Updater) makeUpdates(
	loginData *loginData,
	playerObj, obj object.Object,
	outOfRangeUpdate *OutOfRangeUpdate,
	objectUpdates *[]ObjectUpdate) {
	if obj.Location() != nil {
		if obj.Location().Distance(playerObj.Location()) < MaxObjectDistance {
			update := ObjectUpdate{
				GUID:            obj.GUID(),
				UpdateFlags:     object.UpdateFlags(obj),
				IsSelf:          obj.GUID() == playerObj.GUID(),
				TypeID:          object.TypeID(obj),
				MovementUpdate:  obj.MovementUpdate(),
				NumUpdateFields: object.NumUpdateFields(obj),
				UpdateFields:    obj.UpdateFields(),
				VictimGUID:      0, // TODO: what is this?
				WorldTime:       0, // TODO: what is this?
			}

			// If we have seen the object, update it.
			lastUpdateFields, ok := loginData.UpdateCache[obj.GUID()]
			if ok {
				update.UpdateType = c.UpdateTypeValues

				for k, v := range lastUpdateFields.UpdateFields {
					if update.UpdateFields[k] == v {
						delete(update.UpdateFields, k)
					}
				}
			} else {
				if update.MovementUpdate != nil {
					update.UpdateType = c.UpdateTypeCreateObject2
				} else {
					update.UpdateType = c.UpdateTypeCreateObject
				}

				loginData.UpdateCache[obj.GUID()] = &updateCache{
					UpdateFields:   update.UpdateFields,
					MovementUpdate: update.MovementUpdate,
				}
			}

			*objectUpdates = append(*objectUpdates, update)
		} else {
			// If the object was in range, but no longer is, delete it.
			if _, ok := loginData.UpdateCache[obj.GUID()]; ok {
				delete(loginData.UpdateCache, obj.GUID())
				outOfRangeUpdate.GUIDS = append(outOfRangeUpdate.GUIDS, obj.GUID())
			}
		}
	}
}

func (u *Updater) updatePlayer(guid object.GUID, loginData *loginData) {
	outOfRangeUpdate := OutOfRangeUpdate{
		GUIDS: make([]object.GUID, 0),
	}

	objectUpdates := make([]ObjectUpdate, 0)

	// Find all objects that are close to this player and make sure they
	// have been updated.
	playerObj := u.om.Get(guid)
	for _, obj := range u.om.Objects() {
		u.makeUpdates(loginData, playerObj, obj, &outOfRangeUpdate, &objectUpdates)
	}

	pkt := u.makeUpdateObjectPacketFn(outOfRangeUpdate, objectUpdates)
	loginData.Session.Send(pkt)
}

func (u *Updater) updateOtherPlayers(guid object.GUID) {
	for playerGUID, loginData := range u.sessions {
		if playerGUID == guid {
			continue
		}

		outOfRangeUpdate := OutOfRangeUpdate{
			GUIDS: make([]object.GUID, 0),
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
	// Register an OnChange event.
	u.om.AwaitChange(func(guid object.GUID) {
		u.toUpdateLock.Lock()
		defer u.toUpdateLock.Unlock()
		u.toUpdate = append(u.toUpdate, guid)
	})

	for {
		time.Sleep(time.Second * 1)
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

		u.toUpdate = make([]object.GUID, 0)
		u.toUpdateLock.Unlock()
	}
}
