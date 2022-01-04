package interfaces

import "time"

// AttackInfo encapsulates all information about the result of an attack.
type AttackInfo struct {
	// Damage is the amount of health which should be removed from the target.
	Damage int
}

type Unit interface {
	Object

	/// Game logic.
	// MeleeMainHandAttackRate should return the time between attacks with the main-hand weapon.
	MeleeMainHandAttackRate() time.Duration

	// MeleeOffHandAttackRate should return the time between attacks with the off-hand weapon. A value
	// of nil represents no weapon being equipped in the off-hand slot.
	MeleeOffHandAttackRate() time.Duration

	// ResolveMeleeAttack should simulate an attack from this unit to another unit and return a
	// structure describing the result of the interaction.
	ResolveMeleeAttack(Unit) *AttackInfo
}
