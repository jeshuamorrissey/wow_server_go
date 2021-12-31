package packet

import (
	"encoding/binary"
	"io"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ClientMove is sent from the client periodically.
type ClientMove struct {
	MoveOpCode   system.OpCode
	MovementInfo object.MovementInfo
}

// NewClientMovePacket constructs a new movement packet and returns it.
func NewClientMovePacket(opCode system.OpCode) *ClientMove {
	return &ClientMove{
		MoveOpCode: opCode,
	}
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientMove) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.MoveFlags)
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Time)
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Location)
	if pkt.MovementInfo.MoveFlags|c.MovementFlagOnTransport != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Transport)
	}

	if pkt.MovementInfo.MoveFlags|c.MovementFlagSwimming != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Pitch)
	}

	if pkt.MovementInfo.MoveFlags|c.MovementFlagOnTransport == 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.FallTime)
	}

	if pkt.MovementInfo.MoveFlags|c.MovementFlagFalling != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Jump)
	}

	if pkt.MovementInfo.MoveFlags|c.MovementFlagSplineElevation != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Unk1)
	}

	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientMove) Handle(state *system.State) ([]system.ServerPacket, error) {
	state.Character.MovementInfo = pkt.MovementInfo

	location := state.Character.Location()
	location.X = pkt.MovementInfo.Location.X
	location.Y = pkt.MovementInfo.Location.Y
	location.Z = pkt.MovementInfo.Location.Z
	location.O = pkt.MovementInfo.Location.O
	state.Updater.TriggerUpdate(state.Character.GUID())

	return nil, nil
}

// OpCode gets the opcode of the packet.
func (pkt *ClientMove) OpCode() system.OpCode {
	return pkt.MoveOpCode
}
