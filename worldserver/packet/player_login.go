package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
)

// ClientPlayerLogin is sent from the client periodically.
type ClientPlayerLogin struct {
	GUID object.GUID
}

func (pkt *ClientPlayerLogin) Read(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientPlayerLogin) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	state := stateBase.(*State)

	if !state.OM().Exists(pkt.GUID) {
		stateBase.Log().Errorf("Attempt to log in with unknown GUID %v!", pkt.GUID)
		return []session.ServerPacket{}, nil
	}

	player := state.OM().Get(pkt.GUID).(*object.Player)
	stateBase.Log().Infof("player = %v", player)

	state.OM().Register(player, func(updates map[c.UpdateType][]object.GameObject) {
		for updateType, objs := range updates {
			pkt := new(ServerUpdateObject)
			pkt.Updates = make([]ObjectUpdate, 0)

			for _, obj := range objs {
				update := &Update{
					updateType: updateType,
					Object:     obj,
					Victim:     nil,
					WorldTime:  0,
				}

				if obj.GUID() == player.GUID() {
					update.IsSelf = true
				}

				pkt.Updates = append(pkt.Updates, update)
			}

			state.Session().SendPacket(pkt)
		}
	})

	return []session.ServerPacket{
		&ServerLoginVerifyWorld{
			Character: player,
		},
		&ServerAccountDataTimes{},
		&ServerTutorialFlags{},
		&ServerInitWorldStates{
			Map:  uint32(player.MapID),
			Zone: uint32(player.ZoneID),
		},
	}, nil
}
