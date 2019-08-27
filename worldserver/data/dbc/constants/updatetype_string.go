// Code generated by "stringer -type=UpdateType -trimprefix=UpdateType"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UpdateTypeValues-0]
	_ = x[UpdateTypeMovement-1]
	_ = x[UpdateTypeCreateObject-2]
	_ = x[UpdateTypeCreateObject2-3]
	_ = x[UpdateTypeOutOfRangeObjects-4]
	_ = x[UpdateTypeNearObjects-5]
}

const _UpdateType_name = "ValuesMovementCreateObjectCreateObject2OutOfRangeObjectsNearObjects"

var _UpdateType_index = [...]uint8{0, 6, 14, 26, 39, 56, 67}

func (i UpdateType) String() string {
	if i >= UpdateType(len(_UpdateType_index)-1) {
		return "UpdateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _UpdateType_name[_UpdateType_index[i]:_UpdateType_index[i+1]]
}