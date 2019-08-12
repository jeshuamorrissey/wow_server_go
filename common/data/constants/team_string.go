// Code generated by "stringer -type=Team -trimprefix=Team"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TeamAlliance-67]
	_ = x[TeamHorde-469]
}

const (
	_Team_name_0 = "Alliance"
	_Team_name_1 = "Horde"
)

func (i Team) String() string {
	switch {
	case i == 67:
		return _Team_name_0
	case i == 469:
		return _Team_name_1
	default:
		return "Team(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
