package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerInitWorldStates is sent back in response to ClientPing.
type ServerInitWorldStates struct {
	Map    uint32
	Zone   uint32
	Blocks []WorldStateBlock
}

// WorldStateBlock represents some world state to initialize.
type WorldStateBlock struct {
	State uint32
	Value uint32
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerInitWorldStates) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Map))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Zone))
	binary.Write(buffer, binary.LittleEndian, uint16(len(pkt.Blocks)))
	for _, block := range pkt.Blocks {
		binary.Write(buffer, binary.LittleEndian, uint32(block.State))
		binary.Write(buffer, binary.LittleEndian, uint32(block.Value))
	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerInitWorldStates) OpCode() system.OpCode {
	return system.OpCodeServerInitWorldStates
}
