package gosom

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestDrawDimensions(t *testing.T) {
	som := NewSOM(5, 5)
	som.InitializeWithZeroes(2)

	require.NotNil(t, DrawDimensions(som, 5))
}

func TestDrawUMatrix(t *testing.T) {
	som := NewSOM(5, 5)
	som.InitializeWithZeroes(2)

	require.NotNil(t, DrawUMatrix(som, 5))
}
