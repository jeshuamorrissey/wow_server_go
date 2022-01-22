package components

import (
	"time"

	"github.com/jeshuamorrissey/wow_server_go/server/world/channels"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

type Damage struct {
	Base map[static.SpellSchool]int
}

type combatHook func(*Damage)

type autoAttackTimer struct {
	done   chan bool
	ticker *time.Ticker
}

type Combat struct {
	Target    interfaces.GUID
	Attackers map[interfaces.GUID]bool

	OutgoingDamageMods map[string]func(*Damage)
	IncomingDamageMods map[string]func(*Damage)

	autoAttackTimers []*autoAttackTimer
}

// IsInCombat will return true iff we are in combat.
func (c *Combat) IsInCombat() bool {
	return c.Target != 0 || len(c.Attackers) > 0
}

func (c *Combat) StopAttack() {
	for _, timer := range c.autoAttackTimers {
		timer.done <- true
	}

	c.autoAttackTimers = make([]*autoAttackTimer, 0)
}

// RegisterAttacker will note that the given object is attacking us.
func (c *Combat) RegisterAttacker(attacker interfaces.GUID) {
	if c.Attackers == nil {
		c.Attackers = make(map[interfaces.GUID]bool)
	}

	c.Attackers[attacker] = true
}

func (c *Combat) DeregisterAttacker(attacker interfaces.GUID) {
	delete(c.Attackers, attacker)
}

func (c *Combat) resolveSingleAttack(attacker interfaces.Object, target interfaces.Object, damage *Damage) {
	for _, damageMod := range c.OutgoingDamageMods {
		damageMod(damage)
	}

	finalDamage := 0
	for _, damage := range damage.Base {
		finalDamage += damage
	}

	channels.CombatUpdates <- &channels.CombatUpdate{
		Attacker: attacker,
		Target:   target,
		AttackInfo: &interfaces.AttackInfo{
			Damage: finalDamage,
		},
	}

	target.SendUpdates([]interface{}{
		&messages.UnitModHealth{Amount: -finalDamage},
	})
}

// Attack will start a goroutine which will manage attacking.
func (c *Combat) Attack(attacker interfaces.Object, target interfaces.Object, attackRate time.Duration, calculateBaseDamage func() *Damage) {
	c.Target = target.GUID()

	if c.autoAttackTimers == nil {
		c.autoAttackTimers = make([]*autoAttackTimer, 0)
	}

	autoAttackTimer := &autoAttackTimer{
		done:   make(chan bool),
		ticker: time.NewTicker(attackRate),
	}

	c.autoAttackTimers = append(c.autoAttackTimers, autoAttackTimer)

	c.resolveSingleAttack(attacker, target, calculateBaseDamage())
	go func() {
		for {
			select {
			case <-autoAttackTimer.ticker.C:
				c.resolveSingleAttack(attacker, target, calculateBaseDamage())
			case <-autoAttackTimer.done:
				autoAttackTimer.ticker.Stop()
				return
			}
		}
	}()
}
