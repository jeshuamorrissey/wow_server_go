// Code generated by "stringer -type=Gender -trimprefix=Gender"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GenderMale-0]
	_ = x[GenderFemale-1]
	_ = x[GenderNone-2]
}

const _Gender_name = "MaleFemaleNone"

var _Gender_index = [...]uint8{0, 4, 10, 14}

func (i Gender) String() string {
	if i >= Gender(len(_Gender_index)-1) {
		return "Gender(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Gender_name[_Gender_index[i]:_Gender_index[i+1]]
}
