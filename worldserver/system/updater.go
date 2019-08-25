package system

import (
	"fmt"
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/sirupsen/logrus"
)

// Updater manages sending object updates to sessions based on when objects have been changed.
type Updater struct {
	log           *logrus.Entry
	objectManager *object.Manager

	sessionsLock sync.Mutex
	sessions     map[object.GUID]*Session // logged in character --> Session
}

// NewUpdater makes a new updater object.
func NewUpdater(log *logrus.Entry) *Updater {
	return &Updater{
		log: log.WithFields(logrus.Fields{
			"system": "Updater",
		}),
		objectManager: object.NewManager(log),

		sessions: make(map[object.GUID]*Session),
	}
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

	u.sessions[playerGUID] = session

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

// Run starts the updater, which will constantly scan for object updates.
// Should be run as a goroutine.
func (u *Updater) Run() {
	for {

	}
}
