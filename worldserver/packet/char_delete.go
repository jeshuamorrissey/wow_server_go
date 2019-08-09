package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ClientCharDelete is sent from the client when deleting a character.
type ClientCharDelete struct {
	HighGUID c.HighGUID
	ID       uint32
}

func (pkt *ClientCharDelete) Read(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.HighGUID)
	binary.Read(buffer, binary.LittleEndian, &pkt.ID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharDelete) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	response := new(ServerCharDelete)
	response.Error = CharErrorCodeDeleteSuccess

	// Get the object.
	var charObj database.GameObjectPlayer
	err := stateBase.DB().Where("ID = ?", pkt.ID).First(&charObj).Error
	if err != nil {
		response.Error = CharErrorCodeDeleteFailed
		return []session.ServerPacket{response}, nil
	}

	var char database.Character
	err = stateBase.DB().Where("ID = ?", charObj.CharacterID).First(&char).Error
	if err != nil {
		response.Error = CharErrorCodeDeleteFailed
		return []session.ServerPacket{response}, nil
	}

	stateBase.DB().Unscoped().Delete(&char)

	return []session.ServerPacket{response}, nil
}

// ServerCharDelete is sent from the client when making a character.
type ServerCharDelete struct {
	Error CharErrorCode
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerCharDelete) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerCharDelete) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerCharDelete)
}
