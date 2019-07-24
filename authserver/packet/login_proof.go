package packet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math/big"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver/srp"
)

// OpCodes used by the AuthServer.
const (
	ClientLoginProofOpCode = 1
	ServerLoginProofOpCode = 1
)

// ClientLoginProof encodes proof that the client has the correct information.
type ClientLoginProof struct {
	A             big.Int
	M             big.Int
	CRCHash       big.Int
	NumberOfKeys  uint8
	SecurityFlags uint8
}

func reverse(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	return data
}

// Read will load a ClientLoginProof packet from a buffer.
// An error will be returned if at least one of the fields didn't load correctly.
func (pkt *ClientLoginProof) Read(buffer *bufio.Reader) error {
	aBuffer := make([]byte, 32)
	buffer.Read(aBuffer)
	pkt.A.SetBytes(reverse(aBuffer))

	mBuffer := make([]byte, 20)
	buffer.Read(mBuffer)
	pkt.M.SetBytes(reverse(mBuffer))

	crcHashBuffer := make([]byte, 20)
	buffer.Read(crcHashBuffer)
	pkt.CRCHash.SetBytes(reverse(crcHashBuffer))

	binary.Read(buffer, binary.LittleEndian, &pkt.NumberOfKeys)
	return binary.Read(buffer, binary.LittleEndian, &pkt.SecurityFlags)
}

// ServerLoginProof is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginProof struct {
	Error uint8
	Proof big.Int
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginProof) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(ServerLoginProofOpCode)
	buffer.WriteByte(pkt.Error)

	if pkt.Error == 0 {
		buffer.Write(padBigIntBytes(reverse(pkt.Proof.Bytes()), 20))
		buffer.Write([]byte("\x00\x00\x00\x00")) // unk1
	}

	return buffer.Bytes()
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginProof) Handle(session *Session) ([]ServerPacket, error) {
	response := new(ServerLoginProof)

	K, M := srp.CalculateSessionKey(
		&pkt.A,
		&session.PublicEphemeral,
		&session.PrivateEphemeral,
		&session.Account.Verifier,
		&session.Account.Salt,
		session.Account.Name)

	if M.Cmp(&pkt.M) != 0 {
		response.Error = 4 // TODO(jeshua): make these constants
	} else {
		response.Error = 0
		response.Proof.Set(srp.CalculateServerProof(&pkt.A, M, K))
	}

	return []ServerPacket{response}, nil
}
