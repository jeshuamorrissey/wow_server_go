// Code generated by "stringer -type=InventoryType -trimprefix=InventoryType"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InventoryTypeNonEquip-0]
	_ = x[InventoryTypeHead-1]
	_ = x[InventoryTypeNeck-2]
	_ = x[InventoryTypeShoulders-3]
	_ = x[InventoryTypeBody-4]
	_ = x[InventoryTypeChest-5]
	_ = x[InventoryTypeWaist-6]
	_ = x[InventoryTypeLegs-7]
	_ = x[InventoryTypeFeet-8]
	_ = x[InventoryTypeWrists-9]
	_ = x[InventoryTypeHands-10]
	_ = x[InventoryTypeFinger-11]
	_ = x[InventoryTypeTrinket-12]
	_ = x[InventoryTypeWeapon-13]
	_ = x[InventoryTypeShield-14]
	_ = x[InventoryTypeRanged-15]
	_ = x[InventoryTypeCloak-16]
	_ = x[InventoryType2HWeapon-17]
	_ = x[InventoryTypeBag-18]
	_ = x[InventoryTypeTabard-19]
	_ = x[InventoryTypeRobe-20]
	_ = x[InventoryTypeWeaponMainHand-21]
	_ = x[InventoryTypeWeaponOffHand-22]
	_ = x[InventoryTypeHoldable-23]
	_ = x[InventoryTypeAmmo-24]
	_ = x[InventoryTypeThrown-25]
	_ = x[InventoryTypeRangedRight-26]
	_ = x[InventoryTypeQuiver-27]
	_ = x[InventoryTypeRelic-28]
}

const _InventoryType_name = "NonEquipHeadNeckShouldersBodyChestWaistLegsFeetWristsHandsFingerTrinketWeaponShieldRangedCloak2HWeaponBagTabardRobeWeaponMainHandWeaponOffHandHoldableAmmoThrownRangedRightQuiverRelic"

var _InventoryType_index = [...]uint8{0, 8, 12, 16, 25, 29, 34, 39, 43, 47, 53, 58, 64, 71, 77, 83, 89, 94, 102, 105, 111, 115, 129, 142, 150, 154, 160, 171, 177, 182}

func (i InventoryType) String() string {
	if i >= InventoryType(len(_InventoryType_index)-1) {
		return "InventoryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InventoryType_name[_InventoryType_index[i]:_InventoryType_index[i+1]]
}