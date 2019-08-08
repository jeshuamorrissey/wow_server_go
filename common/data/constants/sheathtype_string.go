// Code generated by "stringer -type=SheathType -trimprefix=SheathType"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SheathTypeNone-0]
	_ = x[SheathTypeMainHand-1]
	_ = x[SheathTypeOffHand-2]
	_ = x[SheathTypeLargeWeaponLeft-3]
	_ = x[SheathTypeLargeWeaponRight-4]
	_ = x[SheathTypeHipWeaponLeft-5]
	_ = x[SheathTypeHipWeaponRight-6]
	_ = x[SheathTypeShield-7]
}

const _SheathType_name = "NoneMainHandOffHandLargeWeaponLeftLargeWeaponRightHipWeaponLeftHipWeaponRightShield"

var _SheathType_index = [...]uint8{0, 4, 12, 19, 34, 50, 63, 77, 83}

func (i SheathType) String() string {
	if i >= SheathType(len(_SheathType_index)-1) {
		return "SheathType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SheathType_name[_SheathType_index[i]:_SheathType_index[i+1]]
}