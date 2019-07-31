package main

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/jeshuamorrissey/mmo_server/authserver/srp"
	db "gitlab.com/jeshuamorrissey/mmo_server/database"
)

// GenerateTestData will generate all data required to have a reasonable test of the
// game system.
func GenerateTestData(database *gorm.DB) error {

	// Generate some accounts.
	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier("JESHUA", "JESHUA", salt)
	err := database.Create(&db.Account{Name: "jeshua", VerifierStr: verifier.Text(16), SaltStr: salt.Text(16)}).Error
	if err != nil {
		return err
	}

	// Make some realms they can connect to.
	err = database.Create(&db.Realm{Name: "Sydney", Host: "localhost:5001"}).Error
	if err != nil {
		return err
	}

	err = database.Create(&db.Realm{Name: "Brisbane", Host: "localhost:5002"}).Error
	if err != nil {
		return err
	}

	return nil
}
