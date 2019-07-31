package packet

import (
	"math/big"

	"github.com/jinzhu/gorm"
	db "gitlab.com/jeshuamorrissey/mmo_server/database"
)

type State struct {
	_db *gorm.DB

	PublicEphemeral  big.Int
	PrivateEphemeral big.Int

	Account db.Account
}

// NewState creates a new state based on the given DB connection.
func NewState(db *gorm.DB) *State {
	return &State{_db: db}
}

// DB returns a reference to the Database object stored in this state.
func (s *State) DB() *gorm.DB {
	return s._db
}
