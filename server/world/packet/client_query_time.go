package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientQueryTime is sent from the client periodically.
type ClientQueryTime struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientQueryTime) FromBytes(state *system.State, buffer io.Reader) error {
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientQueryTime) OpCode() static.OpCode {
	return static.OpCodeClientQueryTime
}
