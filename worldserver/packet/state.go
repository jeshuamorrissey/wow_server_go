package packet

import (
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// State represents all information required by the AuthServer.
type State struct {
	_db *gorm.DB
	log *logrus.Entry

	Account database.Account

	// Some counters required for encrypting the header.
	SendI, SendJ uint8
	RecvI, RecvJ uint8
}

// NewState creates a new state based on the given DB connection.
func NewState(db *gorm.DB, log *logrus.Entry) *State {
	return &State{_db: db, log: log}
}

// DB returns a reference to the Database object stored in this state.
func (s *State) DB() *gorm.DB {
	return s._db
}

// Log returns a reference to the Database object stored in this state.
func (s *State) Log() *logrus.Entry {
	return s.log
}

// AddLogField adds a new field to the logger for this state.
func (s *State) AddLogField(key string, value interface{}) {
	*s.log = *s.log.WithField(key, value)
}
