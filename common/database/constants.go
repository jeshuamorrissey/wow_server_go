package database

// Gender information.
//go:generate stringer -type=Gender -trimprefix=Gender
type Gender uint8

// Gender values.
const (
	GenderMale   Gender = 0
	GenderFemale Gender = 1
	GenderNone   Gender = 2
)

// Race information.
//go:generate stringer -type=Race -trimprefix=Race
type Race uint8

// Race values.
const (
	RaceHuman    Race = 1
	RaceOrc      Race = 2
	RaceDwarf    Race = 3
	RaceNightElf Race = 4
	RaceUndead   Race = 5
	RaceTauren   Race = 6
	RaceGnome    Race = 7
	RaceTroll    Race = 8
	RaceGoblin   Race = 9
)

// Class information.
//go:generate stringer -type=Class -trimprefix=Class
type Class uint8

// Class values.
const (
	ClassWarrior Class = 1
	ClassPaladin Class = 2
	ClassHunter  Class = 3
	ClassRouge   Class = 4
	ClassPriest  Class = 5
	ClassShaman  Class = 7
	ClassMage    Class = 8
	ClassWarlock Class = 9
	ClassDruid   Class = 11
)
