package dbc

import (
	"encoding/json"
	"fmt"
	"os"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

var (
	// Map of "<Class>:<Race>" --> map of stating stats.
	startingLocations map[string]*StartingLocation
)

// StartingLocation contains details about the starting location of a class.
type StartingLocation struct {
	Map, Zone  int
	X, Y, Z, O float32
}

// LoadStartingLocations reads the starting item JSON file and
// populates the startingItems map.
func LoadStartingLocations(jsonFile string) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(&startingLocations)
}

// GetStartingLocation is a utility which will return a mapping of stat
// to the starting value for that stat.
func GetStartingLocation(class c.Class, race c.Race) *StartingLocation {
	return startingLocations[fmt.Sprintf("%d:%d", class, race)]
}
