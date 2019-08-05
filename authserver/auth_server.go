package authserver

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/common/server"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
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
		packet.NewState(db),
	)
}

func setupSession(sess *session.Session) {

}

// RunAuthServer takes as input a database and runs an auth server referencing
// it.
func RunAuthServer(port int, db *gorm.DB) {
	server.RunServer("login", port, db, makeSession, setupSession)
}
