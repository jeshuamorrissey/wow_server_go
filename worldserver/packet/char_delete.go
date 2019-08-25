package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// ClientCharDelete is sent from the client when deleting a character.
type ClientCharDelete struct {
	HighGUID c.HighGUID
	ID       uint32
}

// FromBytes loads the packet from the given data.
func (pkt *ClientCharDelete) FromBytes(state *State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.HighGUID)
	binary.Read(buffer, binary.LittleEndian, &pkt.ID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharDelete) Handle(state *State) ([]ServerPacket, error) {
	response := new(ServerCharDelete)
	response.Error = CharErrorCodeDeleteSuccess

	// Get the object.
	var char database.Character
	err := state.DB.Where("GUID = ?", uint64(pkt.HighGUID)<<32|uint64(pkt.ID)).First(&char).Error
	if err != nil {
		response.Error = CharErrorCodeDeleteFailed
		return []ServerPacket{response}, nil
	}

	state.DB.Unscoped().Delete(&char)

	return []ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharDelete) OpCode() OpCode {
	return OpCodeClientCharDelete
}

// ServerCharDelete is sent from the client when making a character.
type ServerCharDelete struct {
	Error CharErrorCode
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCharDelete) ToBytes(state *State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCharDelete) OpCode() OpCode {
	return OpCodeServerCharDelete
}
