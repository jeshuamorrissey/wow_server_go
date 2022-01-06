package config

import "github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"

type CharacterSettings struct {
	HideHelm  bool `json:"hide_helm"`
	HideCloak bool `json:"hide_cloak"`
}

type Character struct {
	Name        string            `json:"name"`
	GUID        interfaces.GUID   `json:"guid"`
	HasLoggedIn bool              `json:"has_logged_in"`
	Settings    CharacterSettings `json:"settings"`
}

// Flags returns an set of flags based on the character's state.
func (char *Character) Flags() uint32 {
	var flags uint32
	// if char.Settings.HideHelm {
	// 	flags |= uint32(static.CharacterFlagHideHelm)
	// }

	// if char.Settings.HideCloak {
	// 	flags |= uint32(static.CharacterFlagHideCloak)
	// }

	// if char.Settings.IsGhost {
	// 	flags |= uint32(static.CharacterFlagGhost)
	// }

	return flags
}
