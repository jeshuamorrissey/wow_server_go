package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientTutorialFlag is sent from the client periodically.
type ClientTutorialFlag struct {
	Flag uint32
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientTutorialFlag) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Flag)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientTutorialFlag) Handle(state *system.State) ([]system.ServerPacket, error) {
	state.Character.Tutorials[pkt.Flag] = true
	return nil, nil
}

// OpCode gets the opcode of the packet.
func (*ClientTutorialFlag) OpCode() system.OpCode {
	return system.OpCodeClientTutorialFlag
}
