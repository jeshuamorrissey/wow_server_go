package packet

import (
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/objects"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// State represents all information required by the AuthServer.
type State struct {
	db  *gorm.DB
	om  *objects.ObjectManager
	log *logrus.Entry

	Account database.Account
	Realm   *database.Realm

	// Some counters required for encrypting the header.
	SendI, SendJ uint8
	RecvI, RecvJ uint8
}

// NewState creates a new state based on the given DB connection.
func NewState(om *objects.ObjectManager, realm *database.Realm, db *gorm.DB, log *logrus.Entry) *State {
	return &State{db: db, log: log, Realm: realm, om: om}
}

// DB returns a reference to the Database object stored in this state.
func (s *State) DB() *gorm.DB {
	return s.db
}

// Log returns a reference to the Database object stored in this state.
func (s *State) Log() *logrus.Entry {
	return s.log
}

// OM returns a reference to the global object manager.
func (s *State) OM() *objects.ObjectManager {
	return s.om
}

// AddLogField adds a new field to the logger for this state.
func (s *State) AddLogField(key string, value interface{}) {
	*s.log = *s.log.WithField(key, value)
}
