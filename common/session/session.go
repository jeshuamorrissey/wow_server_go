package session

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// OpCode is an integer type which is used to distinguish which packets are which.
type OpCode interface {
	Int() int
	String() string
}

// State is a generic interface which represents some data that needs to be stored
// with the session. This state will vary depending on the server.
type State interface {
	// DB should return a reference to the DB to use in this handler.
	DB() *gorm.DB

	// Log returns a reference to the logger to use.
	Log() *logrus.Entry

	// AddLogField adds a new field to the log entry this state should
	// use.
	AddLogField(string, interface{})
}

// Session management utility.
type Session struct {
	// Function which will read the header of the packet and return the OpCode +
	// the number of bytes that make up the packet.
	readHeader func(io.Reader) (OpCode, int, error)

	// Function which will write the header for the given packet and
	// return it. The arguments are the packet's byte length and the
	// packet's OpCode.
	writeHeader func(int, OpCode) ([]byte, error)

	// Function which will take as input an OpCode and return a valid ClientPacket.
	opCodeToPacket map[OpCode]func() ClientPacket

	// A logger to write to.
	log *logrus.Entry

	// The I/O for this session. Usually will be socket conns.
	input  io.Reader
	output io.Writer

	state State
}

// NewSession makes a new session and returns it.
func NewSession(
	readHeader func(io.Reader) (OpCode, int, error),
	writeHeader func(int, OpCode) ([]byte, error),
	opCodeToPacket map[OpCode]func() ClientPacket,
	log *logrus.Entry,
	input io.Reader,
	output io.Writer,
	state State) *Session {
	return &Session{
		readHeader:     readHeader,
		writeHeader:    writeHeader,
		opCodeToPacket: opCodeToPacket,
		log:            log,
		input:          input,
		output:         output,
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
		s.log.Warnf("<-- %v [UNHANDLED]", opCode.String())
		return nil, nil
	}

	s.log.Tracef("<-- %v", opCode.String())

	pkt := builder()
	pkt.Read(bytes.NewReader(data))

	return pkt, nil
}

// AddLogField will add a new field to the logger for this session.
func (s *Session) AddLogField(key string, value interface{}) {
	s.log = s.log.WithField(key, value)
}

// SendPacket will send a packet back to the given output.
func (s *Session) SendPacket(pkt ServerPacket) error {
	s.log.Tracef("--> %v", pkt.OpCode().String())

	// Write the header.
	pktData := pkt.Bytes()
	header, err := s.writeHeader(len(pktData), pkt.OpCode())
	if err != nil {
		return err
	}

	// Send the data.
	toSend := append(header, pktData...)
	n, err := s.output.Write(toSend)
	if err != nil {
		return err
	}

	if n != len(toSend) {
		return fmt.Errorf("expected %v bytes to send, only sent %v", len(pktData), n)
	}

	return nil
}

// Run takes as input a data source (input) and data destination (output) and
// manages incoming packets, routing them to the handler and then sending the
// output to the appropriate place.
func (s *Session) Run() {
	for {
		packet, err := s.readPacket(s.input)
		if err != nil {
			// This usually happens because of an EOF, so just terminate.
			s.log.Warnf("Terminating connection: %v\n", err)
			return
		}

		// If the packet is nil, we didn't know how to read, so just ignore it.
		if packet == nil {
			continue
		}

		// Load and then handle the packet.
		response, err := packet.Handle(s.state)
		if err != nil {
			s.log.Warnf("Error while handling packet: %v\n", err)
			return
		}

		for _, pkt := range response {
			s.SendPacket(pkt)
		}
	}
}
