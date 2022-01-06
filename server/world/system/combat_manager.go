package system

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
)

type battle struct {
	attacker interfaces.Unit
	target   interfaces.Unit

	attackTimer          *time.Timer
	handleAttackCallback func(interfaces.AttackInfo)
}

// CombatManager manages sending object updates to sessions based on when objects have been changed.
type CombatManager struct {
	log     *logrus.Entry
	om      *dynamic.ObjectManager
	updater *Updater

	battles map[interfaces.GUID]*battle
}

// NewCombatManager makes a new CombatManager dynamic.
func NewCombatManager(log *logrus.Entry, om *dynamic.ObjectManager, updater *Updater) *CombatManager {
	u := &CombatManager{
		log: log.WithFields(logrus.Fields{
			"system": "Combat Manager",
		}),
		om:      om,
		updater: updater,
		battles: make(map[interfaces.GUID]*battle),
	}

	return u
}

func (cm *CombatManager) handleAttack(battle *battle) {
	cm.updater.SendCombatUpdate(battle.attacker, battle.target, *battle.attacker.ResolveMeleeAttack(battle.target))
	time.AfterFunc(battle.attacker.MeleeMainHandAttackRate(), func() {
		cm.handleAttack(battle)
	})
}

func (cm *CombatManager) StartAttack(attacker interfaces.Unit, target interfaces.Unit) {
	battle := &battle{
		attacker: attacker,
		target:   target,
	}

	cm.battles[attacker.GUID()] = battle
	cm.handleAttack(battle)
}

func (cm *CombatManager) StopAttack(attackerGUID interfaces.GUID) {
	battle, ok := cm.battles[attackerGUID]
	if ok {
		if battle.attackTimer != nil {
			battle.attackTimer.Stop()
		}

		delete(cm.battles, attackerGUID)
	}
}
