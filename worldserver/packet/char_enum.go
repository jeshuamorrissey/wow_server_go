package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ClientCharEnum is sent from the client when first connecting.
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
	Characters []*database.Character
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerCharEnum) Bytes(stateBase session.State) []byte {
	state := stateBase.(*State)
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(len(pkt.Characters))) // number of characters

	for _, char := range pkt.Characters {
		charObj := char.Object(state.OM())
		binary.Write(buffer, binary.LittleEndian, charObj.GUID().High())
		binary.Write(buffer, binary.LittleEndian, charObj.GUID().Low())
		buffer.WriteString(char.Name)
		buffer.WriteByte(0)
		buffer.WriteByte(uint8(charObj.Race))
		buffer.WriteByte(uint8(charObj.Class))
		buffer.WriteByte(uint8(charObj.Gender))
		buffer.WriteByte(uint8(charObj.SkinColor))
		buffer.WriteByte(uint8(charObj.Face))
		buffer.WriteByte(uint8(charObj.HairStyle))
		buffer.WriteByte(uint8(charObj.HairColor))
		buffer.WriteByte(uint8(charObj.Feature))
		buffer.WriteByte(uint8(charObj.Level))
		binary.Write(buffer, binary.LittleEndian, uint32(charObj.ZoneID))
		binary.Write(buffer, binary.LittleEndian, uint32(charObj.MapID))
		binary.Write(buffer, binary.LittleEndian, float32(charObj.Location.X))
		binary.Write(buffer, binary.LittleEndian, float32(charObj.Location.Y))
		binary.Write(buffer, binary.LittleEndian, float32(charObj.Location.Z))

		// TODO(jeshua): implement the following fields with comments.
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // GuildID
		binary.Write(buffer, binary.LittleEndian, char.Flags())

		if char.LastLogin == nil {
			buffer.WriteByte(1)
		} else {
			buffer.WriteByte(0)
		}

		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetID
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetLevel
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetFamily

		for slot := c.EquipmentSlotHead; slot <= c.EquipmentSlotTabard; slot++ {
			if item, ok := charObj.Equipment[slot]; ok {
				binary.Write(buffer, binary.LittleEndian, uint32(item.Template().DisplayID))
				binary.Write(buffer, binary.LittleEndian, uint8(item.Template().InventoryType))
			} else {
				binary.Write(buffer, binary.LittleEndian, uint32(0))
				binary.Write(buffer, binary.LittleEndian, uint8(0))
			}
		}

		firstBag := charObj.FirstBag()
		if firstBag != nil {
			binary.Write(buffer, binary.LittleEndian, uint32(firstBag.Template().DisplayID))
			binary.Write(buffer, binary.LittleEndian, uint8(firstBag.Template().InventoryType))
		} else {
			binary.Write(buffer, binary.LittleEndian, uint32(0))
			binary.Write(buffer, binary.LittleEndian, uint8(0))

		}

	}

	return buffer.Bytes()
}

// OpCode gets the opcode of the packet.
func (*ServerCharEnum) OpCode() session.OpCode {
	return session.OpCode(OpCodeServerCharEnum)
}
