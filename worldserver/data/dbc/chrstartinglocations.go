package dbc

// StartingLocations represents data within the ChrStartingLocations.dbc file.
type StartingLocations struct {
	ID   int
	Race *Race
	Map  int
	Zone int
	X    float32
	Y    float32
	Z    float32
	O    float32
}

var (
	// StartingLocationsByID is the primary source of truth, storing data for for this DBC.
	StartingLocationsByID map[int]*StartingLocations
)

// Indexes for this DBC file, to make querying easier.
var (
	StartingLocationsByIndex map[*Race]*StartingLocations

	StartingLocationsHuman *StartingLocations
)

func init() {
	// Set the source of truth.
	StartingLocationsByID = map[int]*StartingLocations{
		0: &StartingLocations{
			ID:   0,
			Race: RaceHuman,
			Map:  0,
			Zone: 12,
			X:    -8949.95,
			Y:    -132.493,
			Z:    83.5312,
			O:    0.0,
		},
	}

	// Set the index.
	StartingLocationsByIndex = make(map[*Race]*StartingLocations)

	// Set the index values.
	StartingLocationsByIndex[RaceHuman] = StartingLocationsByID[0]

	// As there is only a single index, add some special convenience values.
	StartingLocationsHuman = StartingLocationsByID[0]
}
