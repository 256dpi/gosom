package gosom

import (
	"strings"
	"testing"
	"math"

	"github.com/stretchr/testify/assert"
)

func TestNewSOM(t *testing.T) {
	som := NewSOM(2, 2)

	assert.Equal(t, 2, som.Width)
	assert.Equal(t, 2, som.Height)
	assert.Equal(t, "linear", som.CoolingFunction)
	assert.Equal(t, "euclidean", som.DistanceFunction)
	assert.Equal(t, "cone", som.NeighborhoodFunction)
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

	assert.NoError(t, err)
	assert.Equal(t, 2, som.Width)
	assert.Equal(t, 2, som.Height)
	assert.Equal(t, "linear", som.CoolingFunction)
	assert.Equal(t, "euclidean", som.DistanceFunction)
	assert.Equal(t, "cone", som.NeighborhoodFunction)
	assert.Equal(t, 0, som.Nodes[0].X())
	assert.Equal(t, 0, som.Nodes[0].Y())
	assert.Equal(t, []float64{0.1, 0.2}, som.Nodes[0].Weights)
}

func TestLoadSOMFromJSONError(t *testing.T) {
	json := `-`
	reader := strings.NewReader(json)

	_, err := LoadSOMFromJSON(reader)
	assert.Error(t, err)
}

func TestInitialization(t *testing.T) {
	m := NewMatrix(slice)

	som := NewSOM(5, 5)
	som.InitializeWithZeroes(m.Columns)
	som.InitializeWithRandomValues(m)
	som.InitializeWithDataPoints(m)

	// TODO: check node weights
}

func TestClosest(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[4].Weights[1] = 1.0

	assert.Equal(t, som.Closest([]float64{0.0, 1.0}), som.Nodes[4])
}

func TestClosestWithNaNs(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[4].Weights[1] = 1.0

	assert.Equal(t, som.Closest([]float64{math.NaN(), 1.0}), som.Nodes[4])
}

func TestNeighbors(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[0].Weights[1] = 1.0
	som.Nodes[1].Weights[1] = 0.9
	som.Nodes[2].Weights[1] = 0.8

	assert.Equal(t, som.Neighbors([]float64{0.0, 1.0}, 3), []*Node{som.Nodes[0], som.Nodes[1], som.Nodes[2]})
}

func TestNeighborsWithNaNs(t *testing.T) {
	som := NewSOM(3, 3)
	som.Nodes = NewLattice(3, 3, 2)
	som.Nodes[0].Weights[1] = 1.0
	som.Nodes[1].Weights[1] = 0.9
	som.Nodes[2].Weights[1] = 0.8

	assert.Equal(t, som.Neighbors([]float64{math.NaN(), 1.0}, 3), []*Node{som.Nodes[0], som.Nodes[1], som.Nodes[2]})
}
