package database

import (
	"github.com/jinzhu/gorm"
)

// Realm represents a worldserver the client can connect to.
type Realm struct {
	gorm.Model

	Name string
	Host string
}
