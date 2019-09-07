package packet

import (
	"bytes"
	"encoding/binary"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerStandStateUpdate is sent back in response to ClientPing.
type ServerStandStateUpdate struct {
	State c.StandState
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerStandStateUpdate) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, pkt.State)

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerStandStateUpdate) OpCode() system.OpCode {
	return system.OpCodeServerStandstateUpdate
}
