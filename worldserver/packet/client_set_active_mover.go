package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientSetActiveMover is sent from the client periodically.
type ClientSetActiveMover struct {
	GUID object.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientSetActiveMover) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientSetActiveMover) Handle(state *system.State) ([]system.ServerPacket, error) {
	if pkt.GUID != state.Character.GUID() {
		state.Log.Errorf("Incorrect mover GUID: it is %v, but should be %v", pkt.GUID, state.Character.GUID())
	}

	return nil, nil
}

// OpCode gets the opcode of the packet.
func (*ClientSetActiveMover) OpCode() system.OpCode {
	return system.OpCodeClientSetActiveMover
}
