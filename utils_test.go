package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 2))
	require.Equal(t, 1, Min(2, 1))
}

func TestMax(t *testing.T) {
	require.Equal(t, 2, Max(1, 2))
	require.Equal(t, 2, Max(2, 1))
}

func TestAvg(t *testing.T) {
	require.Equal(t, 1.0, Avg([]float64{0, 1, 2}))
}
