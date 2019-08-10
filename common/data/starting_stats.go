package data

import (
	"encoding/json"
	"fmt"
	"os"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

var (
	// Map of "<Class>:<Race>" --> map of stating stats.
	startingStats map[string]map[string]int
)

// LoadStartingStats reads the starting item JSON file and
// populates the startingItems map.
func LoadStartingStats(jsonFile string) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(&startingStats)
}

// GetStartingStats is a utility which will return a mapping of stat
// to the starting value for that stat.
func GetStartingStats(class c.Class, race c.Race) map[c.Stat]int {
	stats := startingStats[fmt.Sprintf("%d:%d", class, race)]

	return map[c.Stat]int{
		c.StatStrength:  stats["strength"],
		c.StatAgility:   stats["agility"],
		c.StatStamina:   stats["stamina"],
		c.StatIntellect: stats["intellect"],
		c.StatSpirit:    stats["spirit"],
	}
}
