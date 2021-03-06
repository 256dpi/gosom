package gosom

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var slice = [][]float64{
	{1.0, 0.5, 0.0},
	{0.0, 0.5, 1.0},
}

var sliceNaN = [][]float64{
	{1.0, 0.5, math.NaN()},
	{math.NaN(), 0.5, 1.0},
}

func abstractTestSlice(t *testing.T, m *Matrix) {
	assert.Equal(t, slice, m.Data)
	assert.Equal(t, 2, m.Rows)
	assert.Equal(t, 3, m.Columns)
	assert.Equal(t, []float64{0.0, 0.5, 0.0}, m.Minimums)
	assert.Equal(t, []float64{1.0, 0.5, 1.0}, m.Maximums)
	assert.Equal(t, 0.0, m.Minimum)
	assert.Equal(t, 1.0, m.Maximum)
	assert.False(t, m.NaNs)
}

func abstractTestSliceNaN(t *testing.T, m *Matrix) {
	assert.Equal(t, fmt.Sprint(sliceNaN), fmt.Sprint(m.Data))
	assert.Equal(t, 2, m.Rows)
	assert.Equal(t, 3, m.Columns)
	assert.Equal(t, []float64{1.0, 0.5, 1.0}, m.Minimums)
	assert.Equal(t, []float64{1.0, 0.5, 1.0}, m.Maximums)
	assert.Equal(t, 0.5, m.Minimum)
	assert.Equal(t, 1.0, m.Maximum)
	assert.True(t, m.NaNs)
}

func TestMatrix(t *testing.T) {
	m := NewMatrix(slice)
	abstractTestSlice(t, m)
}

func TestSubMatrix1(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(0, 2)

	d := [][]float64{
		{1.0, 0.5},
		{0.0, 0.5},
	}

	assert.Equal(t, d, sm.Data)
}

func TestSubMatrix2(t *testing.T) {
	m := NewMatrix(slice)
	sm := m.SubMatrix(2, 1)

	d := [][]float64{
		{0.0},
		{1.0},
	}

	assert.Equal(t, d, sm.Data)
}

func TestRandomRow(t *testing.T) {
	m := NewMatrix(slice)

	t1 := assert.ObjectsAreEqual(m.RandomRow(), slice[0])
	t2 := assert.ObjectsAreEqual(m.RandomRow(), slice[1])

	assert.True(t, t1 || t2)
}

func TestLoadMatrixFromCSV(t *testing.T) {
	csv := "1.0,0.5,0.0\n0.0,0.5,1.0"
	reader := strings.NewReader(csv)

	m, err := LoadMatrixFromCSV(reader)

	assert.NoError(t, err)
	abstractTestSlice(t, m)
}

func TestLoadMatrixFromCSVNull(t *testing.T) {
	csv := "1.0,0.5,NULL\nNULL,0.5,1.0"
	reader := strings.NewReader(csv)

	m, err := LoadMatrixFromCSV(reader)
	assert.NoError(t, err)
	abstractTestSliceNaN(t, m)
}

func TestLoadMatrixFromCSVError(t *testing.T) {
	csv := "1.0,0.5,0.0\n0.0,0.5,1.0,2.0"
	reader := strings.NewReader(csv)

	_, err := LoadMatrixFromCSV(reader)
	assert.Error(t, err)
}

func TestLoadMatrixFromJSON(t *testing.T) {
	json := "[[1.0,0.5,0.0],[0.0,0.5,1.0]]"
	reader := strings.NewReader(json)

	m, err := LoadMatrixFromJSON(reader)
	assert.NoError(t, err)
	abstractTestSlice(t, m)
}

func TestLoadMatrixFromJSONull(t *testing.T) {
	json := "[[1.0,0.5,null],[null,0.5,1.0]]"
	reader := strings.NewReader(json)

	m, err := LoadMatrixFromJSON(reader)
	assert.NoError(t, err)
	abstractTestSliceNaN(t, m)
}

func TestLoadMatrixFromJSONError(t *testing.T) {
	json := "-"
	reader := strings.NewReader(json)

	_, err := LoadMatrixFromJSON(reader)
	assert.Error(t, err)
}
