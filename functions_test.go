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

func TestLinearCooling(t *testing.T) {
	require.Equal(t, 1.0, LinearCooling(1.0, 0.0, 2, 0))
	require.Equal(t, 0.5, LinearCooling(1.0, 0.0, 2, 1))
	require.Equal(t, 0.0, LinearCooling(1.0, 0.0, 2, 2))
}

func TestExponentialCooling(t *testing.T) {
	require.True(t, ExponentialCooling(1.0, 0.0, 2, 0) == 1.0)
	require.True(t, ExponentialCooling(1.0, 0.0, 2, 1) > 0.0)
	require.True(t, ExponentialCooling(1.0, 0.0, 2, 2) == 0.0)
}

func TestBubbleNeighborhood(t *testing.T) {
	require.Equal(t, 1.0, BubbleNeighborhood(0.0, 1.0))
	require.Equal(t, 1.0, BubbleNeighborhood(0.5, 1.0))
	require.Equal(t, 1.0, BubbleNeighborhood(1.0, 1.0))
}

func TestConeNeighborhood(t *testing.T) {
	require.Equal(t, 1.0, ConeNeighborhood(0.0, 1.0))
	require.Equal(t, 0.5, ConeNeighborhood(0.5, 1.0))
	require.Equal(t, 0.0, ConeNeighborhood(1.0, 1.0))
}

func TestGaussianNeighborhood(t *testing.T) {
	require.Equal(t, 1.0, GaussianNeighborhood(0.0, 1.0))
	require.True(t, GaussianNeighborhood(0.5, 1.0) > 0)
	require.True(t, GaussianNeighborhood(1.0, 1.0) > 0)
	require.True(t, GaussianNeighborhood(2.0, 1.0) > 0)
}
