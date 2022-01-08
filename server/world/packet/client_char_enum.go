package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientCharEnum is sent from the client when first connecting.
type ClientCharEnum struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientCharEnum) FromBytes(buffer io.Reader) error {
	return nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharEnum) OpCode() static.OpCode {
	return static.OpCodeClientCharEnum
}
