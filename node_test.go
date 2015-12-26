package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeAdjust1(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{0.0, 0.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestNodeAdjust2(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{1.0, 1.0}, 0.5)

	require.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestNodeAdjust3(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{1.0, 1.0}, 1.0)

	require.Equal(t, []float64{1.0, 1.0}, n.Weights)
}

func TestNodeAdjust4(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{0.0, 0.0}, 1.0)

	require.Equal(t, []float64{0.0, 0.0}, n.Weights)
}
