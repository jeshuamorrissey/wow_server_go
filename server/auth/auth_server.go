package auth

import (
	"net"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/session"
)

var serverName = "LOGIN"

var opCodeToPacket = map[static.OpCode]func() session.ClientPacket{
	static.OpCodeLoginChallenge: func() session.ClientPacket { return new(packet.ClientLoginChallenge) },
	static.OpCodeLoginProof:     func() session.ClientPacket { return new(packet.ClientLoginProof) },
	static.OpCodeRealmlist:      func() session.ClientPacket { return new(packet.ClientRealmlist) },
}

// RunAuthServer starts the authentication server, the first point of contact for the client.
func RunAuthServer(port int, config *config.Config) {
	log := logrus.WithFields(logrus.Fields{"server": serverName, "port": port})

	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	log.Infof("Listening for %v connections on :%v...", serverName, listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving %v connection from %v\n", serverName, conn.RemoteAddr())
		go session.NewSession(
			logrus.WithFields(logrus.Fields{"server": serverName, "account": "???"}),
			conn,
			conn,
			session.NewState(config, log, opCodeToPacket),
		).Run()
	}
}
