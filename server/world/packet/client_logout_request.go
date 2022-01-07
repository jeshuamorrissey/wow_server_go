package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientLogoutRequest is sent from the client periodically.
type ClientLogoutRequest struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientLogoutRequest) FromBytes(state *system.State, buffer io.Reader) error {
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientLogoutRequest) OpCode() static.OpCode {
	return static.OpCodeClientLogoutRequest
}
