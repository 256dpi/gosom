package gosom

import (
	"testing"

	"github.com/stretchr/testify/require"
	"strings"
)

var slice = [][]float64{
	{ 1.0, 0.5, 0.0 },
	{ 0.0, 0.5, 1.0 },
}

func TestDataSet(t *testing.T) {
	ds := NewDataSet(slice)

	require.Equal(t, slice, ds.Data)
	require.Equal(t, 2, ds.Length)
	require.Equal(t, 3, ds.Dimensions)
	require.Equal(t, []float64{0.0, 0.5, 0.0}, ds.Minimums)
	require.Equal(t, []float64{1.0, 0.5, 1.0}, ds.Maximums)
	require.Equal(t, 0.0, ds.Minimum)
	require.Equal(t, 1.0, ds.Maximum)
}

func TestLoadDataSetFromCSV(t *testing.T) {
	csv := "1.0,0.5,0.0\n0.0,0.5,1.0"
	reader := strings.NewReader(csv)

	ds, err := LoadDataSetFromCSV(reader)

	require.NoError(t, err)
	require.Equal(t, slice, ds.Data)
}

func TestLoadDataSetFromJSON(t *testing.T) {
	json := "[[1.0,0.5,0.0],[0.0,0.5,1.0]]"
	reader := strings.NewReader(json)

	ds, err := LoadDataSetFromJSON(reader)

	require.NoError(t, err)
	require.Equal(t, slice, ds.Data)

}
