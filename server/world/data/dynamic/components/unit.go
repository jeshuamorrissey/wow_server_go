package components

import "github.com/jeshuamorrissey/wow_server_go/server/world/data/static"

type Unit struct {
	Level      int
	Race       *static.Race
	Class      *static.Class
	Gender     static.Gender
	Team       static.Team
	StandState static.StandState
}
