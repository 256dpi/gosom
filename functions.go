package gosom

import "math"

type DistanceFunction func(a, b []float64) float64
type CoolingFunction func(start, end float64, steps, step int) float64
type NeighborhoodFunction func(distance float64, radius float64) float64

func EuclideanDistance(a, b []float64) float64 {
	d := 0.0
	l := Min(len(a), len(b))

	for i:=0; i<l; i++ {
		d += (a[i] - b[i]) * (a[i] - b[i])
	}

	return math.Sqrt(d)
}

func ManhattanDistance(a, b []float64) float64 {
	d := 0.0
	l := Min(len(a), len(b))

	for i:=0; i<l; i++ {
		d += math.Abs(b[i]-a[i])
	}

	return d
}

func LinearCooling(start, end, steps, step float64) float64 {
	return start - (float64(step) * (start - end) / float64(steps))
}

func ExponentialCooling(start, end, steps, step float64) float64 {
	//TODO: function hacked (+- 0.2) (is there a better implementation?)
	d := -math.Log((end + 0.2) / (start + 0.2)) / float64(steps)
	return ((start + 0.2) * math.Exp(-float64(step) * d)) - 0.2
}

func BubbleNeighborhood(distance float64, radius float64) float64 {
	return 1.0
}

func ConeNeighborhood(distance float64, radius float64) float64 {
	return (radius - math.Abs(distance)) / radius
}

func GaussianNeighborhood(distance float64, radius float64) float64 {
	stdDev := 2.0;
	norm := (2.0 * math.Pow((radius + 1.0), 2.0)) / math.Pow(stdDev, 2.0);
	return math.Exp((-distance * distance) / norm);
}
