package gosom

import "math"

// A DistanceFunction calculates and returns the distance between to points.
type DistanceFunction func(from, to []float64) (distance float64)

// A CoolingFunction returns the cooling alpha [1..0] for an input value [0..1].
type CoolingFunction func(input float64) (output float64)

// A NeighborhoodFunction returns the influence [1..0] of a distance [0..1].
type NeighborhoodFunction func(distance float64) (influence float64)

func EuclideanDistance(from, to []float64) (distance float64) {
	d := 0.0
	l := Min(len(from), len(to))

	for i:=0; i<l; i++ {
		d += (from[i] - to[i]) * (from[i] - to[i])
	}

	return math.Sqrt(d)
}

func ManhattanDistance(from, to []float64) (distance float64) {
	d := 0.0
	l := Min(len(from), len(to))

	for i:=0; i<l; i++ {
		d += math.Abs(to[i]- from[i])
	}

	return d
}

func LinearCooling(input float64) (output float64) {
	return 1.0 - input
}

func SoftCooling(input float64) (output float64) {
	d := -math.Log(0.2 / 1.2)
	return (1.2 * math.Exp(-input * d)) - 0.2
}

func MediumCooling(input float64) (output float64) {
	return 1.005 * math.Pow(0.005 / 1.0, input) - 0.005
}

func HardCooling(input float64) (output float64) {
	d := 1.0 / 101.0
	return (1.0 + d) / (1 + 100 * input) - d
}

func BubbleNeighborhood(distance float64) (influence float64) {
	d := math.Abs(distance)

	if d < 1.0 {
		return 1.0
	} else {
		return 0.0
	}
}

func ConeNeighborhood(distance float64) (influence float64) {
	d := math.Abs(distance)

	if d < 1.0 {
		return (1.0 - d) / 1.0
	} else {
		return 0.0
	}
}

func GaussianNeighborhood(distance float64) (influence float64) {
	stdDev := 5.5
	norm := (2.0 * math.Pow(2.0, 2.0)) / math.Pow(stdDev, 2.0)
	return math.Exp((-distance * distance) / norm)
}

func MexicanHatNeighborhood(distance float64) (influence float64) {
	norm := 3.0 / (2.0)
	square := math.Pow(distance * norm, 2.0)
	return (1.0 - square) * math.Exp(-square)
}
