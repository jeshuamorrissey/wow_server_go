package system

import (
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// State contains useful state information passed to all packet
// methods.
type State struct {
	Log *logrus.Entry

	DB      *gorm.DB
	OM      *object.Manager
	Updater *Updater

	Realm *database.Realm

	Account   *database.Account
	Character *object.Player
}
