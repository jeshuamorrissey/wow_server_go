package system

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
)

type autoAttackTimer struct {
	done   chan bool
	ticker *time.Ticker
}

type combatInfo struct {
	attacker interfaces.Unit
	target   interfaces.Unit

	autoAttackTimers []*autoAttackTimer
}

// CombatManager manages sending object updates to sessions based on when objects have been changed.
type CombatManager struct {
	log     *logrus.Entry
	om      *dynamic.ObjectManager
	updater *Updater

	lock sync.Mutex

	attackerToTarget  map[interfaces.GUID]*combatInfo
	targetToAttackers map[interfaces.GUID]map[interfaces.GUID]bool
}

// NewCombatManager makes a new CombatManager dynamic.
func NewCombatManager(log *logrus.Entry, om *dynamic.ObjectManager, updater *Updater) *CombatManager {
	u := &CombatManager{
		log: log.WithFields(logrus.Fields{
			"system": "Combat Manager",
		}),
		om:      om,
		updater: updater,

		attackerToTarget:  make(map[interfaces.GUID]*combatInfo, 0),
		targetToAttackers: make(map[interfaces.GUID]map[interfaces.GUID]bool, 0),
	}

	return u
}

func (cm *CombatManager) engageTarget(attacker interfaces.Unit, target interfaces.Unit) *combatInfo {
	cm.disengageTarget(attacker)

	cm.lock.Lock()
	defer cm.lock.Unlock()

	// TODO(jeshua): make it so that things in dynamic. can query the combat manager
	attacker.SetInCombat(true)
	target.SetInCombat(true)

	// Add a new combat entry.
	cm.attackerToTarget[attacker.GUID()] = &combatInfo{
		attacker:         attacker,
		target:           target,
		autoAttackTimers: make([]*autoAttackTimer, 0),
	}

	// Add the reverse-map entry.
	if _, ok := cm.targetToAttackers[target.GUID()]; !ok {
		cm.targetToAttackers[target.GUID()] = make(map[interfaces.GUID]bool, 0)
	}

	cm.targetToAttackers[target.GUID()][attacker.GUID()] = true
	return cm.attackerToTarget[attacker.GUID()]
}

func (cm *CombatManager) disengageTarget(attacker interfaces.Unit) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	if combatInfo, ok := cm.attackerToTarget[attacker.GUID()]; ok {
		// Remove the attacker from the reverse mapping.
		if attackers, ok := cm.targetToAttackers[combatInfo.target.GUID()]; ok {
			delete(attackers, attacker.GUID())
		}

		// The target might also not be in combat anymore, so stop.
		combatInfo.target.SetInCombat(cm.IsInCombat(combatInfo.target))

		// Stop any timers.
		for _, timer := range combatInfo.autoAttackTimers {
			timer.done <- true
		}

		// Forget about this combat.
		delete(cm.attackerToTarget, attacker.GUID())
	}

	attacker.SetInCombat(cm.IsInCombat(attacker))
}

func (cm *CombatManager) StartMeleeAttack(attacker interfaces.Unit, target interfaces.Unit) {
	combatInfo := cm.engageTarget(attacker, target)

	// We are performing melee attacks, so add timers for the two melee weapons.
	if attacker.MeleeOffHandAttackRate() != 0 {
		offHandAutoAttackTimer := &autoAttackTimer{
			done:   make(chan bool),
			ticker: time.NewTicker(attacker.MeleeOffHandAttackRate()),
		}

		combatInfo.autoAttackTimers = append(combatInfo.autoAttackTimers, offHandAutoAttackTimer)
		go cm.resolveMeleeAttack(offHandAutoAttackTimer, combatInfo, func() *interfaces.AttackInfo {
			return attacker.ResolveMainHandAttack(target)
		})
	}

	mainHandAutoAttackTimer := &autoAttackTimer{
		done:   make(chan bool),
		ticker: time.NewTicker(attacker.MeleeMainHandAttackRate()),
	}
	combatInfo.autoAttackTimers = append(combatInfo.autoAttackTimers, mainHandAutoAttackTimer)
	go cm.resolveMeleeAttack(mainHandAutoAttackTimer, combatInfo, func() *interfaces.AttackInfo {
		return attacker.ResolveMainHandAttack(target)
	})
}

func (cm *CombatManager) StopAttack(attacker interfaces.Unit) {
	cm.disengageTarget(attacker)
}

func (cm *CombatManager) resolveMeleeAttackSingle(attackInfo *interfaces.AttackInfo, combatInfo *combatInfo) {
	switch typedUnit := combatInfo.target.(type) {
	case *dynamic.Unit:
		typedUnit.TakeDamage(attackInfo.Damage)
	case *dynamic.Player:
		typedUnit.TakeDamage(attackInfo.Damage)
	}
	cm.updater.TriggerUpdateFor(combatInfo.target)
	cm.updater.SendCombatUpdate(combatInfo.attacker, combatInfo.target, attackInfo)
}

func (cm *CombatManager) resolveMeleeAttack(autoAttackTimer *autoAttackTimer, combatInfo *combatInfo, resolveAttack func() *interfaces.AttackInfo) {
	cm.resolveMeleeAttackSingle(resolveAttack(), combatInfo)
	for {
		select {
		case <-autoAttackTimer.ticker.C:
			cm.resolveMeleeAttackSingle(resolveAttack(), combatInfo)
		case <-autoAttackTimer.done:
			autoAttackTimer.ticker.Stop()
			return
		}
	}
}

// IsInCombat determines whether the given unit is in combat. A unit is in combat if at least one
// other unit is attacking it.
func (cm *CombatManager) IsInCombat(unit interfaces.Unit) bool {
	if cm.GetTargetOf(unit) != nil {
		return true
	}

	if attackerMap, ok := cm.targetToAttackers[unit.GUID()]; ok {
		return len(attackerMap) >= 1
	}

	return false
}

// GetTargetOf will return the target of the given unit, or None if it isn't in combat.
func (cm *CombatManager) GetTargetOf(unit interfaces.Unit) interfaces.Unit {
	if attackerMap, ok := cm.attackerToTarget[unit.GUID()]; ok {
		return attackerMap.target
	}

	return nil
}
