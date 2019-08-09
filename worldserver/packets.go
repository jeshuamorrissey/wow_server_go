package worldserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
)

var (
	opCodeToPacket = map[session.OpCode]func() session.ClientPacket{
		packet.OpCodeClientAuthSession: func() session.ClientPacket { return new(packet.ClientAuthSession) },
		packet.OpCodeClientCharCreate:  func() session.ClientPacket { return new(packet.ClientCharCreate) },
		packet.OpCodeClientCharEnum:    func() session.ClientPacket { return new(packet.ClientCharEnum) },
		packet.OpCodeClientPing:        func() session.ClientPacket { return new(packet.ClientPing) },
	}
)

func readHeader(stateBase session.State, buffer io.Reader) (session.OpCode, int, error) {
	state := stateBase.(*packet.State)

	headerData := make([]byte, 6)
	n, err := buffer.Read(headerData)
	if err != nil {
		return packet.OpCode(0), 0, fmt.Errorf("erorr while reading header: %v", err)
	}

	if n != len(headerData) {
		return packet.OpCode(0), 0, errors.New("short read when reading opcode data")
	}

	// If there is a session key in the state, then we need to decrypt.
	if state.Account.SessionKey() != nil {
		sessionKeyBytes := common.ReverseBytes(state.Account.SessionKey().Bytes())

		for i := 0; i < 6; i++ {
			state.RecvI %= uint8(len(sessionKeyBytes))
			x := (headerData[i] - state.RecvJ) ^ sessionKeyBytes[state.RecvI]
			state.RecvI++
			state.RecvJ = headerData[i]
			headerData[i] = x
		}
	}

	// In the world server, the length is the first 2 bytes in the pkt.
	length := int(binary.BigEndian.Uint16(headerData[:2]))
	opCode := packet.OpCode(binary.LittleEndian.Uint32(headerData[2:]))

	return opCode, length - 4, nil
}

func writeHeader(stateBase session.State, packetLen int, opCode session.OpCode) ([]byte, error) {
	state := stateBase.(*packet.State)

	lengthData := make([]byte, 2)
	opCodeData := make([]byte, 2)

	binary.BigEndian.PutUint16(lengthData, uint16(packetLen)+2)
	binary.LittleEndian.PutUint16(opCodeData, uint16(opCode.Int()))

	header := make([]byte, 0)
	header = append(header, lengthData...)
	header = append(header, opCodeData...)

	// If there is a session key in the state, then we need to encrypt.
	if state.Account.SessionKey() != nil {
		sessionKeyBytes := common.ReverseBytes(state.Account.SessionKey().Bytes())

		for i := 0; i < 4; i++ {
			state.SendI %= uint8(len(sessionKeyBytes))
			x := (header[i] ^ sessionKeyBytes[state.SendI]) + state.SendJ
			state.SendI++

			header[i] = x
			state.SendJ = x
		}
	}

	// TODO(jeshua): implement this.
	return header, nil
}
