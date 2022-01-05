package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientNameQuery is sent from the client periodically.
type ClientNameQuery struct {
	GUID interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientNameQuery) FromBytes(state *system.State, buffer io.Reader) error {
	return binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
}

// Handle will ensure that the given account exists.
func (pkt *ClientNameQuery) Handle(state *system.State) ([]system.ServerPacket, error) {
	if !state.OM.Exists(pkt.GUID) {
		return nil, nil
	}

	response := new(ServerNameQueryResponse)
	response.Character = state.Account.Character
	return []system.ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientNameQuery) OpCode() static.OpCode {
	return static.OpCodeClientNameQuery
}
