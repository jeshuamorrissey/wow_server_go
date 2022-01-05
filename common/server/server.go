package server

import (
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/config"
	"github.com/sirupsen/logrus"
)

// RunServer will take control of the process and run a server. This is
// designed to be run as a goroutine.
func RunServer(
	name string,
	port int,
	config *config.Config,
	makeSession func(io.Reader, io.Writer, *logrus.Entry, *config.Config) *session.Session,
	setupSession func(*session.Session)) {
	log := logrus.WithFields(logrus.Fields{"server": name, "port": port})

	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	log.Infof("Listening for %v connections on :%v...", strings.ToUpper(name), listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving %v connection from %v\n", strings.ToUpper(name), conn.RemoteAddr())
		sessLog := logrus.WithFields(logrus.Fields{"server": name, "account": "???"})
		sess := makeSession(conn, conn, sessLog, config)
		setupSession(sess)
		go sess.Run()
	}
}
