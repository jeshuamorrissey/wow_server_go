package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ClientStandStateChange is sent from the client periodically.
type ClientStandStateChange struct {
	State static.StandState
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientStandStateChange) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.State)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientStandStateChange) OpCode() static.OpCode {
	return static.OpCodeClientStandstatechange
}
