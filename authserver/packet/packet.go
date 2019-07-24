package packet

import (
	"bufio"
	"log"
)

var (
	opCodeToPacket = map[uint8]func() ClientPacket{
		ClientLoginChallengeOpCode: func() ClientPacket { return new(ClientLoginChallenge) },
		ClientLoginProofOpCode:     func() ClientPacket { return new(ClientLoginProof) },
	}
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
	// to the client. It takes as input some session information (which
	// depends on the type of session).
	Handle(session *Session) ([]ServerPacket, error)
}

func padBigIntBytes(data []byte, nBytes int) []byte {
	if len(data) > nBytes {
		return data[:nBytes]
	}

	currSize := len(data)
	for i := 0; i < nBytes-currSize; i++ {
		data = append(data, '\x00')
	}

	return data
}

// ReadClientPacket will read the next available client packet from the connection
// and return it.
func ReadClientPacket(buffer *bufio.Reader) (ClientPacket, error) {
	// First, read the OpCode.
	opCode, err := buffer.ReadByte()
	if err != nil {
		return nil, err
	}

	builder, ok := opCodeToPacket[opCode]
	if !ok {
		log.Printf("Unknown opcode %v", opCode)
		return nil, nil
	}

	log.Printf("Received OpCode %v", opCode)
	return builder(), nil
}
