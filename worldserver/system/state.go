package system

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
	"github.com/sirupsen/logrus"
)

// State contains useful state information passed to all packet
// methods.
type State struct {
	Session *Session

	Log *logrus.Entry

	Config        *world.WorldConfig
	OM            *object.Manager
	Updater       *Updater
	CombatManager *CombatManager

	Account   *world.Account
	Character *object.Player
}
