package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientPing is sent from the client periodically.
type ClientPing struct {
	Ping    uint32
	Latency uint32
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientPing) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Ping)
	binary.Read(buffer, binary.LittleEndian, &pkt.Latency)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientPing) OpCode() static.OpCode {
	return static.OpCodeClientPing
}
