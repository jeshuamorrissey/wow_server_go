package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientCreatureQuery is sent from the client periodically.
type ClientCreatureQuery struct {
	Entry uint32
	GUID  interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientCreatureQuery) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Entry)
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCreatureQuery) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCreatureQueryResponse)

	response.Unit = nil
	if unit, ok := static.Units[int(pkt.Entry)]; ok {
		response.Unit = unit
		response.Entry = pkt.Entry
	} else if pkt.GUID != 0 && state.OM.Exists(pkt.GUID) {
		response.Unit = state.OM.GetUnit(pkt.GUID).Template()
		response.Entry = uint32(response.Unit.Entry)
	}

	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientCreatureQuery) OpCode() static.OpCode {
	return static.OpCodeClientCreatureQuery
}
