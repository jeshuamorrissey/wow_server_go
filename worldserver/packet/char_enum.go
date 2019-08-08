package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ClientCharEnum is sent from the client periodically.
type ClientCharEnum struct {
}

func (pkt *ClientCharEnum) Read(buffer io.Reader) error {
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientCharEnum) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	state := stateBase.(*State)
	response := new(ServerCharEnum)

	err := stateBase.DB().Where(&database.Character{AccountID: state.Account.ID, RealmID: state.Realm.ID}).Find(&response.Characters).Error
	if err != nil {
		return nil, err
	}

	return []session.ServerPacket{response}, nil
}

// ServerCharEnum is sent back in response to ClientPing.
type ServerCharEnum struct {
	Characters []database.Character
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerCharEnum) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(len(pkt.Characters))) // number of characters

	for _, char := range pkt.Characters {
		binary.Write(buffer, binary.LittleEndian, char.Object.GetHighGUID())
		binary.Write(buffer, binary.LittleEndian, uint32(char.Object.Model.ID))
		buffer.WriteString(char.Name)
		buffer.WriteByte(0)
		buffer.WriteByte(uint8(char.Object.Race))
		buffer.WriteByte(uint8(char.Object.Class))
		buffer.WriteByte(uint8(char.Object.Gender))
		buffer.WriteByte(char.Object.SkinColor)
		buffer.WriteByte(char.Object.Face)
		buffer.WriteByte(char.Object.HairStyle)
		buffer.WriteByte(char.Object.HairColor)
		buffer.WriteByte(char.Object.Feature)
		buffer.WriteByte(char.Object.Level)
		binary.Write(buffer, binary.LittleEndian, char.Object.ZoneID)
		binary.Write(buffer, binary.LittleEndian, char.Object.MapID)
		binary.Write(buffer, binary.LittleEndian, char.Object.X)
		binary.Write(buffer, binary.LittleEndian, char.Object.Y)
		binary.Write(buffer, binary.LittleEndian, char.Object.Z)

		// TODO(jeshua): implement the following fields.
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // GuildID
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // Flags
		binary.Write(buffer, binary.LittleEndian, uint8(0))  // FirstLogin
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetID
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetLevel
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetFamily

		equipmentMap := char.Object.EquipmentMap()
		for i := c.EquipmentSlotHead; i <= c.EquipmentSlotTabard; i++ {
			if item, ok := equipmentMap[i]; ok {
				binary.Write(buffer, binary.LittleEndian, uint32(item.Template().DisplayID))
				binary.Write(buffer, binary.LittleEndian, uint8(item.Template().InventoryType))
			} else {
				binary.Write(buffer, binary.LittleEndian, uint32(0)) // ItemDisplayID
				binary.Write(buffer, binary.LittleEndian, uint8(0))  // ItemInventoryType
			}
		}

		binary.Write(buffer, binary.LittleEndian, uint32(0)) // FirstBagDisplayID
		binary.Write(buffer, binary.LittleEndian, uint8(0))  // FirstBagInventoryType
	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerCharEnum) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerCharEnum)
}
