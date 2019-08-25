package packet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"regexp"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jinzhu/gorm"
)

func normalizeCharacterName(name string) string {
	re, _ := regexp.Compile("[^a-zA-Z]+")
	name = re.ReplaceAllString(name, "")
	return strings.Title(name)
}

func validateCharacterName(name string) CharErrorCode {
	if len(name) == 0 {
		return CharErrorCodeNameNoName
	} else if len(name) > c.MaxPlayerNameLength {
		return CharErrorCodeNameTooLong
	} else if len(name) < c.MinPlayerNameLength {
		return CharErrorCodeNameTooShort
	}

	return CharErrorCodeCreateSuccess
}

// ClientCharCreate is sent from the client when making a character.
type ClientCharCreate struct {
	Name      string
	Race      c.Race
	Class     c.Class
	Gender    c.Gender
	SkinColor uint8
	Face      uint8
	HairStyle uint8
	HairColor uint8
	Feature   uint8
}

// FromBytes loads the packet from the given data.
func (pkt *ClientCharCreate) FromBytes(state *State, bufferBase io.Reader) error {
	buffer := bufio.NewReader(bufferBase)
	pkt.Name, _ = buffer.ReadString('\x00')
	pkt.Name = normalizeCharacterName(pkt.Name)
	binary.Read(buffer, binary.LittleEndian, &pkt.Race)
	binary.Read(buffer, binary.LittleEndian, &pkt.Class)
	binary.Read(buffer, binary.LittleEndian, &pkt.Gender)
	binary.Read(buffer, binary.LittleEndian, &pkt.SkinColor)
	binary.Read(buffer, binary.LittleEndian, &pkt.Face)
	binary.Read(buffer, binary.LittleEndian, &pkt.HairStyle)
	binary.Read(buffer, binary.LittleEndian, &pkt.HairColor)
	binary.Read(buffer, binary.LittleEndian, &pkt.Feature)
	common.ReadBytes(buffer, 1) // OutfitID
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharCreate) Handle(state *State) ([]ServerPacket, error) {
	response := new(ServerCharCreate)
	response.Error = CharErrorCodeCreateSuccess

	// Check for invalid names.
	response.Error = validateCharacterName(pkt.Name)
	if response.Error != CharErrorCodeCreateSuccess {
		return []ServerPacket{response}, nil
	}

	// Check if name already exists.
	var char database.Character
	err := state.DB.Where(&database.Character{Name: pkt.Name}).First(&char).Error
	if err != gorm.ErrRecordNotFound {
		response.Error = CharErrorCodeCreateNameInUse
		return []ServerPacket{response}, nil
	}

	// Make the character.
	charObj, err := database.NewCharacter(
		state.OM, state.Account, state.Realm,
		pkt.Name,
		pkt.Race, pkt.Class, pkt.Gender,
		pkt.SkinColor, pkt.Face, pkt.HairStyle, pkt.HairColor, pkt.Feature)
	if err != nil {
		response.Error = CharErrorCodeCreateError
		return []ServerPacket{response}, nil
	}

	err = state.DB.Create(charObj).Error
	if err != nil {
		response.Error = CharErrorCodeCreateFailed
		return []ServerPacket{response}, nil
	}

	return []ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharCreate) OpCode() OpCode {
	return OpCodeClientCharCreate
}

// ServerCharCreate is sent from the client when making a character.
type ServerCharCreate struct {
	Error CharErrorCode
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCharCreate) ToBytes(state *State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(pkt.Error))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCharCreate) OpCode() OpCode {
	return OpCodeServerCharCreate
}
