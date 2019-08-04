package database

import (
	"github.com/jinzhu/gorm"
)

// Realm represents a worldserver the client can connect to.
type Realm struct {
	gorm.Model

	Name string `gorm:"unique"`
	Host string
}
