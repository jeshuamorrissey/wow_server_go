package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/common/database"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
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

func main() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	log.Printf("Created in-memory database")

	database.Setup(db)

	defer db.Close()

	// Create database testdata.
	err = GenerateTestData(db)
	if err != nil {
		log.Fatalf("Failed to generate test data: %v\n", err)
	}

	authserver.RunAuthServer(db)
}
