package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerUpdateObject is the UPDATE_OBJECT packet.
type ServerUpdateObject struct {
	Updates []ObjectUpdate
}

// ObjectUpdate is a common parent class for different types of updates.
type ObjectUpdate interface {
	UpdateType() c.UpdateType
}

// OutOfRangeUpdate represents making a series of objects out of range.
type OutOfRangeUpdate struct {
	GUIDS []object.GUID
}

// UpdateType returns the update type in this structure.
func (u *OutOfRangeUpdate) UpdateType() c.UpdateType { return c.UpdateTypeOutOfRangeObjects }

// Update represents a generic game object update.
type Update struct {
	updateType c.UpdateType
	IsSelf     bool
	Object     object.Object
	Victim     object.Object
	WorldTime  uint32
}

// UpdateType returns the update type in this structure.
func (u *Update) UpdateType() c.UpdateType { return u.updateType }

// FieldsUpdate returns the bytes representation of the fields.
func (u *Update) FieldsUpdate() []byte {
	mask := big.NewInt(0)
	fields := bytes.NewBufferString("")

	for field, valueGeneric := range u.Object.UpdateFields() {
		switch value := valueGeneric.(type) {
		case float32:
			binary.Write(fields, binary.LittleEndian, value)
			mask.SetBit(mask, int(field), 1)
		case uint32:
			binary.Write(fields, binary.LittleEndian, value)
			mask.SetBit(mask, int(field), 1)
		default:
			panic(fmt.Sprintf("Unknown field type %T in update fields (%v)", value, field))
		}
	}

	nBlocks := uint8((len(u.Object.UpdateFields()) + 32 - 1) / 32)
	nBytes := uint8((nBlocks * 32) / 8)

	fieldBytes := make([]byte, 0)
	fieldBytes = append(fieldBytes, uint8(nBlocks))
	fieldBytes = append(fieldBytes, common.PadBigIntBytes(common.ReverseBytes(mask.Bytes()), int(nBytes))...)
	fieldBytes = append(fieldBytes, fields.Bytes()...)
	return fieldBytes
}

// Bytes converts the packet into an array of bytes.
func (pkt *ServerUpdateObject) Bytes(stateBase session.State) []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, len(pkt.Updates))
	buffer.WriteByte('\x00') // hasTransportUpdate

	for _, update := range pkt.Updates {
		buffer.WriteByte(uint8(update.UpdateType()))

		if update.UpdateType() == c.UpdateTypeOutOfRangeObjects {
			outOfRangeUpdate := update.(*OutOfRangeUpdate)
			binary.Write(buffer, binary.LittleEndian, uint32(len(outOfRangeUpdate.GUIDS)))
			for _, guid := range outOfRangeUpdate.GUIDS {
				buffer.Write(guid.Pack())
			}
		} else {
			objUpdate := update.(*Update)
			buffer.Write(objUpdate.Object.GUID().Pack())

			if update.UpdateType() != c.UpdateTypeValues {
				updateFlags := object.UpdateFlags(objUpdate.Object)

				if objUpdate.IsSelf {
					updateFlags |= c.UpdateFlagsSelf
				}

				buffer.WriteByte(uint8(object.TypeID(objUpdate.Object)))
				buffer.WriteByte(uint8(updateFlags))

				buffer.Write(objUpdate.Object.MovementUpdate())

				if updateFlags&c.UpdateFlagsHighGUID != 0 {
					binary.Write(buffer, binary.LittleEndian, uint32(objUpdate.Object.GUID().High()))
				}

				if updateFlags&c.UpdateFlagsAll != 0 {
					binary.Write(buffer, binary.LittleEndian, uint32(1))
				}

				if updateFlags&c.UpdateFlagsFullGUID != 0 && objUpdate.Victim != nil {
					buffer.Write(objUpdate.Victim.GUID().Pack())
				}

				if updateFlags&c.UpdateFlagsTransport != 0 {
					binary.Write(buffer, binary.LittleEndian, uint32(objUpdate.WorldTime))
				}
			}

			buffer.Write(objUpdate.FieldsUpdate())
		}
	}

	return buffer.Bytes()
}

// OpCode returns the OpCode of the packet.
func (pkt *ServerUpdateObject) OpCode() session.OpCode {
	return system.OpCodeServerUpdateObject
}
