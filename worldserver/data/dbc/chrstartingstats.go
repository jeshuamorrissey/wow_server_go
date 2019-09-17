package dbc

// StartingStats represents data within the ChrStartingStats.dbc file.
type StartingStats struct {
	ID       int
	Class    *Class
	Race     *Race
	Strength int
}

var (
	// StartingStatsByID is the primary source of truth, storing data for for this DBC.
	StartingStatsByID map[int]*StartingStats
)

// Indexes for this DBC file, to make querying easier.
var (
	StartingStatsByIndex map[*Class]map[*Race]*StartingStats
)

func init() {
	// Set the source of truth.
	StartingStatsByID = map[int]*StartingStats{
		0: &StartingStats{
			ID:       0,
			Class:    ClassWarrior,
			Race:     RaceHuman,
			Strength: 0,
		},
	}

	// Set the index.
	StartingStatsByIndex = make(map[*Class]map[*Race]*StartingStats)

	// Initialize sub-maps for each indexed field.
	StartingStatsByIndex[ClassWarrior] = make(map[*Race]*StartingStats)

	// Set the index values.
	StartingStatsByIndex[ClassWarrior][RaceHuman] = StartingStatsByID[0]

}
