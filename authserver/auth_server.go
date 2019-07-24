package authserver

import (
	"bufio"
	"log"
	"net"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver/packet"
)

// RunAuthServer takes as input a database and runs an auth server referencing
// it.
func RunAuthServer() {
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

		// Make a new buffer to read from.
		go packet.RunSession(bufio.NewReader(conn), bufio.NewWriter(conn))
	}
}
