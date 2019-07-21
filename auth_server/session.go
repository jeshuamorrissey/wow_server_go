package main

import (
	"bufio"
	"log"

	"gitlab.com/jeshuamorrissey/mmo_server/auth_server/packet"
)

// RunSession takes control of the thread and listens for packets and responds to them.
func RunSession(input *bufio.Reader, output *bufio.Writer) {
	for {
		packet, err := packet.ReadClientPacket(input)
		if err != nil {
			// This usually happens because of an EOF, so just terminate.
			log.Printf("Terminating connection: %v\n", err)
			return
		}

		// If the packet is nil, we didn't know how to read, so just ignore it.
		if packet == nil {
			continue
		}

		// Load and then handle the packet.
		packet.Read(input)
		response, err := packet.Handle()
		if err != nil {
			log.Printf("Error while handling packet: %v\n", err)
			return
		}

		for _, pkt := range response {
			log.Printf("--> %v (%v)\n", pkt.Bytes(), len(pkt.Bytes()))
			output.Write(pkt.Bytes())
			output.Flush()
		}
	}
}
