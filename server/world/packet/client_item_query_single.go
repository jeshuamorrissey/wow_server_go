package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientItemQuerySingle is sent from the client periodically.
type ClientItemQuerySingle struct {
	Entry uint32
	GUID  interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientItemQuerySingle) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Entry)
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientItemQuerySingle) OpCode() static.OpCode {
	return static.OpCodeClientItemQuerySingle
}
