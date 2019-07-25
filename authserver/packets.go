package authserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"gitlab.com/jeshuamorrissey/mmo_server/common"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver/packet"
	"gitlab.com/jeshuamorrissey/mmo_server/session"
)

var (
	opCodeToPacket = map[session.OpCode]func() session.ClientPacket{
		packet.ClientLoginChallengeOpCode: func() session.ClientPacket { return new(packet.ClientLoginChallenge) },
		packet.ClientLoginProofOpCode:     func() session.ClientPacket { return new(packet.ClientLoginProof) },
		packet.ClientRealmlistOpCode:      func() session.ClientPacket { return new(packet.ClientRealmlist) },
	}
)

func readHeader(buffer io.Reader) (session.OpCode, int, error) {
	opCodeData := make([]byte, 1)
	n, err := buffer.Read(opCodeData)
	if err != nil {
		return 0, 0, fmt.Errorf("erorr while reading opcode: %v", err)
	}

	if n != 1 {
		return 0, 0, errors.New("short read when reading opcode data")
	}

	// In the auth server, the length is based on the packet type.
	opCode := session.OpCode(opCodeData[0])
	length := 0
	if opCode == packet.ClientLoginChallengeOpCode {
		lenData, err := common.ReadBytes(buffer, 3)
		if err != nil {
			return 0, 0, fmt.Errorf("error while reading header length: %v", err)
		}

		length = int(binary.LittleEndian.Uint16(lenData[1:]))
	} else if opCode == packet.ClientLoginProofOpCode {
		length = 74
	} else if opCode == packet.ClientRealmlistOpCode {
		length = 4
	}

	return opCode, length, nil
}
