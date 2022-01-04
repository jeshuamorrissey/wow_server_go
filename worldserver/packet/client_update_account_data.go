package packet

import (
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientUpdateAccountData is sent from the client periodically.
type ClientUpdateAccountData struct{}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientUpdateAccountData) FromBytes(state *system.State, buffer io.Reader) error {
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientUpdateAccountData) Handle(state *system.State) ([]system.ServerPacket, error) {
	// Not implemented.
	return nil, nil
}

// OpCode gets the opcode of the packet.
func (*ClientUpdateAccountData) OpCode() static.OpCode {
	return static.OpCodeClientUpdateAccountData
}
