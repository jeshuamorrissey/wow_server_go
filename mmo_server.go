package main

import (
	"log"
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver"
	"github.com/jinzhu/gorm"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

// GenerateTestData will generate all data required to have a reasonable test of the
// game system.
func GenerateTestData(db *gorm.DB) error {
	// Generate some accounts.
	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier("JESHUA", "JESHUA", salt)
	err := db.Create(&database.Account{Name: "JESHUA", VerifierStr: verifier.Text(16), SaltStr: salt.Text(16)}).Error
	if err != nil {
		return err
	}

	// Make some realms they can connect to.
	err = db.Create(&database.Realm{Name: "Sydney", Host: "localhost:5001"}).Error
	if err != nil {
		return err
	}

	err = db.Create(&database.Realm{Name: "Brisbane", Host: "localhost:5002"}).Error
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

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		authserver.RunAuthServer(5000, db)
	}()

	go func() {
		defer wg.Done()
		worldserver.RunWorldServer(5001, db)
	}()

	wg.Wait()
}
