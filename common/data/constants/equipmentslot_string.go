// Code generated by "stringer -type=EquipmentSlot -trimprefix=EquipmentSlot"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EquipmentSlotHead-0]
	_ = x[EquipmentSlotNeck-1]
	_ = x[EquipmentSlotShoulders-2]
	_ = x[EquipmentSlotBody-3]
	_ = x[EquipmentSlotChest-4]
	_ = x[EquipmentSlotWaist-5]
	_ = x[EquipmentSlotLegs-6]
	_ = x[EquipmentSlotFeet-7]
	_ = x[EquipmentSlotWrists-8]
	_ = x[EquipmentSlotHands-9]
	_ = x[EquipmentSlotFinger1-10]
	_ = x[EquipmentSlotFinger2-11]
	_ = x[EquipmentSlotTrinket1-12]
	_ = x[EquipmentSlotTrinket2-13]
	_ = x[EquipmentSlotBack-14]
	_ = x[EquipmentSlotMainHand-15]
	_ = x[EquipmentSlotOffHand-16]
	_ = x[EquipmentSlotRanged-17]
	_ = x[EquipmentSlotTabard-18]
}

const _EquipmentSlot_name = "HeadNeckShouldersBodyChestWaistLegsFeetWristsHandsFinger1Finger2Trinket1Trinket2BackMainHandOffHandRangedTabard"

var _EquipmentSlot_index = [...]uint8{0, 4, 8, 17, 21, 26, 31, 35, 39, 45, 50, 57, 64, 72, 80, 84, 92, 99, 105, 111}

func (i EquipmentSlot) String() string {
	if i >= EquipmentSlot(len(_EquipmentSlot_index)-1) {
		return "EquipmentSlot(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EquipmentSlot_name[_EquipmentSlot_index[i]:_EquipmentSlot_index[i+1]]
}