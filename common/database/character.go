package database

import (
	"time"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jinzhu/gorm"
)

// Character represents a character in the game, linked to an account.
// The character's information is stored over this structure and a
// game object.
type Character struct {
	gorm.Model

	Name      string `gorm:"unique"`
	LastLogin *time.Time

	AccountID uint
	RealmID   uint

	GUID object.GUID

	// Flags.
	HideHelm        bool
	HideCloak       bool
	IsGhost         bool
	RenameNextLogin bool
}

// NewCharacter creates a new character entry in the database and
// returns a pointer to it.
func NewCharacter(
	om *object.Manager,
	name string,
	race c.Race, class c.Class, gender c.Gender,
	skinColor, face, hairStyle, hairColor, feature uint8) *Character {
	// TODO(jeshua): re-implement this.
	return nil
}

// Flags returns an set of flags based on the character's state.
func (char *Character) Flags() uint32 {
	var flags uint32
	if char.HideHelm {
		flags |= uint32(c.CharacterFlagHideHelm)
	}

	if char.HideCloak {
		flags |= uint32(c.CharacterFlagHideCloak)
	}

	if char.IsGhost {
		flags |= uint32(c.CharacterFlagGhost)
	}

	if char.RenameNextLogin {
		flags |= uint32(c.CharacterFlagRename)
	}

	return flags
}
