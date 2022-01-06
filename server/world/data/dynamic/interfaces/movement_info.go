package interfaces

import "github.com/jeshuamorrissey/wow_server_go/server/world/data/static"

// MovementInfo records movement information for a unit.
type MovementInfo struct {
	MoveFlags static.MovementFlag
	Time      uint32
	Location  Location

	Transport struct {
		GUID     GUID
		Location Location
		Time     uint32
	}

	Pitch    float32  // Swimming pitch.
	FallTime uint32   // Last time the unit fell.
	Jump     struct { // Information about the character's jump.
		Velocity, SinAngle, CosAngle, XYSpeed float32
	}

	// Spline related?
	Unk1 float32
}
