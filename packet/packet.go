package packet

import (
	"bufio"
)

// ServerPacket is a type of packet which can have it's contents written to a
// byte buffer.
type ServerPacket interface {
	// Bytes writes out the packet as a byte array.
	Bytes() []byte
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
