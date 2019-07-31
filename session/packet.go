package session

import "io"

// ServerPacket is a type of packet which can have it's contents written to a
// byte buffer.
type ServerPacket interface {
	// OpCode returns the opcode for the given packet as an int.
	OpCode() OpCode

	// Bytes writes out the packet as a byte array.
	Bytes() []byte
}

// ClientPacket is a type of packet which can have it's contents filled in
// from a byte buffer.
type ClientPacket interface {
	// Read takes as input a buffer and populates the fields of the packet.
	Read(io.Reader) error

	// Handle the packet and return a list of server packets to send back
	// to the client. It takes as input some session information (which
	// depends on the type of session).
	Handle(State) ([]ServerPacket, error)
}