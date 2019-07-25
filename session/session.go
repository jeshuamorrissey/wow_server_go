package session

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

// OpCode is an integer type which is used to distinguish which packets are which.
type OpCode int

// State is a generic interface which represents some data that needs to be stored
// with the session. This state will vary depending on the server.
type State interface{}

// Session management utility.
type Session struct {
	// Function which will read the header of the packet and return the OpCode +
	// the number of bytes that make up the packet.
	readHeader func(io.Reader) (OpCode, int, error)

	// Function which will take as input an OpCode and return a valid ClientPacket.
	opCodeToPacket map[OpCode]func() ClientPacket

	// Function which will take as input an OpCode and return a string name.
	opCodeName func(OpCode) string

	state State
}

// NewSession makes a new session and returns it.
func NewSession(
	readHeader func(io.Reader) (OpCode, int, error),
	opCodeToPacket map[OpCode]func() ClientPacket,
	opCodeName func(OpCode) string,
	state State) *Session {
	return &Session{
		readHeader:     readHeader,
		opCodeToPacket: opCodeToPacket,
		opCodeName:     opCodeName,
		state:          state,
	}
}

func (s *Session) readPacket(buffer io.Reader) (ClientPacket, error) {
	opCode, length, err := s.readHeader(buffer)
	if err != nil {
		return nil, err
	}

	// Read all required bytes from the buffer. If less were read than expected,
	// or an EOF was not found, then error.
	data := make([]byte, length)
	n, err := buffer.Read(data)

	// If this is a valid packet, then the length should match exactly.
	if n != length {
		return nil, fmt.Errorf("short read: wanted %v bytes, only got %v bytes", length, n)
	}

	// We shouldn't see any error (even an EOF).
	if err != nil {
		return nil, fmt.Errorf("unknown error while reading packet data: %v", err)
	}

	builder, ok := s.opCodeToPacket[opCode]
	if !ok {
		log.Printf("Unhandled opcode %v", s.opCodeName(opCode))
		return nil, nil
	}

	log.Printf("<-- %v", s.opCodeName(opCode))

	pkt := builder()
	pkt.Read(bytes.NewReader(data))

	return pkt, nil
}

// Run takes as input a data source (input) and data destination (output) and
// manages incoming packets, routing them to the handler and then sending the
// output to the appropriate place.
func (s *Session) Run(input io.Reader, output io.Writer) {
	for {
		packet, err := s.readPacket(input)
		if err != nil {
			// This usually happens because of an EOF, so just terminate.
			log.Printf("Terminating connection: %v\n", err)
			return
		}

		// If the packet is nil, we didn't know how to read, so just ignore it.
		if packet == nil {
			continue
		}

		// Load and then handle the packet.
		response, err := packet.Handle(s.state)
		if err != nil {
			log.Printf("Error while handling packet: %v\n", err)
			return
		}

		for _, pkt := range response {
			log.Printf("--> %v", s.opCodeName(pkt.OpCode()))
			output.Write(pkt.Bytes())
		}
	}
}
