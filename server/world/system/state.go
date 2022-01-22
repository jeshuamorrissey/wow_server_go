package system

import (
	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
)

// State contains useful state information passed to all packet
// methods.
type State struct {
	Session *Session

	Log *logrus.Entry

	Config  *config.Config
	OM      *dynamic.ObjectManager
	Updater *Updater

	Account   *config.Account
	Character *dynamic.Player
}
