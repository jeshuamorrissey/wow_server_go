// Code generated by "stringer -type=Race -trimprefix=Race"; DO NOT EDIT.

package database

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RaceHuman-1]
	_ = x[RaceOrc-2]
	_ = x[RaceDwarf-3]
	_ = x[RaceNightElf-4]
	_ = x[RaceUndead-5]
	_ = x[RaceTauren-6]
	_ = x[RaceGnome-7]
	_ = x[RaceTroll-8]
	_ = x[RaceGoblin-9]
}

const _Race_name = "HumanOrcDwarfNightElfUndeadTaurenGnomeTrollGoblin"

var _Race_index = [...]uint8{0, 5, 8, 13, 21, 27, 33, 38, 43, 49}

func (i Race) String() string {
	i -= 1
	if i >= Race(len(_Race_index)-1) {
		return "Race(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Race_name[_Race_index[i]:_Race_index[i+1]]
}
