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

// Handle will ensure that the given account exists.
func (pkt *ClientLogoutRequest) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerLogoutResponse)

	// TODO: Actually implement this!
	response.Reason = 0
	response.InstantLogout = true

	if state.Character != nil {
		state.Updater.Logout(state.Character.GUID())
	}

	return []system.ServerPacket{
		response, new(ServerLogoutComplete),
	}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientLogoutRequest) OpCode() static.OpCode {
	return static.OpCodeClientLogoutRequest
}
