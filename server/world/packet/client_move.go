package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// ClientMove is sent from the client periodically.
type ClientMove struct {
	MoveOpCode   static.OpCode
	MovementInfo interfaces.MovementInfo
}

// NewClientMovePacket constructs a new movement packet and returns it.
func NewClientMovePacket(opCode static.OpCode) *ClientMove {
	return &ClientMove{
		MoveOpCode: opCode,
	}
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientMove) FromBytes(state *system.State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.MoveFlags)
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Time)
	binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Location)
	if pkt.MovementInfo.MoveFlags|static.MovementFlagOnTransport != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Transport)
	}

	if pkt.MovementInfo.MoveFlags|static.MovementFlagSwimming != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Pitch)
	}

	if pkt.MovementInfo.MoveFlags|static.MovementFlagOnTransport == 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.FallTime)
	}

	if pkt.MovementInfo.MoveFlags|static.MovementFlagFalling != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Jump)
	}

	if pkt.MovementInfo.MoveFlags|static.MovementFlagSplineElevation != 0 {
		binary.Read(buffer, binary.LittleEndian, &pkt.MovementInfo.Unk1)
	}

	return nil
}

// OpCode gets the opcode of the packet.
func (pkt *ClientMove) OpCode() static.OpCode {
	return pkt.MoveOpCode
}
