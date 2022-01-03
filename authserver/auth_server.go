package authserver

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/common/server"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
	"github.com/sirupsen/logrus"
)

func makeSession(reader io.Reader, writer io.Writer, log *logrus.Entry, config *world.WorldConfig) *session.Session {
	return session.NewSession(
		readHeader,
		writeHeader,
		opCodeToPacket,
		log,
		reader,
		writer,
		packet.NewState(config, log),
	)
}

func setupSession(sess *session.Session) {

}

// RunAuthServer takes as input a database and runs an auth server referencing
// it.
func RunAuthServer(port int, config *world.WorldConfig) {
	server.RunServer("LOGIN", port, config, makeSession, setupSession)
}
