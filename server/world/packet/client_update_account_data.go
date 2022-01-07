package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientUpdateAccountData is sent from the client periodically.
type ClientUpdateAccountData struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientUpdateAccountData) FromBytes(state *system.State, buffer io.Reader) error {
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientUpdateAccountData) OpCode() static.OpCode {
	return static.OpCodeClientUpdateAccountData
}
