package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerCreatureQueryResponse is sent back in response to ClientPing.
type ServerCreatureQueryResponse struct {
	Entry uint32
	Unit  *static.Unit
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCreatureQueryResponse) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	if pkt.Unit == nil {
		binary.Write(buffer, binary.LittleEndian, uint32(pkt.Entry|0x80000000))
		return buffer.Bytes(), nil
	}

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Entry))
	buffer.WriteString(pkt.Unit.Name)
	buffer.WriteByte('\x00')
	buffer.WriteByte('\x00') // name2
	buffer.WriteByte('\x00') // name3
	buffer.WriteByte('\x00') // name4
	buffer.WriteString(pkt.Unit.SubName)
	buffer.WriteByte('\x00')

	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // CreatureTypeFlags
	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // CreatureType
	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // CreatureFamily
	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // CreatureRank
	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // unk
	binary.Write(buffer, binary.LittleEndian, uint32(0))                  // PetSpellDataID
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Unit.DisplayID)) // DisplayID
	binary.Write(buffer, binary.LittleEndian, uint8(0))                   // IsCivilian
	binary.Write(buffer, binary.LittleEndian, uint8(0))                   // IsRacialLeader

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCreatureQueryResponse) OpCode() static.OpCode {
	return static.OpCodeServerCreatureQueryResponse
}
