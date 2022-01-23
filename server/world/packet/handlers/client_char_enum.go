package handlers

import (
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientCharEnum(pkt *packet.ClientCharEnum, state *system.State) ([]interfaces.ServerPacket, error) {
	charObj := state.OM.GetPlayer(state.Account.Character.GUID)
	equipment := make(map[static.EquipmentSlot]*packet.ItemSummary)
	for slot, itemGUID := range charObj.Equipment {
		item := state.OM.GetItem(itemGUID)
		equipment[slot] = &packet.ItemSummary{
			DisplayID:     item.GetTemplate().DisplayID,
			InventoryType: item.GetTemplate().InventoryType,
		}
	}

	var firstBagSummary *packet.ItemSummary = nil
	if charObj.FirstBag() != nil {
		firstBagSummary = &packet.ItemSummary{
			DisplayID:     charObj.FirstBag().GetTemplate().DisplayID,
			InventoryType: charObj.FirstBag().GetTemplate().InventoryType,
		}
	}

	return []interfaces.ServerPacket{&packet.ServerCharEnum{
		Characters: []*packet.CharacterSummary{
			{
				Name:        state.Account.Character.Name,
				GUID:        state.Account.Character.GUID,
				Race:        charObj.Race,
				Class:       charObj.Class,
				Gender:      charObj.Gender,
				SkinColor:   charObj.SkinColor,
				Face:        charObj.Face,
				HairStyle:   charObj.HairStyle,
				HairColor:   charObj.HairColor,
				Feature:     charObj.Feature,
				Level:       charObj.Level,
				ZoneID:      charObj.ZoneID,
				MapID:       charObj.MapID,
				Location:    charObj.Location,
				HasLoggedIn: state.Account.Character.HasLoggedIn,
				Flags:       state.Account.Character.Flags(),
				Equipment:   equipment,
				FirstBag:    firstBagSummary,
			},
		},
	}}, nil
}
