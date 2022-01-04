package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientCharDelete is sent from the client when deleting a character.
type ClientCharDelete struct {
	HighGUID static.HighGUID
	ID       uint32
}

// FromBytes loads the packet from the given data.
func (pkt *ClientCharDelete) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.HighGUID)
	binary.Read(buffer, binary.LittleEndian, &pkt.ID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharDelete) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCharDelete)
	response.Error = static.CharErrorCodeDeleteFailed
	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharDelete) OpCode() static.OpCode {
	return static.OpCodeClientCharDelete
}
