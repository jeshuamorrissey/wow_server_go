package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/lib/util"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/session"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/srp"
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
func (pkt *ClientLoginChallenge) FromBytes(state *session.State, buffer io.Reader) error {
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

// OpCode gets the opcode of the packet.
func (*ClientLoginChallenge) OpCode() static.OpCode {
	return static.OpCodeLoginChallenge
}

// ServerLoginChallenge is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginChallenge struct {
	Error   static.LoginErrorCode
	B       big.Int
	Salt    big.Int
	SaltCRC big.Int
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginChallenge) ToBytes(state *session.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(0) // unk1
	buffer.WriteByte(uint8(pkt.Error))

	if pkt.Error == 0 {
		buffer.Write(util.PadBigIntBytes(util.ReverseBytes(pkt.B.Bytes()), 32))
		buffer.WriteByte(1)
		buffer.WriteByte(srp.G)
		buffer.WriteByte(32)
		buffer.Write(util.ReverseBytes(srp.N().Bytes()))
		buffer.Write(util.PadBigIntBytes(util.ReverseBytes(pkt.Salt.Bytes()), 32))
		buffer.Write(util.PadBigIntBytes(util.ReverseBytes(pkt.SaltCRC.Bytes()), 16))
		buffer.WriteByte(0) // unk2
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLoginChallenge) OpCode() static.OpCode {
	return static.OpCodeLoginChallenge
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginChallenge) Handle(state *session.State) ([]session.ServerPacket, error) {
	response := new(ServerLoginChallenge)
	response.Error = static.LoginOK

	// Validate the packet.
	gameName := strings.TrimRight(string(pkt.GameName[:]), "\x00")
	if gameName != static.SupportedGameName {
		response.Error = static.LoginFailed
	} else if pkt.Version != static.SupportedGameVersion || pkt.Build != static.SupportedGameBuild {
		response.Error = static.LoginBadVersion
	} else {
		for _, account := range state.Config.Accounts {
			if strings.ToLower(account.Name) == strings.ToLower(string(pkt.AccountName)) {
				state.Account = account
				break
			}
		}

		if state.Account == nil {
			response.Error = static.LoginUnknownAccount
		}
	}

	if response.Error == static.LoginOK {
		b, B := srp.GenerateEphemeralPair(state.Account.Verifier())
		state.PrivateEphemeral.Set(b)
		state.PublicEphemeral.Set(B)

		response.B.Set(B)
		response.Salt.Set(state.Account.Salt())
		response.SaltCRC.SetInt64(0)
	}

	return []session.ServerPacket{response}, nil
}
