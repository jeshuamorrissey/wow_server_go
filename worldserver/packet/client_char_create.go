package packet

import (
	"bufio"
	"encoding/binary"
	"io"
	"regexp"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
	"github.com/jinzhu/gorm"
)

func normalizeCharacterName(name string) string {
	re, _ := regexp.Compile("[^a-zA-Z]+")
	name = re.ReplaceAllString(name, "")
	return strings.Title(name)
}

func validateCharacterName(name string) c.CharErrorCode {
	if len(name) == 0 {
		return c.CharErrorCodeNameNoName
	} else if len(name) > c.MaxPlayerNameLength {
		return c.CharErrorCodeNameTooLong
	} else if len(name) < c.MinPlayerNameLength {
		return c.CharErrorCodeNameTooShort
	}

	return c.CharErrorCodeCreateSuccess
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
func (pkt *ClientCharCreate) FromBytes(state *system.State, bufferBase io.Reader) error {
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
func (pkt *ClientCharCreate) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCharCreate)
	response.Error = c.CharErrorCodeCreateSuccess

	// Check for invalid names.
	response.Error = validateCharacterName(pkt.Name)
	if response.Error != c.CharErrorCodeCreateSuccess {
		return []system.ServerPacket{response}, nil
	}

	// Check if name already exists.
	var char database.Character
	err := state.DB.Where(&database.Character{Name: pkt.Name}).First(&char).Error
	if err != gorm.ErrRecordNotFound {
		response.Error = c.CharErrorCodeCreateNameInUse
		return []system.ServerPacket{response}, nil
	}

	// Make the character.
	charObj, err := database.NewCharacter(
		state.OM, state.Account, state.Realm,
		pkt.Name,
		pkt.Race, pkt.Class, pkt.Gender,
		pkt.SkinColor, pkt.Face, pkt.HairStyle, pkt.HairColor, pkt.Feature)
	if err != nil {
		response.Error = c.CharErrorCodeCreateError
		return []system.ServerPacket{response}, nil
	}

	err = state.DB.Create(charObj).Error
	if err != nil {
		response.Error = c.CharErrorCodeCreateFailed
		return []system.ServerPacket{response}, nil
	}

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharCreate) OpCode() system.OpCode {
	return system.OpCodeClientCharCreate
}