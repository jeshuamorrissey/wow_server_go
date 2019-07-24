package main

import (
	"log"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver"
	"gitlab.com/jeshuamorrissey/mmo_server/database"
)

func main() {
	defer database.Cleanup()

	// Create database testdata.
	err := GenerateTestData()
	if err != nil {
		log.Fatalf("Failed to generate test data: %v\n", err)
	}

	authserver.RunAuthServer()
}
