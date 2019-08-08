package data

import (
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
)

// GetStartingItems is a utility which will return pointers to the item
// templates that a certain gender/race/class combination should have.
func GetStartingItems(gender c.Gender, class c.Class, race c.Race) []*Item {
	if gender == c.GenderMale {
		if class == c.ClassWarrior {
			if race == c.RaceHuman {
				return []*Item{
					Items[25],
				}
			}
		}
	}
	return []*Item{}
}
