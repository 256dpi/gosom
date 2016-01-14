package gosom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, 1, min(2, 1))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, max(1, 2))
	assert.Equal(t, 2, max(2, 1))
}

func TestAvg(t *testing.T) {
	assert.Equal(t, 1.0, avg([]float64{0, 1, 2}))
}
