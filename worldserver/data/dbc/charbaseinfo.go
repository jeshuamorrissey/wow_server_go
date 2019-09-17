package dbc

// CharBaseInfo represents data within the CharBaseInfo.dbc file.
type CharBaseInfo struct {
	Race  uint8
	Class uint8
}

var (
	// CharBaseInfoByID is the primary source of truth, storing data for for this DBC.
	CharBaseInfoByID map[int]*CharBaseInfo
)

// Indexes for this DBC file, to make querying easier.
var ()

func init() {
	// Set the source of truth.
	CharBaseInfoByID = map[int]*CharBaseInfo{
		0: &CharBaseInfo{
			Race:  1,
			Class: 1,
		},
		1: &CharBaseInfo{
			Race:  1,
			Class: 2,
		},
		2: &CharBaseInfo{
			Race:  1,
			Class: 4,
		},
		3: &CharBaseInfo{
			Race:  1,
			Class: 5,
		},
		4: &CharBaseInfo{
			Race:  1,
			Class: 8,
		},
		5: &CharBaseInfo{
			Race:  1,
			Class: 9,
		},
		6: &CharBaseInfo{
			Race:  2,
			Class: 1,
		},
		7: &CharBaseInfo{
			Race:  2,
			Class: 3,
		},
		8: &CharBaseInfo{
			Race:  2,
			Class: 4,
		},
		9: &CharBaseInfo{
			Race:  2,
			Class: 7,
		},
		10: &CharBaseInfo{
			Race:  2,
			Class: 9,
		},
		11: &CharBaseInfo{
			Race:  3,
			Class: 1,
		},
		12: &CharBaseInfo{
			Race:  3,
			Class: 2,
		},
		13: &CharBaseInfo{
			Race:  3,
			Class: 3,
		},
		14: &CharBaseInfo{
			Race:  3,
			Class: 4,
		},
		15: &CharBaseInfo{
			Race:  3,
			Class: 5,
		},
		16: &CharBaseInfo{
			Race:  3,
			Class: 8,
		},
		17: &CharBaseInfo{
			Race:  4,
			Class: 1,
		},
		18: &CharBaseInfo{
			Race:  4,
			Class: 3,
		},
		19: &CharBaseInfo{
			Race:  4,
			Class: 4,
		},
		20: &CharBaseInfo{
			Race:  4,
			Class: 5,
		},
		21: &CharBaseInfo{
			Race:  4,
			Class: 11,
		},
		22: &CharBaseInfo{
			Race:  5,
			Class: 1,
		},
		23: &CharBaseInfo{
			Race:  5,
			Class: 4,
		},
		24: &CharBaseInfo{
			Race:  5,
			Class: 5,
		},
		25: &CharBaseInfo{
			Race:  5,
			Class: 8,
		},
		26: &CharBaseInfo{
			Race:  5,
			Class: 9,
		},
		27: &CharBaseInfo{
			Race:  6,
			Class: 1,
		},
		28: &CharBaseInfo{
			Race:  6,
			Class: 3,
		},
		29: &CharBaseInfo{
			Race:  6,
			Class: 7,
		},
		30: &CharBaseInfo{
			Race:  6,
			Class: 11,
		},
		31: &CharBaseInfo{
			Race:  7,
			Class: 1,
		},
		32: &CharBaseInfo{
			Race:  7,
			Class: 4,
		},
		33: &CharBaseInfo{
			Race:  7,
			Class: 8,
		},
		34: &CharBaseInfo{
			Race:  7,
			Class: 9,
		},
		35: &CharBaseInfo{
			Race:  8,
			Class: 1,
		},
		36: &CharBaseInfo{
			Race:  8,
			Class: 3,
		},
		37: &CharBaseInfo{
			Race:  8,
			Class: 4,
		},
		38: &CharBaseInfo{
			Race:  8,
			Class: 5,
		},
		39: &CharBaseInfo{
			Race:  8,
			Class: 7,
		},
		40: &CharBaseInfo{
			Race:  8,
			Class: 8,
		},
	}

}
