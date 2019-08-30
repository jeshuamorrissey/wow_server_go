package packet

import (
	"bytes"
	"encoding/binary"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerUpdateObject is the UPDATE_OBJECT packet.
type ServerUpdateObject struct {
	OutOfRangeUpdates system.OutOfRangeUpdate
	ObjectUpdates     []system.ObjectUpdate
}

// ToBytes converts the packet into an array of bytes.
func (pkt *ServerUpdateObject) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	nUpdates := len(pkt.ObjectUpdates)
	if len(pkt.OutOfRangeUpdates.GUIDS) > 0 {
		nUpdates++
	}

	binary.Write(buffer, binary.LittleEndian, uint32(nUpdates))
	binary.Write(buffer, binary.LittleEndian, uint8(0)) // hasTransportUpdate

	if len(pkt.OutOfRangeUpdates.GUIDS) > 0 {
		buffer.WriteByte(uint8(c.UpdateTypeOutOfRangeObjects))
		binary.Write(buffer, binary.LittleEndian, uint32(len(pkt.OutOfRangeUpdates.GUIDS)))
		for _, guid := range pkt.OutOfRangeUpdates.GUIDS {
			buffer.Write(guid.Pack())
		}
	}

	for _, update := range pkt.ObjectUpdates {
		buffer.WriteByte(uint8(update.UpdateType))
		buffer.Write(update.GUID.Pack())

		if update.UpdateType != c.UpdateTypeValues {
			updateFlags := update.UpdateFlags

			if update.IsSelf {
				updateFlags |= c.UpdateFlagsSelf
			}

			buffer.WriteByte(uint8(update.TypeID))
			buffer.WriteByte(uint8(updateFlags))

			if update.MovementUpdate != nil {
				buffer.Write(update.MovementUpdate)
			}

			if updateFlags&c.UpdateFlagsHighGUID != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(update.GUID.High()))
			}

			if updateFlags&c.UpdateFlagsAll != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(1))
			}

			if updateFlags&c.UpdateFlagsFullGUID != 0 && update.VictimGUID != 0 {
				buffer.Write(update.VictimGUID.Pack())
			}

			if updateFlags&c.UpdateFlagsTransport != 0 {
				binary.Write(buffer, binary.LittleEndian, uint32(update.WorldTime))
			}
		}

		buffer.Write(update.UpdateFields.ToBytes(update.NumUpdateFields))
	}

	return buffer.Bytes(), nil
}

// OpCode returns the OpCode of the packet.
func (pkt *ServerUpdateObject) OpCode() system.OpCode {
	return system.OpCodeServerUpdateObject
}
