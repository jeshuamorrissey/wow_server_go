package packet

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
)

// OutOfRangeUpdate represents a list of GUIDs which are no longer in range.
type OutOfRangeUpdate struct {
	GUIDS []interfaces.GUID
}

// ObjectUpdate represents an update to an individual object.
type ObjectUpdate struct {
	GUID        interfaces.GUID
	UpdateType  static.UpdateType
	UpdateFlags static.UpdateFlags
	IsSelf      bool
	TypeID      static.TypeID

	MovementUpdate  []byte
	NumUpdateFields int
	UpdateFields    interfaces.UpdateFieldsMap

	VictimGUID interfaces.GUID
	WorldTime  uint32
}

// ServerUpdateObject is the UPDATE_OBJECT packet.
type ServerUpdateObject struct {
	OutOfRangeUpdates OutOfRangeUpdate
	ObjectUpdates     []ObjectUpdate
}

// ToBytes converts the packet into an array of bytes.
func (pkt *ServerUpdateObject) ToBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("")

	nUpdates := len(pkt.ObjectUpdates)
	if len(pkt.OutOfRangeUpdates.GUIDS) > 0 {
		nUpdates++
	}

	binary.Write(buffer, binary.LittleEndian, uint32(nUpdates))
	binary.Write(buffer, binary.LittleEndian, uint8(0)) // hasTransportUpdate

	if len(pkt.OutOfRangeUpdates.GUIDS) > 0 {
		buffer.WriteByte(uint8(static.UpdateTypeOutOfRangeObjects))
		binary.Write(buffer, binary.LittleEndian, uint32(len(pkt.OutOfRangeUpdates.GUIDS)))
		for _, guid := range pkt.OutOfRangeUpdates.GUIDS {
			buffer.Write(guid.Pack())
		}
	}

	for _, update := range pkt.ObjectUpdates {
		buffer.WriteByte(uint8(update.UpdateType))
		buffer.Write(update.GUID.Pack())

		if update.UpdateType != static.UpdateTypeValues {
			updateFlags := update.UpdateFlags

			if update.IsSelf {
				updateFlags |= static.UpdateFlagsSelf
			}

			buffer.WriteByte(uint8(update.TypeID))
			buffer.WriteByte(uint8(updateFlags))

			if update.MovementUpdate != nil {
				buffer.Write(update.MovementUpdate)
			}

			if updateFlags&static.UpdateFlagsHighGUID != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(update.GUID.High()))
			}

			if updateFlags&static.UpdateFlagsAll != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(1))
			}

			if updateFlags&static.UpdateFlagsFullGUID != 0 && update.VictimGUID != 0 {
				buffer.Write(update.VictimGUID.Pack())
			}

			if updateFlags&static.UpdateFlagsTransport != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(update.WorldTime))
			}
		}

		buffer.Write(update.UpdateFields.ToBytes(update.NumUpdateFields))
	}

	return buffer.Bytes(), nil
}

// OpCode returns the OpCode of the packet.
func (pkt *ServerUpdateObject) OpCode() static.OpCode {
	return static.OpCodeServerUpdateObject
}
