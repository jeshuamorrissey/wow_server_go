package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"gitlab.com/jeshuamorrissey/mmo_server/authserver"
	"gitlab.com/jeshuamorrissey/mmo_server/database"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

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
