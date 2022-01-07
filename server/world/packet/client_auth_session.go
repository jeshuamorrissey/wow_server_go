package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientAuthSession is the initial message sent from the server
// to the client to start authorization.
type ClientAuthSession struct {
	BuildNumber      uint32
	AccountName      []byte
	ClientSeed       uint32
	ClientProof      [20]byte
	AddonSize        uint32
	AddonsCompressed []byte
}

// FromBytes reads a ClientAuthSession pcket from the byter buffer.
func (pkt *ClientAuthSession) FromBytes(state *system.State, buffer io.Reader) error {
	var unk uint32

	binary.Read(buffer, binary.LittleEndian, &pkt.BuildNumber)
	binary.Read(buffer, binary.LittleEndian, &unk)

	// Null-terminated account name
	var b byte
	binary.Read(buffer, binary.LittleEndian, &b)
	for b != '\x00' {
		pkt.AccountName = append(pkt.AccountName, b)
		binary.Read(buffer, binary.LittleEndian, &b)
	}

	binary.Read(buffer, binary.LittleEndian, &pkt.ClientSeed)
	binary.Read(buffer, binary.LittleEndian, &pkt.ClientProof)
	binary.Read(buffer, binary.LittleEndian, &pkt.AddonSize)

	pkt.AddonsCompressed = make([]byte, pkt.AddonSize)
	buffer.Read(pkt.AddonsCompressed)
	return nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientAuthSession) OpCode() static.OpCode {
	return static.OpCodeClientAuthSession
}
