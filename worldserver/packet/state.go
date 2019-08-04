package packet

import (
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jinzhu/gorm"
)

// State represents all information required by the AuthServer.
type State struct {
	_db *gorm.DB

	Account database.Account
}

// NewState creates a new state based on the given DB connection.
func NewState(db *gorm.DB) *State {
	return &State{_db: db}
}

// DB returns a reference to the Database object stored in this state.
func (s *State) DB() *gorm.DB {
	return s._db
}
