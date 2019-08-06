package database

// Gender information.
//go:generate stringer -type=Gender -trimprefix=Gender
type Gender uint8

// Gender values.
const (
	GenderMale   Gender = 0
	GenderFemale Gender = 1
)

// Race information.
//go:generate stringer -type=Race -trimprefix=Race
type Race uint8

// Race values.
const (
	RaceHuman    Race = 0
	RaceOrc      Race = 1
	RaceDwarf    Race = 2
	RaceNightElf Race = 3
	RaceUndead   Race = 4
	RaceTauren   Race = 5
	RaceGnome    Race = 6
	RaceTroll    Race = 7
	RaceGoblin   Race = 8
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
