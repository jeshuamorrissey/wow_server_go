package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
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

// OpCode gets the opcode of the packet.
func (*ClientTutorialFlag) OpCode() static.OpCode {
	return static.OpCodeClientTutorialFlag
}
