package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"

	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientSetActiveMover is sent from the client periodically.
type ClientSetActiveMover struct {
	GUID interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientSetActiveMover) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientSetActiveMover) OpCode() static.OpCode {
	return static.OpCodeClientSetActiveMover
}
