package system

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/sirupsen/logrus"
)

type battle struct {
	attacker object.UnitInterface
	target   object.UnitInterface

	attackTimer          *time.Timer
	handleAttackCallback func(object.AttackInfo)
}

// CombatManager manages sending object updates to sessions based on when objects have been changed.
type CombatManager struct {
	log     *logrus.Entry
	om      *object.Manager
	updater *Updater

	battles map[object.GUID]*battle
}

// NewCombatManager makes a new CombatManager object.
func NewCombatManager(log *logrus.Entry, om *object.Manager, updater *Updater) *CombatManager {
	u := &CombatManager{
		log: log.WithFields(logrus.Fields{
			"system": "Combat Manager",
		}),
		om:      om,
		updater: updater,
		battles: make(map[object.GUID]*battle),
	}

	return u
}

func (cm *CombatManager) handleAttack(battle *battle) {
	cm.updater.SendCombatUpdate(battle.attacker, battle.target, battle.attacker.Attack(battle.target))
	time.AfterFunc(battle.attacker.MeleeAttackRate(), func() {
		cm.handleAttack(battle)
	})
}

func (cm *CombatManager) StartAttack(attacker object.UnitInterface, target object.UnitInterface) {
	battle := &battle{
		attacker: attacker,
		target:   target,
	}

	cm.battles[attacker.GUID()] = battle
	cm.handleAttack(battle)
}

func (cm *CombatManager) StopAttack(attackerGUID object.GUID) {
	battle, ok := cm.battles[attackerGUID]
	if ok {
		if battle.attackTimer != nil {
			battle.attackTimer.Stop()
		}

		delete(cm.battles, attackerGUID)
	}
}
