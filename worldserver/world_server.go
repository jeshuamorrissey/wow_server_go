package worldserver

import (
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/objects"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/common/server"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func makeSession(om *objects.ObjectManager, realm *database.Realm, reader io.Reader, writer io.Writer, log *logrus.Entry, db *gorm.DB) *session.Session {
	return session.NewSession(
		readHeader,
		writeHeader,
		opCodeToPacket,
		log,
		reader,
		writer,
		packet.NewState(om, realm, db, log),
	)
}

func setupSession(sess *session.Session) {
	pkt := packet.ServerAuthChallenge{Seed: 0}
	sess.SendPacket(&pkt)
}

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(realmName string, port int, om *objects.ObjectManager, db *gorm.DB) {
	var realm database.Realm
	err := db.Where("name = ?", realmName).First(&realm).Error
	if err != nil {
		panic(fmt.Sprintf("Unknown realm %v", realmName))
	}

	server.RunServer(
		"world",
		port,
		db,
		func(reader io.Reader, writer io.Writer, log *logrus.Entry, db *gorm.DB) *session.Session {
			return makeSession(om, &realm, reader, writer, log, db)
		},
		setupSession)
}
