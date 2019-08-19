package object

import "math"

// Location represents an objects location within the world.
type Location struct {
	X, Y, Z, O float64
}

// Distance calculates the distance between two locations.
func (loc *Location) Distance(other *Location) float64 {
	return math.Sqrt(
		math.Pow(loc.X-other.X, 2) +
			math.Pow(loc.X-other.X, 2) +
			math.Pow(loc.X-other.X, 2))
}
