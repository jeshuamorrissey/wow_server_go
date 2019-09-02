package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientItemQuerySingle is sent from the client periodically.
type ClientItemQuerySingle struct {
	Entry uint32
	GUID  object.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientItemQuerySingle) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Entry)
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientItemQuerySingle) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerItemQuerySingleResponse)

	response.Entry = pkt.Entry
	response.Item = nil
	if item, ok := dbc.Items[int(pkt.Entry)]; ok {
		response.Item = item
	} else if pkt.GUID != 0 && state.OM.Exists(pkt.GUID) {
		itemObjGeneric := state.OM.Get(pkt.GUID)
		if itemObjGeneric != nil {
			switch itemObj := itemObjGeneric.(type) {
			case *object.Item:
				response.Item = itemObj.Template()
			case *object.Container:
				response.Item = itemObj.Template()
			}
		}
	}

	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientItemQuerySingle) OpCode() system.OpCode {
	return system.OpCodeClientItemQuerySingle
}
