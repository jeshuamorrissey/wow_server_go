package components

import (
	"bytes"
	"encoding/binary"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
)

type Movement struct {
	interfaces.MovementInfo

	SpeedWalk         float32
	SpeedRun          float32
	SpeedRunBackward  float32
	SpeedSwim         float32
	SpeedSwimBackward float32
	SpeedTurn         float32
}

// MovementUpdate returns a buffer which represents the movement update object for this unit.
func (m *Movement) MovementUpdate() []byte {
	buffer := bytes.NewBufferString("")

	binary.Write(buffer, binary.LittleEndian, uint32(m.movementFlags()))
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time

	binary.Write(buffer, binary.LittleEndian, float32(m.Location.X))
	binary.Write(buffer, binary.LittleEndian, float32(m.Location.Y))
	binary.Write(buffer, binary.LittleEndian, float32(m.Location.Z))
	binary.Write(buffer, binary.LittleEndian, float32(m.Location.O))

	// TODO(jeshua): transport.
	// if GetObjectManager().Exists(m.Transport) {
	// transportObj := GetObjectManager().GetTransport(m.Transport)
	// binary.Write(buffer, binary.LittleEndian, uint64(m.Transport))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location.X))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location.Y))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location.Z))
	// binary.Write(buffer, binary.LittleEndian, float32(transportObj.Location.O))
	// binary.Write(buffer, binary.LittleEndian, uint32(0)) // Time
	// }

	// if m.IsSwimming {
	// 	binary.Write(buffer, binary.LittleEndian, float32(m.MovementInfo.Pitch))
	// }

	// if !GetObjectManager().Exists(m.Transport) {
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // LastFallTime
	// }

	// if m.IsFalling {
	// 	binary.Write(buffer, binary.LittleEndian, float32(m.MovementInfo.Jump.Velocity))
	// 	binary.Write(buffer, binary.LittleEndian, float32(m.MovementInfo.Jump.SinAngle))
	// 	binary.Write(buffer, binary.LittleEndian, float32(m.MovementInfo.Jump.CosAngle))
	// 	binary.Write(buffer, binary.LittleEndian, float32(m.MovementInfo.Jump.XYSpeed))
	// }

	// SplineElevation update goes HERE.

	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedWalk))
	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedRun))
	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedRunBackward))
	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedSwim))
	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedSwimBackward))
	binary.Write(buffer, binary.LittleEndian, float32(m.SpeedTurn))

	// Spline update goes HERE.

	return buffer.Bytes()
}

func (m *Movement) movementFlags() uint32 {
	movementFlags := uint32(0)
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagForward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagBackward)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagStrafeLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagStrafeRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagTurnLeft)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagTurnRight)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagPitchUp)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagPitchDown)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagWalkMode)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagLevitating)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFlying)
	// }
	// if u.IsFalling {
	// 	movementFlags |= uint32(static.MovementFlagFalling)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFallingFar)
	// }
	// if u.IsSwimming {
	// 	movementFlags |= uint32(static.MovementFlagSwimming)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSplineEnabled)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagCanFly)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagFlyingOld)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagOnTransport)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSplineElevation)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagRoot)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagWaterWalking)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagSafeFall)
	// }
	// if u.Is {
	// 	movementFlags |= uint32(static.MovementFlagHover)
	// }
	return movementFlags
}
