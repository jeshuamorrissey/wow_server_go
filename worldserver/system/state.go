package system

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/config"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic"
	"github.com/sirupsen/logrus"
)

// State contains useful state information passed to all packet
// methods.
type State struct {
	Session *Session

	Log *logrus.Entry

	Config        *config.Config
	OM            *dynamic.ObjectManager
	Updater       *Updater
	CombatManager *CombatManager

	Account   *config.Account
	Character *dynamic.Player
}
