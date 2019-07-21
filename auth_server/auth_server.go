package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
		return
	}

	log.Printf("Listening for connections on :5000...\n")

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
			return
		}

		log.Printf("Receiving connection from %v\n", conn.RemoteAddr())

		// Make a new buffer to read from.
		go RunSession(bufio.NewReader(conn), bufio.NewWriter(conn))
	}
}
