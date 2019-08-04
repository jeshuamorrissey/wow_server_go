package authserver

import (
	"log"
	"net"
	"strconv"

	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jinzhu/gorm"
)

// RunAuthServer takes as input a database and runs an auth server referencing
// it.
func RunAuthServer(port int, db *gorm.DB) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	// Main control loop.
	log.Printf("Listening for AUTH connections on :%v...\n", port)

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving AUTH connection from %v\n", conn.RemoteAddr())

		go session.NewSession(
			readHeader,
			writeHeader,
			opCodeToPacket,
			conn,
			conn,
			packet.NewState(db),
		).Run()
	}
}
