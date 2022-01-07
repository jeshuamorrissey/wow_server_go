package handlers

import (
	"regexp"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
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

// Handle will ensure that the given account exists.
func HandleClientCharCreate(pkt *packet.ClientCharCreate, state *system.State) ([]system.ServerPacket, error) {
	response := new(packet.ServerCharCreate)
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
