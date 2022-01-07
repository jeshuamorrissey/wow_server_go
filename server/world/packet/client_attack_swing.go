package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientAttackSwing is sent from the client periodically.
type ClientAttackSwing struct {
	Target interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientAttackSwing) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Target)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientAttackSwing) OpCode() static.OpCode {
	return static.OpCodeClientAttackswing
}
