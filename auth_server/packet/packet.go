package packet

import (
	"bufio"
	"log"

	"gitlab.com/jeshuamorrissey/mmo_server/packet"
)

var (
	opCodeToPacket = map[uint8]func() packet.ClientPacket{
		ClientLoginChallengeOpCode: func() packet.ClientPacket { return new(ClientLoginChallenge) },
		ClientLoginProofOpCode:     func() packet.ClientPacket { return new(ClientLoginProof) },
	}
)

func padBigIntBytes(data []byte, nBytes int) []byte {
	if len(data) > nBytes {
		return data[:nBytes]
	}

	currSize := len(data)
	for i := 0; i < nBytes-currSize; i++ {
		data = append(data, '\x00')
	}

	return data
}

// ReadClientPacket will read the next available client packet from the connection
// and return it.
func ReadClientPacket(buffer *bufio.Reader) (packet.ClientPacket, error) {
	// First, read the OpCode.
	opCode, err := buffer.ReadByte()
	if err != nil {
		return nil, err
	}

	builder, ok := opCodeToPacket[opCode]
	if !ok {
		log.Printf("Unknown opcode %v", opCode)
		return nil, nil
	}

	log.Printf("Received OpCode %v", opCode)
	return builder(), nil
}
