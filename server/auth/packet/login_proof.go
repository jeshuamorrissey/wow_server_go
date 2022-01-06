package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/lib/util"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/session"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/srp"
)

// ClientLoginProof encodes proof that the client has the correct information.
type ClientLoginProof struct {
	A             big.Int
	M             big.Int
	CRCHash       big.Int
	NumberOfKeys  uint8
	SecurityFlags uint8
}

// Read will load a ClientLoginProof packet from a buffer.
// An error will be returned if at least one of the fields didn't load correctly.
func (pkt *ClientLoginProof) FromBytes(state *session.State, buffer io.Reader) error {
	aBuffer := make([]byte, 32)
	buffer.Read(aBuffer)
	pkt.A.SetBytes(util.ReverseBytes(aBuffer))

	mBuffer := make([]byte, 20)
	buffer.Read(mBuffer)
	pkt.M.SetBytes(util.ReverseBytes(mBuffer))

	crcHashBuffer := make([]byte, 20)
	buffer.Read(crcHashBuffer)
	pkt.CRCHash.SetBytes(util.ReverseBytes(crcHashBuffer))

	binary.Read(buffer, binary.LittleEndian, &pkt.NumberOfKeys)
	return binary.Read(buffer, binary.LittleEndian, &pkt.SecurityFlags)
}

// OpCode gets the opcode of the packet.
func (*ClientLoginProof) OpCode() static.OpCode {
	return static.OpCodeLoginProof
}

// ServerLoginProof is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginProof struct {
	Error static.LoginErrorCode
	Proof big.Int
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginProof) ToBytes(state *session.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	if pkt.Error == 0 {
		buffer.Write(util.PadBigIntBytes(util.ReverseBytes(pkt.Proof.Bytes()), 20))
		buffer.Write([]byte("\x00\x00\x00\x00")) // unk1
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerLoginProof) OpCode() static.OpCode {
	return static.OpCodeLoginProof
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginProof) Handle(state *session.State) ([]session.ServerPacket, error) {
	response := new(ServerLoginProof)

	K, M := srp.CalculateSessionKey(
		&pkt.A,
		state.PublicEphemeral,
		state.PrivateEphemeral,
		state.Account.Verifier(),
		state.Account.Salt(),
		state.Account.Name)

	if M.Cmp(&pkt.M) != 0 {
		response.Error = 4 // TODO(jeshua): make these constants
	} else {
		response.Error = 0
		response.Proof.Set(srp.CalculateServerProof(&pkt.A, M, K))

		state.AddLogField("account", state.Account.Name)

		state.Account.SessionKeyStr = K.Text(16)
	}

	return []session.ServerPacket{response}, nil
}
