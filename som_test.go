package gosom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSOM(t *testing.T) {
	som := NewSOM(2, 2)

	require.Equal(t, 2, som.Width)
	require.Equal(t, 2, som.Height)
	require.Equal(t, "linear", som.CoolingFunction)
	require.Equal(t, "euclidean", som.DistanceFunction)
	require.Equal(t, "cone", som.NeighborhoodFunction)
}

func TestLoadSOMFromJSON(t *testing.T) {
	json := `{
	  "Width": 2,
	  "Height": 2,
	  "CoolingFunction": "linear",
	  "DistanceFunction": "euclidean",
	  "NeighborhoodFunction": "cone",
	  "Nodes": [
	    {
	      "Position": [0, 0],
	      "Weights": [0.1, 0.2]
	    }
	  ]
	}`
	reader := strings.NewReader(json)

	som, err := LoadSOMFromJSON(reader)

	require.NoError(t, err)
	require.Equal(t, 2, som.Width)
	require.Equal(t, 2, som.Height)
	require.Equal(t, "linear", som.CoolingFunction)
	require.Equal(t, "euclidean", som.DistanceFunction)
	require.Equal(t, "cone", som.NeighborhoodFunction)
	require.Equal(t, 0, som.Nodes[0].X())
	require.Equal(t, 0, som.Nodes[0].Y())
	require.Equal(t, []float64{0.1, 0.2}, som.Nodes[0].Weights)
}

// Init...

func TestClosest(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[4].Weights[1] = 1.0

	require.Equal(t, som.Closest([]float64{0.0, 1.0}), som.Nodes[4])
}

func TestNeighbors(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[0].Weights[1] = 1.0
	som.Nodes[1].Weights[1] = 0.9
	som.Nodes[2].Weights[1] = 0.8

	require.Equal(t, som.Neighbors([]float64{0.0, 1.0}, 3), []*Node{som.Nodes[0], som.Nodes[1], som.Nodes[2]})
}
