package gosom

import (
	"testing"
	"math"

	"github.com/stretchr/testify/require"
)

func TestEuclideanDistance(t *testing.T) {
	a := []float64{1.0, 1.0}
	b := []float64{0.0, 0.0}

	require.Equal(t, math.Sqrt(2), EuclideanDistance(a, b))
}

func TestManhattanDistance(t *testing.T) {
	a := []float64{1.0, 1.0}
	b := []float64{0.0, 0.0}

	require.Equal(t, 2.0, ManhattanDistance(a, b))
}

// other functions are tested using the plot
