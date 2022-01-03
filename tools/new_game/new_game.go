package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world/initial_data"
)

func main() {
	namePtr := flag.String("name", "", "The name of the save file.")

	flag.Parse()

	name := *namePtr
	if name == "" {
		log.Fatal("Missing required flag -name")
		return
	}

	// Make an empty configuration file.
	config := world.NewWorldConfig(name)

	// Make a character.
	var err error
	config.Accounts[0].Character, err = initial_data.NewCharacter(
		config,
		"Jeshua",
		dbc.RaceHuman, dbc.ClassWarrior, c.GenderMale,
		1, 1, 1, 1, 1)
	if err != nil {
		log.Fatalf("Failed to create character: %v", err)
		return
	}

	// Populate the world.
	initial_data.PopulateWorld(config)

	jsonContent, err := json.Marshal(&config)
	if err != nil {
		log.Fatalf("Failed to save config to JSON: %v", err)
		return
	}

	fmt.Print(string(jsonContent))
}
