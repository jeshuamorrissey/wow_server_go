package worldserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

var (
	opCodeToPacket = map[session.OpCode]func() session.ClientPacket{}
)

func readHeader(buffer io.Reader) (session.OpCode, int, error) {
	headerData := make([]byte, 6)
	n, err := buffer.Read(headerData)
	if err != nil {
		return 0, 0, fmt.Errorf("erorr while reading header: %v", err)
	}

	if n != len(headerData) {
		return 0, 0, errors.New("short read when reading opcode data")
	}

	// If there is a session key in the state, then we need to decrypt.
	// TODO(jeshua): implement this.

	// In the auth server, the length is based on the packet type.
	opCode := session.OpCode(binary.LittleEndian.Uint32(headerData[2:]))
	length := int(binary.BigEndian.Uint16(headerData[:2]))

	fmt.Printf("got opCode = %v, length = %v\n", opCode, length)

	return opCode, length - 4, nil
}
