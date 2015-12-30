package functions

import "math"

// A DistanceFunction calculates the distance between to points.
type DistanceFunction func(from, to []float64) (distance float64)

// A CoolingFunction calculates the cooling alpha [1..0] for an input value [0..1].
type CoolingFunction func(progress float64) (factor float64)

// A NeighborhoodFunction calculates the influence [1..0] of a distance [0..1..2].
type NeighborhoodFunction func(distance float64) (influence float64)

// EuclideanDistance returns the euclidean distance between two points.
func EuclideanDistance(from, to []float64) (distance float64) {
	d := 0.0
	l := min(len(from), len(to))

	for i := 0; i < l; i++ {
		d += (from[i] - to[i]) * (from[i] - to[i])
	}

	return math.Sqrt(d)
}

// ManhattanDistance returns the manhattan distance between two points.
func ManhattanDistance(from, to []float64) (distance float64) {
	d := 0.0
	l := min(len(from), len(to))

	for i := 0; i < l; i++ {
		d += math.Abs(to[i] - from[i])
	}

	return d
}

// LinearCooling returns the linear cooling factor for progress.
func LinearCooling(progress float64) (factor float64) {
	return 1.0 - progress
}

// SoftCooling returns the soft exponential cooling factor for progress.
func SoftCooling(progress float64) (factor float64) {
	d := -math.Log(0.2 / 1.2)
	return (1.2 * math.Exp(-progress*d)) - 0.2
}

// MediumCooling returns the medium exponential cooling factor for progress.
func MediumCooling(progress float64) (factor float64) {
	return 1.005*math.Pow(0.005/1.0, progress) - 0.005
}

// HardCooling returns the hard exponential cooling factor for progress.
func HardCooling(progress float64) (factor float64) {
	d := 1.0 / 101.0
	return (1.0+d)/(1+100*progress) - d
}

// BubbleNeighborhood returns the influence for the specified distance.
func BubbleNeighborhood(distance float64) (influence float64) {
	d := math.Abs(distance)

	if d < 1.0 {
		return 1.0
	}

	return 0.0
}

// ConeNeighborhood returns the influence for the specified distance.
func ConeNeighborhood(distance float64) (influence float64) {
	d := math.Abs(distance)

	if d < 1.0 {
		return (1.0 - d) / 1.0
	}

	return 0.0
}

// GaussianNeighborhood returns the influence for the specified distance.
func GaussianNeighborhood(distance float64) (influence float64) {
	stdDev := 4.0
	norm := (2.0 * math.Pow(2.0, 2.0)) / math.Pow(stdDev, 2.0)
	return math.Exp((-distance * distance) / norm)
}

// CoolingFactor returns the cooling factors based on the selected coolingFunction.
func CoolingFactor(coolingFunction string, progress float64) (factor float64) {
	switch coolingFunction {
	case "linear":
		return LinearCooling(progress)
	case "soft":
		return SoftCooling(progress)
	case "medium":
		return MediumCooling(progress)
	case "hard":
		return HardCooling(progress)
	}

	return 0.0
}

// Distance returns the distance between two points based on the selected distanceFunction.
func Distance(distanceFunction string, from, to []float64) (distance float64) {
	switch distanceFunction {
	case "euclidean":
		return EuclideanDistance(from, to)
	case "manhattan":
		return ManhattanDistance(from, to)
	}

	return 0.0
}

// NeighborhoodInfluence returns the influence of the distnance based on the selected neighborhoodFunction.
func NeighborhoodInfluence(neighborhoodFunction string, distance float64) (influence float64) {
	switch neighborhoodFunction {
	case "bubble":
		return BubbleNeighborhood(distance)
	case "cone":
		return ConeNeighborhood(distance)
	case "gaussian":
		return GaussianNeighborhood(distance)
	}

	return 0.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
