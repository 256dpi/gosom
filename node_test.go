package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodePosition(t *testing.T) {
	n1 := NewNode(0, 0)

	require.Equal(t, 0, n1.X)
	require.Equal(t, 0, n1.Y)

	n2 := NewNode(1, 1)

	require.Equal(t, 1, n2.X)
	require.Equal(t, 1, n2.Y)
}

func TestNodeAdjust1(t *testing.T) {
	n := NewNode(0, 0)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{0.0, 0.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestNodeAdjust2(t *testing.T) {
	n := NewNode(0, 0)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{1.0, 1.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestNodeAdjust3(t *testing.T) {
	n := NewNode(0, 0)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{1.0, 1.0}, 1.0)

	require.Equal(t, []float64{1.0, 1.0}, n.Weights)
}

func TestNodeAdjust4(t *testing.T) {
	n := NewNode(0, 0)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{0.0, 0.0}, 1.0)

	require.Equal(t, []float64{0.0, 0.0}, n.Weights)
}
