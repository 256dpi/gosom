package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdjust1(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{0.0, 0.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestAdjust2(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{1.0, 1.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestAdjust3(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{1.0, 1.0}, 1.0)

	require.Equal(t, []float64{1.0, 1.0}, n.Weights)
}

func TestAdjust4(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{0.0, 0.0}, 1.0)

	require.Equal(t, []float64{0.0, 0.0}, n.Weights)
}

func TestX(t *testing.T) {
	n := NewNode(2, 3, 0)

	require.Equal(t, 2, n.X())
}

func TestY(t *testing.T) {
	n := NewNode(2, 3, 0)

	require.Equal(t, 3, n.Y())
}
