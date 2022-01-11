package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientAttackStop is sent from the client periodically.
type ClientAttackStop struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientAttackStop) FromBytes(buffer io.Reader) error {
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientAttackStop) OpCode() static.OpCode {
	return static.OpCodeClientAttackstop
}
