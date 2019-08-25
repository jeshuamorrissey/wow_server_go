package worldserver

import (
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func makeSession(om *object.Manager, realm *database.Realm, reader io.Reader, writer io.Writer, log *logrus.Entry, db *gorm.DB) *system.Session {
	return system.NewSession(
		reader,
		writer,
		opCodeToPacket,
		db,
		om,
		log,
		realm,
	)
}

func setupSession(sess *system.Session) {
	pkt := packet.ServerAuthChallenge{Seed: 0}
	sess.Send(&pkt)
}

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(realmName string, port int, om *object.Manager, db *gorm.DB) {
	var realm database.Realm
	err := db.Where("name = ?", realmName).First(&realm).Error
	if err != nil {
		panic(fmt.Sprintf("Unknown realm %v", realmName))
	}

	log := logrus.WithFields(logrus.Fields{"server": "WORLD", "port": port})

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	log.Infof("Listening for WORLD connections on :%v...", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving WORLD connection from %v\n", conn.RemoteAddr())
		sessLog := logrus.WithFields(logrus.Fields{"server": "WORLD", "account": "???"})
		sess := makeSession(om, &realm, conn, conn, sessLog, db)
		setupSession(sess)
		go sess.Run()
	}
}
