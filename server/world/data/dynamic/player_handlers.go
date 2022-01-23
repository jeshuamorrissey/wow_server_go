package dynamic

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/components"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/messages"
)

// HandleModHealth is the handler for the ModHealth message.
//
// This handler will change the health by the requested amount.
func (p *Player) HandleModHealth(msg *messages.ModHealth) {
	p.ModHealth(msg.Amount, p.MaxHealth())
}

// HandleModPower is the handler for the ModPower message.
//
// This handler will change the power by the requested amount.
func (p *Player) HandleModPower(msg *messages.ModPower) {
	p.ModPower(msg.Amount, p.MaxPower())
}

// HandleAttack is the handler for the UnitAttack message.
//
// This handler will notify the target that they are being attacked, and then proceed
// to attack them with each weapon they have equipped.
func (p *Player) HandleAttack(msg *messages.UnitAttack) {
	target := GetObjectManager().Get(msg.Target)
	if target == nil {
		return
	}

	// First, we want to attack the target; so tell them we wish to attack them.
	target.SendUpdates([]interface{}{
		&messages.UnitRegisterAttack{Attacker: p.GUID()},
	})

	// Now, we have to attack with our weapons. We make one timer for each weapon (melee and ranged)
	// with the assumption that only one will be within range at a time (and only that one will do
	// damage).
	for _, weapon := range p.weapons() {
		p.Attack(p, target, weapon.GetTemplate().AttackRate, func() *components.Damage {
			return &components.Damage{
				Base: weapon.CalculateDamage(),
			}
		})
	}
}

func (p *Player) HandleRegisterAttacker(msg *messages.UnitRegisterAttack) {
	p.RegisterAttacker(msg.Attacker)
}

func (p *Player) HandleDeregisterAttacker(msg *messages.UnitDeregisterAttacker) {
	p.DeregisterAttacker(msg.Attacker)
}

func (p *Player) HandleAttackStop(msg *messages.UnitStopAttack) {
	// First, if we have a target, tell them we aren't attacking them.
	if p.Target > 0 {
		if target := GetObjectManager().Get(p.Target); target != nil {
			target.SendUpdates([]interface{}{
				&messages.UnitDeregisterAttacker{Attacker: p.GUID()},
			})
		}
	}

	// Now, we can stop attacking and clean up our attack timers.
	p.StopAttack()
}
