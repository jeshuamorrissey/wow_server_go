package packet

import (
	"bufio"
	"encoding/binary"
	"log"
	"math/big"

	"gitlab.com/jeshuamorrissey/mmo_server/auth_server/srp"
	"gitlab.com/jeshuamorrissey/mmo_server/packet"
)

// OpCodes used by the AuthServer.
const (
	ClientLoginChallengeOpCode = 0
	ServerLoginChallengeOpCode = 0
)

// ClientLoginChallenge encodes information about a new connection to the
// login server.
type ClientLoginChallenge struct {
	Error          uint8
	GameName       [4]byte
	Version        [3]uint8
	Build          uint16
	Platform       [4]byte
	OS             [4]byte
	Locale         [4]byte
	TimezoneOffset uint32
	IPAddress      uint32
	AccountName    []byte
}

// Read will load a ClientLoginChallenge packet from a buffer.
// An error will be returned if at least one of the fields didn't load correctly.
func (pkt *ClientLoginChallenge) Read(buffer *bufio.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Error)
	buffer.Read(make([]byte, 2)) // unused: packet length
	binary.Read(buffer, binary.LittleEndian, &pkt.GameName)
	binary.Read(buffer, binary.LittleEndian, &pkt.Version)
	binary.Read(buffer, binary.LittleEndian, &pkt.Build)
	binary.Read(buffer, binary.LittleEndian, &pkt.Platform)
	binary.Read(buffer, binary.LittleEndian, &pkt.OS)
	binary.Read(buffer, binary.LittleEndian, &pkt.Locale)
	binary.Read(buffer, binary.LittleEndian, &pkt.TimezoneOffset)
	binary.Read(buffer, binary.BigEndian, &pkt.IPAddress)

	var accountNameLen uint8
	binary.Read(buffer, binary.LittleEndian, &accountNameLen)

	pkt.AccountName = make([]byte, accountNameLen)
	return binary.Read(buffer, binary.LittleEndian, &pkt.AccountName)
}

// ServerLoginChallenge is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginChallenge struct {
	Error   uint8
	B       big.Int
	Salt    big.Int
	SaltCRC big.Int
}

// Write writes the data stored in the ServerLoginChallenge packet to a byte buffer.
func (pkt *ServerLoginChallenge) Write(buffer *bufio.Writer) error {
	buffer.WriteByte(ServerLoginChallengeOpCode)
	buffer.WriteByte(0) // unk1
	buffer.WriteByte(pkt.Error)

	if pkt.Error == 0 {
		buffer.Write(pkt.B.Bytes())
		buffer.WriteByte(1)
		buffer.WriteByte(srp.G)
		buffer.WriteByte(32)
		buffer.Write(srp.N().Bytes())
		buffer.Write(pkt.Salt.Bytes())
		buffer.Write(pkt.SaltCRC.Bytes())
		buffer.WriteByte(0) // unk2
	}

	return nil
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginChallenge) Handle() ([]packet.ServerPacket, error) {
	response := new(ServerLoginChallenge)

	log.Println(string(pkt.AccountName))

	return []packet.ServerPacket{response}, nil
}
