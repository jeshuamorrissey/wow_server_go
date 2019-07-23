package packet

import (
	"bufio"
	"log"
	"math/big"

	"gitlab.com/jeshuamorrissey/mmo_server/database"
	"go.mongodb.org/mongo-driver/mongo"
)

// Session contains information required to maintain a user session. One of these
// structures will be created for each session.
type Session struct {
	Database *mongo.Database

	PublicEphemeral  big.Int
	PrivateEphemeral big.Int

	Account *database.Account
}

// RunSession takes control of the thread and listens for packets and responds to them.
func RunSession(database *mongo.Database, input *bufio.Reader, output *bufio.Writer) {
	session := Session{
		Database: database,

		PublicEphemeral:  big.Int{},
		PrivateEphemeral: big.Int{},

		Account: nil,
	}

	for {
		packet, err := ReadClientPacket(input)
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
		response, err := packet.Handle(&session)
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
