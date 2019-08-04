package authserver

import (
	"log"
	"net"

	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jinzhu/gorm"
)

// RunAuthServer takes as input a database and runs an auth server referencing
// it.
func RunAuthServer(db *gorm.DB) {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	// Main control loop.
	log.Printf("Listening for connections on :5000...\n")

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving connection from %v\n", conn.RemoteAddr())

		go session.NewSession(
			readHeader,
			opCodeToPacket,
			packet.OpCodeName,
			packet.NewState(db)).Run(conn, conn)
	}
}
