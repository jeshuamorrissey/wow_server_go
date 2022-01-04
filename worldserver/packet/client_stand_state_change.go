package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientStandStateChange is sent from the client periodically.
type ClientStandStateChange struct {
	State static.StandState
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientStandStateChange) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.State)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientStandStateChange) Handle(state *system.State) ([]system.ServerPacket, error) {
	state.Character.StandState = pkt.State

	response := new(ServerStandStateUpdate)
	response.State = pkt.State

	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientStandStateChange) OpCode() static.OpCode {
	return static.OpCodeClientStandstatechange
}
