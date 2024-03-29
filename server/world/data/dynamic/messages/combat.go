package messages

import "github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"

type (
	UnitAttack struct {
		Target interfaces.GUID
	}

	UnitStopAttack struct{}

	UnitRegisterAttack struct {
		Attacker interfaces.GUID
	}

	UnitDeregisterAttacker struct {
		Attacker interfaces.GUID
	}

	UnitDied struct {
		DeadUnit interfaces.GUID
	}
)
