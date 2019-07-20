package packet

import (
	"bufio"
)

// ServerPacket is a type of packet which can have it's contents written to a
// byte buffer.
type ServerPacket interface {
	// Write takes as input a buffer and writes it's contents to it.
	Write(*bufio.Writer) error
}

// ClientPacket is a type of packet which can have it's contents filled in
// from a byte buffer.
type ClientPacket interface {
	// Read takes as input a buffer and populates the fields of the packet.
	Read(*bufio.Reader) error

	// Handle the packet and return a list of server packets to send back
	// to the client.
	Handle() ([]ServerPacket, error)
}
