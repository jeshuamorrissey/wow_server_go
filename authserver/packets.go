package authserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

var (
	opCodeToPacket = map[session.OpCode]func() session.ClientPacket{
		packet.OpCodeLoginChallenge: func() session.ClientPacket { return new(packet.ClientLoginChallenge) },
		packet.OpCodeLoginProof:     func() session.ClientPacket { return new(packet.ClientLoginProof) },
		packet.OpCodeRealmlist:      func() session.ClientPacket { return new(packet.ClientRealmlist) },
	}
)

func readHeader(state session.State, buffer io.Reader) (session.OpCode, int, error) {
	opCodeData := make([]byte, 1)
	n, err := buffer.Read(opCodeData)
	if err != nil {
		return packet.OpCode(0), 0, fmt.Errorf("erorr while reading opcode: %v", err)
	}

	if n != 1 {
		return packet.OpCode(0), 0, errors.New("short read when reading opcode data")
	}

	// In the auth server, the length is based on the packet type.
	opCode := packet.OpCode(opCodeData[0])
	length := 0
	if opCode == packet.OpCodeLoginChallenge {
		lenData, err := common.ReadBytes(buffer, 3)
		if err != nil {
			return packet.OpCode(0), 0, fmt.Errorf("error while reading header length: %v", err)
		}

		length = int(binary.LittleEndian.Uint16(lenData[1:]))
	} else if opCode == packet.OpCodeLoginProof {
		length = 74
	} else if opCode == packet.OpCodeRealmlist {
		length = 4
	}

	return opCode, length, nil
}

func writeHeader(stateBase session.State, packetLen int, opCode session.OpCode) ([]byte, error) {
	return []byte{uint8(opCode.Int())}, nil
}
