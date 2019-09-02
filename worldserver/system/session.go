package system

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Session represents a single client connection to the World Server.
type Session struct {
	// The input/output ports for this session.
	inputLock  sync.Mutex
	input      io.Reader
	outputLock sync.Mutex
	output     io.Writer

	// A mapping of opCode --> callback to create the client packet.
	opCodeToPacket map[OpCode]func() ClientPacket

	// State that is to be passed to each handler.
	state *State

	// Private data required to encrypt/decrypt packet headers.
	sendI, sendJ, recvI, recvJ uint8
}

// NewSession constructs a new session object and returns it.
func NewSession(
	input io.Reader,
	output io.Writer,
	opCodeToPacket map[OpCode]func() ClientPacket,
	db *gorm.DB,
	objectManager *object.Manager,
	log *logrus.Entry,
	realm *database.Realm,
	updater *Updater,
) *Session {
	state := &State{
		Log: log,

		DB:      db,
		OM:      objectManager,
		Updater: updater,

		Realm:     realm,
		Account:   nil,
		Character: nil,
	}

	session := &Session{
		input:          input,
		output:         output,
		opCodeToPacket: opCodeToPacket,
		state:          state,
	}

	state.Session = session
	return session
}

// Send sends a single packet to this session's client.
func (s *Session) Send(pkt ServerPacket) error {
	s.outputLock.Lock()
	defer s.outputLock.Unlock()

	opCode := pkt.OpCode()
	pktData, err := pkt.ToBytes(s.state)
	if err != nil {
		return err
	}

	// TODO(jeshua): Get this working!
	// if opCode == OpCodeServerUpdateObject && len(pktData) > 100 {
	// 	opCode = OpCodeServerCompressedUpdateObject

	// 	var compressedPktData bytes.Buffer
	// 	writer, err := zlib.NewWriterLevel(&compressedPktData, 1)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	writer.Write(pktData)

	// 	pktData := make([]byte, 4, compressedPktData.Len()+4)
	// 	binary.LittleEndian.PutUint32(pktData, uint32(compressedPktData.Len()))
	// 	pktData = append(pktData, compressedPktData.Bytes()...)
	// }

	header, err := s.makeHeader(len(pktData), opCode)
	if err != nil {
		return err
	}

	s.state.Log.Tracef("--> %s", opCode)
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

// Run starts the given session object. This should be called as a separate
// goroutine. The session will end when the user disconnects.
func (s *Session) Run() {
	for {
		pkt, err := s.readPacket()
		if err != nil {
			s.state.Log.Warnf("Terminating connection: %v\n", err)
			s.state.Updater.Logout(s.state.Character.GUID())
			return
		}

		// If the packet is nil, we don't know how to read it yet.
		if pkt == nil {
			continue
		}

		// Load and then handle the packet.
		responses, err := pkt.Handle(s.state)
		if err != nil {
			s.state.Log.Warnf("Error while handling packet %s: %v", pkt.OpCode(), err)
			continue
		}

		for _, response := range responses {
			s.Send(response)
		}
	}
}

func (s *Session) readPacket() (ClientPacket, error) {
	s.inputLock.Lock()
	defer s.inputLock.Unlock()

	opCode, length, err := s.readHeader()
	if err != nil {
		return nil, err
	}

	data, err := common.ReadBytes(s.input, length)
	if err != nil {
		return nil, err
	}

	builder, ok := s.opCodeToPacket[opCode]
	if !ok {
		s.state.Log.Warnf("<-- %v [UNHANDLED]", opCode.String())
		return nil, nil
	}

	pkt := builder()
	pkt.FromBytes(s.state, bytes.NewReader(data))

	s.state.Log.Tracef("<-- %s", opCode)

	return pkt, nil
}

func (s *Session) readHeader() (OpCode, int, error) {
	headerData := make([]byte, 6)
	n, err := s.input.Read(headerData)
	if err != nil {
		return OpCode(0), 0, fmt.Errorf("erorr while reading header: %v", err)
	}

	if n != len(headerData) {
		return OpCode(0), 0, errors.New("short read when reading opcode data")
	}

	// If there is a session key in the state, then we need to decrypt.
	if s.state.Account != nil && s.state.Account.SessionKey() != nil {
		sessionKeyBytes := common.ReverseBytes(s.state.Account.SessionKey().Bytes())

		for i := 0; i < 6; i++ {
			s.recvI %= uint8(len(sessionKeyBytes))
			x := (headerData[i] - s.recvJ) ^ sessionKeyBytes[s.recvI]
			s.recvI++
			s.recvJ = headerData[i]
			headerData[i] = x
		}
	}

	// In the world server, the length is the first 2 bytes in the pkt.
	length := int(binary.BigEndian.Uint16(headerData[:2]))
	opCode := OpCode(binary.LittleEndian.Uint32(headerData[2:]))

	return opCode, length - 4, nil
}

func (s *Session) makeHeader(packetLen int, opCode OpCode) ([]byte, error) {
	lengthData := make([]byte, 2)
	opCodeData := make([]byte, 2)

	binary.BigEndian.PutUint16(lengthData, uint16(packetLen)+2)
	binary.LittleEndian.PutUint16(opCodeData, uint16(opCode.Int()))

	header := make([]byte, 0, 4)
	header = append(header, lengthData...)
	header = append(header, opCodeData...)

	// If there is a session key in the state, then we need to encrypt.
	if s.state.Account != nil && s.state.Account.SessionKey() != nil {
		sessionKeyBytes := common.ReverseBytes(s.state.Account.SessionKey().Bytes())

		for i := 0; i < 4; i++ {
			s.sendI %= uint8(len(sessionKeyBytes))
			x := (header[i] ^ sessionKeyBytes[s.sendI]) + s.sendJ
			s.sendI++

			header[i] = x
			s.sendJ = x
		}
	}

	return header, nil
}
