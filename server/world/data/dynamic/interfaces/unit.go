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
	// Initialize should start any long-running operations (such as timers) which are relevant to
	// the unit (e.g. a timer to restore mana, or to check nearby for threats).
	Initialize()

	// MeleeMainHandAttackRate should return the time between attacks with the main-hand weapon.
	MeleeMainHandAttackRate() time.Duration

	// ResolveMainHandAttack should simulate an attack from this unit to another unit and return a
	// structure describing the result of the interaction.
	ResolveMainHandAttack(Unit) *AttackInfo

	// MeleeOffHandAttackRate should return the time between attacks with the off-hand weapon. A value
	// of nil represents no weapon being equipped in the off-hand slot.
	MeleeOffHandAttackRate() time.Duration

	// ResolveOffHandAttack should simulate an attack from this unit to another unit and return a
	// structure describing the result of the interaction.
	ResolveOffHandAttack(Unit) *AttackInfo

	// SetInCombat can be used to specify that the given unit is in combat.
	SetInCombat(bool)
}
