package session

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/lib/util"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
)

// Packet is a generic packet.
type Packet interface {
	// OpCode returns the opcode for the given packet as an int.
	OpCode() static.OpCode
}

// ServerPacket is a packet sent from this server to a client.
type ServerPacket interface {
	Packet

	// ToBytes writes the packet out to an array of bytes.
	ToBytes(*State) ([]byte, error)
}

// ClientPacket is a packet sent from the client to this server.
type ClientPacket interface {
	Packet

	// FromBytes reads the packet from a generic reader.
	FromBytes(*State, io.Reader) error

	// Handle the packet and return a list of server packets to send back
	// to the client.
	Handle(*State) ([]ServerPacket, error)
}

func readPacket(state *State, buffer io.Reader) (ClientPacket, static.OpCode, error) {
	opCode, length, err := readHeader(state, buffer)
	if err != nil {
		return nil, 0, err
	}

	data, err := util.ReadBytes(buffer, length)
	if err != nil {
		return nil, 0, err
	}

	builder, ok := state.opCodeToPacket[opCode]
	if !ok {
		return nil, opCode, nil
	}

	pkt := builder()
	pkt.FromBytes(state, bytes.NewReader(data))

	return pkt, opCode, nil
}

func readHeader(state *State, buffer io.Reader) (static.OpCode, int, error) {
	opCodeData := make([]byte, 1)
	n, err := buffer.Read(opCodeData)
	if err != nil {
		return static.OpCode(0), 0, fmt.Errorf("erorr while reading opcode: %v", err)
	}

	if n != 1 {
		return static.OpCode(0), 0, errors.New("short read when reading opcode data")
	}

	// In the auth server, the length is based on the packet type.
	opCode := static.OpCode(opCodeData[0])
	length := 0
	if opCode == static.OpCodeLoginChallenge {
		lenData, err := util.ReadBytes(buffer, 3)
		if err != nil {
			return static.OpCode(0), 0, fmt.Errorf("error while reading header length: %v", err)
		}

		length = int(binary.LittleEndian.Uint16(lenData[1:]))
	} else if opCode == static.OpCodeLoginProof {
		length = 74
	} else if opCode == static.OpCodeRealmlist {
		length = 4
	}

	return opCode, length, nil
}

func writeHeader(state *State, packetLen int, opCode static.OpCode) ([]byte, error) {
	return []byte{uint8(opCode.Int())}, nil
}
