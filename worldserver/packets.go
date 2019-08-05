package worldserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
)

var (
	opCodeToPacket = map[session.OpCode]func() session.ClientPacket{
		packet.OpCodeClientAuthSession: func() session.ClientPacket { return new(packet.ClientAuthSession) },
	}
)

func readHeader(buffer io.Reader) (session.OpCode, int, error) {
	headerData := make([]byte, 6)
	n, err := buffer.Read(headerData)
	if err != nil {
		return packet.OpCode(0), 0, fmt.Errorf("erorr while reading header: %v", err)
	}

	if n != len(headerData) {
		return packet.OpCode(0), 0, errors.New("short read when reading opcode data")
	}

	// If there is a session key in the state, then we need to decrypt.
	// TODO(jeshua): implement this.

	// In the world server, the length is the first 2 bytes in the pkt.
	length := int(binary.BigEndian.Uint16(headerData[:2]))
	opCode := packet.OpCode(binary.LittleEndian.Uint32(headerData[2:]))

	return opCode, length - 4, nil
}

func writeHeader(packetLen int, opCode session.OpCode) ([]byte, error) {
	lengthData := make([]byte, 2)
	opCodeData := make([]byte, 2)

	binary.BigEndian.PutUint16(lengthData, uint16(packetLen))
	binary.LittleEndian.PutUint16(opCodeData, uint16(opCode.Int()))

	// If there is a session key in the state, then we need to encrypt.
	// TODO(jeshua): implement this.
	return append(lengthData, opCodeData...), nil
}
