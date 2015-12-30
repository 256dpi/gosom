package gosom

import (
	"testing"
	"strings"

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
