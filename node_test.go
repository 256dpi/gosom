package gosom

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjust1(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{0.0, 0.0}, 0.5)

	assert.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestAdjust2(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{1.0, 1.0}, 0.5)

	assert.Equal(t, []float64{0.5, 0.5}, n.Weights)
}

func TestAdjust3(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{1.0, 1.0}

	n.Adjust([]float64{1.0, 1.0}, 1.0)

	assert.Equal(t, []float64{1.0, 1.0}, n.Weights)
}

func TestAdjust4(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{0.0, 0.0}, 1.0)

	assert.Equal(t, []float64{0.0, 0.0}, n.Weights)
}

func TestAdjustNaN(t *testing.T) {
	n := NewNode(0, 0, 2)
	n.Weights = []float64{0.0, 0.0}

	n.Adjust([]float64{math.NaN(), 1.0}, 1.0)

	assert.Equal(t, []float64{0.0, 1.0}, n.Weights)
}

func TestX(t *testing.T) {
	n := NewNode(2, 3, 0)

	assert.Equal(t, 2, n.X())
}

func TestY(t *testing.T) {
	n := NewNode(2, 3, 0)

	assert.Equal(t, 3, n.Y())
}
