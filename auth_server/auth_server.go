package main

import (
	"bufio"
	"fmt"
	"net"

	"gitlab.com/jeshuamorrissey/mmo_server/auth_server/packet"
	packet_common "gitlab.com/jeshuamorrissey/mmo_server/packet"
)

// ReadClientPacket will read the next available client packet from the connection
// and return it.
func ReadClientPacket(buffer *bufio.Reader) (packet_common.ClientPacket, error) {
	// First, read the OpCode.
	opCode, err := buffer.ReadByte()
	if err != nil {
		return nil, err
	}

	if opCode == packet.ClientLoginChallengeOpCode {
		return new(packet.ClientLoginChallenge), nil
	} else {
		return nil, fmt.Errorf("unknown opcode %v", opCode)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Printf("Error while opening port: %v", err)
		return
	}

	// Accept a connection.
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Error while receiving client connection: %v", err)
		return
	}

	// Make a new buffer to read from.
	buffer := bufio.NewReader(conn)

	// Receive packets from the connection.
	for {
		packet, err := ReadClientPacket(buffer)
		if err != nil {
			fmt.Printf("Error while reading client pakcet: %v", err)
			return
		}

		// Load and then handle the packet.
		packet.Read(buffer)
		response, err := packet.Handle()
		if err != nil {
			fmt.Printf("Error while handling packet: %v", err)
		}

		for pkt := range response {
			// send response
			fmt.Printf("%v", pkt)
		}
	}
}
