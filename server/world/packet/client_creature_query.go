package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientCreatureQuery is sent from the client periodically.
type ClientCreatureQuery struct {
	Entry uint32
	GUID  interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientCreatureQuery) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Entry)
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientCreatureQuery) OpCode() static.OpCode {
	return static.OpCodeClientCreatureQuery
}
