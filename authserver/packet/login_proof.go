package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
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
func (pkt *ClientLoginProof) Read(buffer io.Reader) error {
	aBuffer := make([]byte, 32)
	buffer.Read(aBuffer)
	pkt.A.SetBytes(common.ReverseBytes(aBuffer))

	mBuffer := make([]byte, 20)
	buffer.Read(mBuffer)
	pkt.M.SetBytes(common.ReverseBytes(mBuffer))

	crcHashBuffer := make([]byte, 20)
	buffer.Read(crcHashBuffer)
	pkt.CRCHash.SetBytes(common.ReverseBytes(crcHashBuffer))

	binary.Read(buffer, binary.LittleEndian, &pkt.NumberOfKeys)
	return binary.Read(buffer, binary.LittleEndian, &pkt.SecurityFlags)
}

// ServerLoginProof is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginProof struct {
	Error LoginErrorCode
	Proof big.Int
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginProof) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	if pkt.Error == 0 {
		buffer.Write(common.PadBigIntBytes(common.ReverseBytes(pkt.Proof.Bytes()), 20))
		buffer.Write([]byte("\x00\x00\x00\x00")) // unk1
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerLoginProof) OpCode() session.OpCode {
	return OpCodeLoginProof
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginProof) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	state := stateBase.(*State)
	response := new(ServerLoginProof)

	K, M := srp.CalculateSessionKey(
		&pkt.A,
		&state.PublicEphemeral,
		&state.PrivateEphemeral,
		state.Account.Verifier(),
		state.Account.Salt(),
		state.Account.Name)

	if M.Cmp(&pkt.M) != 0 {
		response.Error = 4 // TODO(jeshua): make these constants
	} else {
		response.Error = 0
		response.Proof.Set(srp.CalculateServerProof(&pkt.A, M, K))

		state.AddLogField("account", state.Account.Name)

		state.Account.SessionKeyStr = new(string)
		*state.Account.SessionKeyStr = K.Text(16)
		stateBase.DB().Save(state.Account)
	}

	return []session.ServerPacket{response}, nil
}
