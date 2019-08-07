package database

import (
	"github.com/jinzhu/gorm"
)

// Character represents a character in the game, linked to an account.
// The character's information is stored over this structure and a
// game object.
type Character struct {
	gorm.Model

	Name string

	AccountID uint
	RealmID   uint

	Object GameObjectPlayer
}
