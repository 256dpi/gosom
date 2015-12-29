package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
	"strings"
)

var slice = [][]float64{
	{1.0, 0.5, 0.0},
	{0.0, 0.5, 1.0},
}

func TestMatrix(t *testing.T) {
	m := NewMatrix(slice)

	require.Equal(t, slice, m.Data)
	require.Equal(t, 2, m.Rows)
	require.Equal(t, 3, m.Columns)
	require.Equal(t, []float64{0.0, 0.5, 0.0}, m.Minimums)
	require.Equal(t, []float64{1.0, 0.5, 1.0}, m.Maximums)
	require.Equal(t, 0.0, m.Minimum)
	require.Equal(t, 1.0, m.Maximum)
}

func TestSubMatrix1(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(0, 2)

	d := [][]float64{
		{1.0, 0.5},
		{0.0, 0.5},
	}

	require.Equal(t, d, sm.Data)
}

func TestSubMatrix2(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(2, 1)

	d := [][]float64{
		{0.0},
		{1.0},
	}

	require.Equal(t, d, sm.Data)
}

func TestLoadMatrixFromCSV(t *testing.T) {
	csv := "1.0,0.5,0.0\n0.0,0.5,1.0"
	reader := strings.NewReader(csv)

	m, err := LoadMatrixFromCSV(reader)

	require.NoError(t, err)
	require.Equal(t, slice, m.Data)
}

func TestLoadMatrixFromJSON(t *testing.T) {
	json := "[[1.0,0.5,0.0],[0.0,0.5,1.0]]"
	reader := strings.NewReader(json)

	m, err := LoadMatrixFromJSON(reader)

	require.NoError(t, err)
	require.Equal(t, slice, m.Data)

}
