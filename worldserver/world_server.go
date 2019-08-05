package worldserver

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/server"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func makeSession(reader io.Reader, writer io.Writer, log *logrus.Entry, db *gorm.DB) *session.Session {
	return session.NewSession(
		readHeader,
		writeHeader,
		opCodeToPacket,
		log,
		reader,
		writer,
		packet.NewState(db, log),
	)
}

func setupSession(sess *session.Session) {
	pkt := packet.ServerAuthChallenge{Seed: 0}
	sess.SendPacket(&pkt)
}

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(port int, db *gorm.DB) {
	server.RunServer("world", port, db, makeSession, setupSession)
}
