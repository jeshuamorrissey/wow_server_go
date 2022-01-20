package system

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
)

const (
	// MaxObjectDistance is the distance at which objects will be sent to
	// the client. Anything with a distance larger than this is considered
	// "out of range".
	// TODO(jeshua): Make this reasonable.
	MaxObjectDistance = 500
)

type updateCache struct {
	UpdateFields   interfaces.UpdateFieldsMap
	MovementUpdate []byte
}

type loginData struct {
	Session     *Session
	UpdateCache map[interfaces.GUID]*updateCache
}

func (ld *loginData) updateCache(obj interfaces.Object, update *packet.ObjectUpdate) {
	if _, ok := ld.UpdateCache[obj.GUID()]; !ok {
		ld.UpdateCache[obj.GUID()] = &updateCache{
			UpdateFields:   interfaces.UpdateFieldsMap{},
			MovementUpdate: []byte{},
		}
	}

	cache := ld.UpdateCache[obj.GUID()]

	// First, copy over any fields which are in the new update.
	for k, v := range update.UpdateFields {
		cache.UpdateFields[k] = v
	}

	// Second, replace the movement update (if it has changed).
	if bytes.Compare(cache.MovementUpdate, update.MovementUpdate) != 0 {
		cache.MovementUpdate = update.MovementUpdate[:]
	}
}

// Updater manages sending object updates to sessions based on when objects have been changed.
type Updater struct {
	log *logrus.Entry
	om  *dynamic.ObjectManager

	sessionsLock sync.Mutex
	sessions     map[interfaces.GUID]*loginData // logged in character --> Session

	toUpdateLock sync.Mutex
	toUpdate     []interfaces.GUID
}

// NewUpdater makes a new updater object.
func NewUpdater(log *logrus.Entry, om *dynamic.ObjectManager) *Updater {
	u := &Updater{
		log: log.WithFields(logrus.Fields{
			"system": "Updater",
		}),
		om:       om,
		sessions: make(map[interfaces.GUID]*loginData),
		toUpdate: make([]interfaces.GUID, 0),
	}

	om.SetTriggerUpdateFor(func(obj interfaces.Object) {
		u.TriggerUpdateFor(obj)
	})

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
	u.om.GetPlayer(playerGUID).Unit.IsActive = true

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
	u.om.GetPlayer(playerGUID).Unit.IsActive = false

	delete(u.sessions, playerGUID)
	return nil
}

func (u *Updater) makeUpdates(
	loginData *loginData,
	player interfaces.Object,
	objectToUpdate interfaces.Object,
	outOfRangeUpdate *packet.OutOfRangeUpdate,
	objectUpdates *[]packet.ObjectUpdate) {
	if objectToUpdate.GetLocation() != nil {
		var movementUpdate []byte = nil
		switch objectToUpdateTyped := objectToUpdate.(type) {
		case *dynamic.Unit:
			movementUpdate = objectToUpdateTyped.MovementUpdate()
		case *dynamic.Player:
			movementUpdate = objectToUpdateTyped.MovementUpdate()
		}

		if objectToUpdate.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance {
			update := packet.ObjectUpdate{
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

				loginData.updateCache(objectToUpdate, &update)
			} else {
				if update.MovementUpdate != nil {
					update.UpdateType = static.UpdateTypeCreateObject2
				} else {
					update.UpdateType = static.UpdateTypeCreateObject
				}

				loginData.updateCache(objectToUpdate, &update)
			}

			if len(update.UpdateFields) != 0 {
				*objectUpdates = append(*objectUpdates, update)
			}
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
	outOfRangeUpdate := packet.OutOfRangeUpdate{
		GUIDS: make([]interfaces.GUID, 0),
	}

	objectUpdates := make([]packet.ObjectUpdate, 0)

	// Find all objects that are close to this player and make sure they have been updated.
	playerObj := u.om.Get(guid)
	for guid := range u.om.ActiveIDs {
		u.makeUpdates(loginData, playerObj, u.om.Get(guid), &outOfRangeUpdate, &objectUpdates)
	}

	if len(outOfRangeUpdate.GUIDS) == 0 && len(objectUpdates) == 0 {
		return
	}

	pkt := u.makeUpdateObjectPacket(outOfRangeUpdate, objectUpdates)
	loginData.Session.Send(pkt)
}

func (u *Updater) updateOtherPlayers(guid interfaces.GUID) {
	for playerGUID, loginData := range u.sessions {
		if playerGUID == guid {
			continue
		}

		outOfRangeUpdate := packet.OutOfRangeUpdate{
			GUIDS: make([]interfaces.GUID, 0),
		}

		objectUpdates := make([]packet.ObjectUpdate, 0)

		player := u.om.Get(playerGUID)
		obj := u.om.Get(guid)
		u.makeUpdates(loginData, player, obj, &outOfRangeUpdate, &objectUpdates)

		if len(outOfRangeUpdate.GUIDS) == 0 && len(objectUpdates) == 0 {
			return
		}

		loginData.Session.Send(u.makeUpdateObjectPacket(outOfRangeUpdate, objectUpdates))
	}
}

func (u *Updater) makeUpdateObjectPacket(outOfRangeUpdate packet.OutOfRangeUpdate, objectUpdates []packet.ObjectUpdate) ServerPacket {
	return &packet.ServerUpdateObject{
		OutOfRangeUpdates: outOfRangeUpdate,
		ObjectUpdates:     objectUpdates,
	}
}

func (u *Updater) doUpdate(guid interfaces.GUID) {
	// First, check to see if this is a player. If it is, then we
	// have to update the _player_ with all objects around them.
	if loginData, ok := u.sessions[guid]; ok {
		u.updatePlayer(guid, loginData)
	}

	// Second, we need to notify other players of this update.
	u.updateOtherPlayers(guid)
}

// Run starts the updater, which will constantly scan for object updates.
// Should be run as a goroutine.
func (u *Updater) Run() {
	for range time.Tick(time.Millisecond * 30) {
		u.toUpdateLock.Lock()

		// There are some object to update!
		for _, guid := range u.toUpdate {
			u.doUpdate(guid)
		}

		u.toUpdate = make([]interfaces.GUID, 0)
		u.toUpdateLock.Unlock()
	}
}

func (u *Updater) SendCombatUpdate(attacker interfaces.Unit, target interfaces.Unit, attackInfo *interfaces.AttackInfo) {
	// Find all players in range of either the attacker or target.
	for playerGUID, loginData := range u.sessions {
		player := u.om.Get(playerGUID)
		if player == nil {
			continue
		}

		if attacker.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance || target.GetLocation().Distance(player.GetLocation()) < MaxObjectDistance {
			loginData.Session.Send(&packet.ServerAttackerStateUpdate{
				HitInfo:      static.HitInfoNormalSwing,
				Attacker:     attacker.GUID(),
				Target:       target.GUID(),
				Damage:       int32(attackInfo.Damage),
				TargetState:  static.AttackTargetStateHit,
				MeleeSpellID: 0,
			})
		}
	}
}

func (u *Updater) TriggerUpdateFor(obj interfaces.Object) {
	u.doUpdate(obj.GUID())
}
