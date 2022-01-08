package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// ServerItemQuerySingleResponse is sent back in response to ClientPing.
type ServerItemQuerySingleResponse struct {
	Entry uint32
	Item  *static.Item
}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerItemQuerySingleResponse) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	if pkt.Item == nil {
		binary.Write(buffer, binary.LittleEndian, uint32(pkt.Entry|0x80000000))
		return buffer.Bytes(), nil
	}

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Entry))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Class))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.SubClass))

	buffer.WriteString(pkt.Item.Name)
	buffer.WriteByte('\x00')
	buffer.WriteByte('\x00') // Name2
	buffer.WriteByte('\x00') // Name3
	buffer.WriteByte('\x00') // Name4

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.DisplayID)) // DisplayID
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Quality))   // Quality
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Flags()))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.VendorBuyPrice))  // VendorBuyPrice
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.VendorSellPrice)) // VendorSellPrice
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.InventoryType))   // InventoryType
	binary.Write(buffer, binary.LittleEndian, uint32(0xFFFFFFFF))               // AllowableClassMask
	binary.Write(buffer, binary.LittleEndian, uint32(0xFFFFFFFF))               // AllowableRaceMask
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Level))           // ItemLevel
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredLevel
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredSkill
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredSkillRank
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredSpell
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredHonorRank
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredCityRank, deprecated
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredReputationFaction
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // RequiredReputationRank
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // VendorStackSize
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // MaxStackSize
	binary.Write(buffer, binary.LittleEndian, uint32(0))                        // ContainerSlots
	for i := 0; i < 10; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0)) // StatType
		binary.Write(buffer, binary.LittleEndian, int32(0))  // StatValue
	}

	for damageType, damage := range pkt.Item.Damages {
		binary.Write(buffer, binary.LittleEndian, float32(damage.Min))
		binary.Write(buffer, binary.LittleEndian, float32(damage.Max))
		binary.Write(buffer, binary.LittleEndian, uint32(damageType))
	}

	for i := 0; i < 5-len(pkt.Item.Damages); i++ {
		binary.Write(buffer, binary.LittleEndian, float32(0)) // DamageMin
		binary.Write(buffer, binary.LittleEndian, float32(0)) // DamageMax
		binary.Write(buffer, binary.LittleEndian, uint32(0))  // DamageType
	}

	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolPhysical]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolHoly]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolFire]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolNature]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolFrost]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolShadow]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.Resistances[static.SpellSchoolArcane]))
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.AttackRate.Seconds()*1000))
	binary.Write(buffer, binary.LittleEndian, uint32(0))  // RequiredAmmoType
	binary.Write(buffer, binary.LittleEndian, float32(0)) // RangedModRange

	for i := 0; i < 5; i++ {
		binary.Write(buffer, binary.LittleEndian, uint32(0))          // spell.ID
		binary.Write(buffer, binary.LittleEndian, uint32(0))          // spell.Trigger
		binary.Write(buffer, binary.LittleEndian, uint32(0))          // spell.Charges
		binary.Write(buffer, binary.LittleEndian, uint32(0xFFFFFFFF)) // spell.Cooldown
		binary.Write(buffer, binary.LittleEndian, uint32(0))          // spell.Category
		binary.Write(buffer, binary.LittleEndian, uint32(0xFFFFFFFF)) // spell.CategoryCooldown
	}
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Bonding
	buffer.WriteString(pkt.Item.Description)             // Description
	buffer.WriteByte('\x00')
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // PageText
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // Language
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // PageMaterial
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // StartQuest
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // LockID
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // Material
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.SheathType))    // SheathType
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // RandomProperty
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // Block
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // ItemSet
	binary.Write(buffer, binary.LittleEndian, uint32(pkt.Item.MaxDurability)) // MaxDurability
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // Area
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // Map
	binary.Write(buffer, binary.LittleEndian, uint32(0))                      // BagFamily

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerItemQuerySingleResponse) OpCode() static.OpCode {
	return static.OpCodeServerItemQuerySingleResponse
}
