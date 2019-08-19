package packet

import (
	"io"
)

// Packet is a generic packet.
type Packet interface {
	// OpCode returns the opcode for the given packet as an int.
	OpCode() OpCode
}

// ServerPacket is a packet sent from this server to a client.
type ServerPacket interface {
	Packet

	// ToBytes writes the packet out to an array of bytes.
	ToBytes(*State) ([]byte, error)
}

// ClientPacket is a packet sent from the client to this server.
type ClientPacket interface {
	Packet

	// FromBytes reads the packet from a generic reader.
	FromBytes(*State, io.Reader) error

	// Handle the packet and return a list of server packets to send back
	// to the client.
	Handle(*State) ([]ServerPacket, error)
}
