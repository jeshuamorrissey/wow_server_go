package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

type ItemSummary struct {
	DisplayID     static.DisplayID
	InventoryType static.InventoryType
}

type CharacterSummary struct {
	Name        string
	GUID        interfaces.GUID
	Race        *static.Race
	Class       *static.Class
	Gender      static.Gender
	SkinColor   int
	Face        int
	HairStyle   int
	HairColor   int
	Feature     int
	Level       int
	ZoneID      int
	MapID       int
	Location    interfaces.Location
	HasLoggedIn bool
	Flags       uint32
	Equipment   map[static.EquipmentSlot]*ItemSummary
	FirstBag    *ItemSummary
}

// ServerCharEnum is sent back in response to ClientPing.
type ServerCharEnum struct {
	Characters []*CharacterSummary
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerCharEnum) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(uint8(len(pkt.Characters))) // number of characters

	for _, char := range pkt.Characters {
		binary.Write(buffer, binary.LittleEndian, char.GUID.Low())
		binary.Write(buffer, binary.LittleEndian, char.GUID.High())
		buffer.WriteString(char.Name)
		buffer.WriteByte(0)
		buffer.WriteByte(uint8(char.Race.ID))
		buffer.WriteByte(uint8(char.Class.ID))
		buffer.WriteByte(uint8(char.Gender))
		buffer.WriteByte(uint8(char.SkinColor))
		buffer.WriteByte(uint8(char.Face))
		buffer.WriteByte(uint8(char.HairStyle))
		buffer.WriteByte(uint8(char.HairColor))
		buffer.WriteByte(uint8(char.Feature))
		buffer.WriteByte(uint8(char.Level))
		binary.Write(buffer, binary.LittleEndian, uint32(char.ZoneID))
		binary.Write(buffer, binary.LittleEndian, uint32(char.MapID))
		binary.Write(buffer, binary.LittleEndian, float32(char.Location.X))
		binary.Write(buffer, binary.LittleEndian, float32(char.Location.Y))
		binary.Write(buffer, binary.LittleEndian, float32(char.Location.Z))

		// TODO(jeshua): implement the following fields with comments.
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // GuildID
		binary.Write(buffer, binary.LittleEndian, char.Flags)

		if !char.HasLoggedIn {
			buffer.WriteByte(1)
		} else {
			buffer.WriteByte(0)
		}

		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetID
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetLevel
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // PetFamily

		for slot := static.EquipmentSlotHead; slot <= static.EquipmentSlotTabard; slot++ {
			if itemSummary, ok := char.Equipment[slot]; ok {
				binary.Write(buffer, binary.LittleEndian, uint32(itemSummary.DisplayID))
				binary.Write(buffer, binary.LittleEndian, uint8(itemSummary.InventoryType))
			} else {
				binary.Write(buffer, binary.LittleEndian, uint32(0))
				binary.Write(buffer, binary.LittleEndian, uint8(0))
			}
		}

		if char.FirstBag != nil {
			binary.Write(buffer, binary.LittleEndian, uint32(char.FirstBag.DisplayID))
			binary.Write(buffer, binary.LittleEndian, uint8(char.FirstBag.InventoryType))
		} else {
			binary.Write(buffer, binary.LittleEndian, uint32(0))
			binary.Write(buffer, binary.LittleEndian, uint8(0))

		}

	}

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerCharEnum) OpCode() static.OpCode {
	return static.OpCodeServerCharEnum
}
