package gosom

import (
	"testing"
	"math"

	"github.com/stretchr/testify/require"
)

func TestEuclideanDistance(t *testing.T) {
	require.Equal(t, math.Sqrt(2), EuclideanDistance(
		[]float64{1.0, 1.0},
		[]float64{0.0, 0.0},
	))

	require.Equal(t, 1.0, EuclideanDistance(
		[]float64{0.0, 1.0},
		[]float64{0.0, 0.0},
	))
}

func TestManhattanDistance(t *testing.T) {
	require.Equal(t, 2.0, ManhattanDistance(
		[]float64{1.0, 1.0},
		[]float64{0.0, 0.0},
	))

	require.Equal(t, 1.0, ManhattanDistance(
		[]float64{0.0, 1.0},
		[]float64{0.0, 0.0},
	))
}

// other functions are tested using the plot
