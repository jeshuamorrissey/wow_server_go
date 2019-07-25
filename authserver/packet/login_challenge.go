package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver/srp"
	"gitlab.com/jeshuamorrissey/mmo_server/common"
	"gitlab.com/jeshuamorrissey/mmo_server/database"
	"gitlab.com/jeshuamorrissey/mmo_server/session"
)

// ClientLoginChallenge encodes information about a new connection to the
// login server.
type ClientLoginChallenge struct {
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
func (pkt *ClientLoginChallenge) Read(buffer io.Reader) error {
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

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginChallenge) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(ServerLoginChallengeOpCode))
	buffer.WriteByte(0) // unk1
	buffer.WriteByte(pkt.Error)

	if pkt.Error == 0 {
		buffer.Write(common.PadBigIntBytes(reverse(pkt.B.Bytes()), 32))
		buffer.WriteByte(1)
		buffer.WriteByte(srp.G)
		buffer.WriteByte(32)
		buffer.Write(reverse(srp.N().Bytes()))
		buffer.Write(common.PadBigIntBytes(reverse(pkt.Salt.Bytes()), 32))
		buffer.Write(common.PadBigIntBytes(reverse(pkt.SaltCRC.Bytes()), 16))
		buffer.WriteByte(0) // unk2
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerLoginChallenge) OpCode() session.OpCode {
	return ServerLoginChallengeOpCode
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginChallenge) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	state := stateBase.(*State)
	response := new(ServerLoginChallenge)

	// Get information from the session.
	account, err := database.GetAccount(string(pkt.AccountName))
	if err != nil {
		return nil, err
	}

	b, B := srp.GenerateEphemeral(&account.Verifier)
	state.PrivateEphemeral.Set(b)
	state.PublicEphemeral.Set(B)
	state.Account = account

	response.Error = 0
	response.B.Set(B)
	response.Salt.Set(&account.Salt)
	response.SaltCRC.SetInt64(0)

	return []session.ServerPacket{response}, nil
}
