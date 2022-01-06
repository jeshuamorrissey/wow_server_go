package packet

import (
	"bufio"
	"encoding/binary"
	"io"
	"regexp"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/lib/util"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
	"github.com/jeshuamorrissey/wow_server_go/tools/new_game/initial_data"
)

func normalizeCharacterName(name string) string {
	re, _ := regexp.Compile("[^a-zA-Z]+")
	name = re.ReplaceAllString(name, "")
	return strings.Title(name)
}

func validateCharacterName(name string) static.CharErrorCode {
	if len(name) == 0 {
		return static.CharErrorCodeNameNoName
	} else if len(name) > static.MaxPlayerNameLength {
		return static.CharErrorCodeNameTooLong
	} else if len(name) < static.MinPlayerNameLength {
		return static.CharErrorCodeNameTooShort
	}

	return static.CharErrorCodeCreateSuccess
}

// ClientCharCreate is sent from the client when making a character.
type ClientCharCreate struct {
	Name      string
	Race      *static.Race
	Class     *static.Class
	Gender    static.Gender
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
	util.ReadBytes(buffer, 1) // OutfitID
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharCreate) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerCharCreate)
	response.Error = static.CharErrorCodeCreateSuccess

	// Check for invalid names.
	response.Error = validateCharacterName(pkt.Name)
	if response.Error != static.CharErrorCodeCreateSuccess {
		return []system.ServerPacket{response}, nil
	}

	// If the character already exists, fail.
	if state.Account.Character != nil {
		response.Error = static.CharErrorCodeCreateFailed
		return []system.ServerPacket{response}, nil
	}

	// Make the character.
	charObj, err := initial_data.NewCharacter(
		state.Config,
		pkt.Name,
		pkt.Race, pkt.Class, pkt.Gender,
		pkt.SkinColor, pkt.Face, pkt.HairStyle, pkt.HairColor, pkt.Feature)
	if err != nil {
		response.Error = static.CharErrorCodeCreateError
		return []system.ServerPacket{response}, nil
	}

	state.Account.Character = charObj

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientCharCreate) OpCode() static.OpCode {
	return static.OpCodeClientCharCreate
}
