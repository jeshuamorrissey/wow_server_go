package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientCreatureQuery is sent from the client periodically.
type ClientCreatureQuery struct {
	Entry uint32
	GUID  object.GUID
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
	if unit, ok := dbc.Units[int(pkt.Entry)]; ok {
		response.Unit = unit
		response.Entry = pkt.Entry
	} else if pkt.GUID != 0 && state.OM.Exists(pkt.GUID) {
		response.Unit = state.OM.Get(pkt.GUID).(*object.Unit).Template()
		response.Entry = uint32(response.Unit.Entry)
	}

	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientCreatureQuery) OpCode() system.OpCode {
	return system.OpCodeClientCreatureQuery
}
