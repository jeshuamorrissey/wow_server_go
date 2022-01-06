package session

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

// Session represents a single user's connection to the server,
type Session struct {
	// A logger to write to.
	log *logrus.Entry

	// The I/O for this session. Usually will be socket conns.
	input  io.Reader
	output io.Writer

	// The state with some dynamic information.
	state *State
}

// NewSession makes a new session and returns it.
func NewSession(log *logrus.Entry, input io.Reader, output io.Writer, state *State) *Session {
	return &Session{
		log:    log,
		input:  input,
		output: output,
		state:  state,
	}
}

// SendPacket will send a packet back to the given output.
func (s *Session) SendPacket(pkt ServerPacket) error {
	s.log.Tracef("--> %v", pkt.OpCode())

	// Write the header.
	pktData, err := pkt.ToBytes(s.state)
	if err != nil {
		return err
	}

	header, err := writeHeader(s.state, len(pktData), pkt.OpCode())
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
		packet, opCode, err := readPacket(s.state, s.input)
		if err != nil {
			// This usually happens because of an EOF, so just terminate.
			s.log.Warnf("Terminating connection: %v\n", err)
			return
		}

		// If the packet is nil, we didn't know how to read, so just ignore it.
		if packet == nil {
			s.log.Warnf("<-- %v [UNHANDLED]", opCode)
			continue
		}

		s.log.Tracef("<-- %v", opCode)

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
