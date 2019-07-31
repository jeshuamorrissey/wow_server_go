package database

import (
	"github.com/jinzhu/gorm"
)

// Setup creates performs the automigration of the required types.
func Setup(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Realm{})
}
