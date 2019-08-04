package worldserver

import (
	"log"
	"net"
	"strconv"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jinzhu/gorm"
)

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(port int, db *gorm.DB) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	// Main control loop.
	log.Printf("Listening for WORLD connections on :%v...\n", port)

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving WORLD connection from %v\n", conn.RemoteAddr())

		// TODO(jeshua): send an AUTH_CHALLENGE packet.
		go session.NewSession(
			readHeader,
			opCodeToPacket,
			packet.OpCodeName,
			packet.NewState(db)).Run(conn, conn)
	}
}
