package session

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/config"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
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
	// Config should return a reference to the world config object.
	Config() *config.Config

	// Log returns a reference to the logger to use.
	Log() *logrus.Entry

	// AddLogField adds a new field to the log entry this state should
	// use.
	AddLogField(string, interface{})

	SetSession(*Session)
	Session() *Session
}

// Session management utility.
type Session struct {
	// Function which will read the header of the packet and return the OpCode +
	// the number of bytes that make up the packet.
	readHeader func(State, io.Reader) (OpCode, int, error)

	// Function which will write the header for the given packet and
	// return it. The arguments are the packet's byte length and the
	// packet's OpCode.
	writeHeader func(State, int, OpCode) ([]byte, error)

	// Function which will take as input an OpCode and return a valid ClientPacket.
	opCodeToPacket map[OpCode]func() ClientPacket

	// A logger to write to.
	log *logrus.Entry

	// The I/O for this session. Usually will be socket conns.
	input  io.Reader
	output io.Writer

	// Update object field cache.
	updateFieldCache map[interfaces.GUID]map[static.UpdateField]interface{}

	state State
}

// NewSession makes a new session and returns it.
func NewSession(
	readHeader func(State, io.Reader) (OpCode, int, error),
	writeHeader func(State, int, OpCode) ([]byte, error),
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
	opCode, length, err := s.readHeader(s.state, buffer)
	if err != nil {
		return nil, err
	}

	data, err := common.ReadBytes(buffer, length)
	if err != nil {
		return nil, err
	}

	builder, ok := s.opCodeToPacket[opCode]
	if !ok {
		s.log.Warnf("<-- %v [UNHANDLED]", opCode.String())
		return nil, nil
	}

	pkt := builder()
	pkt.Read(bytes.NewReader(data))

	s.log.Tracef("<-- %v", opCode.String())

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
	pktData := pkt.Bytes(s.state)
	header, err := s.writeHeader(s.state, len(pktData), pkt.OpCode())
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
