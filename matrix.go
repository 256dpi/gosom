package gosom

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"math"
	"math/rand"
	"strconv"

	"github.com/gonum/floats"
)

// A Matrix holds and extends a two dimensional float slice.
type Matrix struct {
	Data     [][]float64
	Rows     int
	Columns  int
	Minimums []float64
	Maximums []float64
	Minimum  float64
	Maximum  float64
	NaNs     bool
}

// NewMatrix will create a new Matrix and work out the meta information.
// The function expects the float slice to be consistent in columns.
func NewMatrix(data [][]float64) *Matrix {
	m := &Matrix{
		Data:    data,
		Rows:    len(data),
		Columns: len(data[0]),
	}

	m.Minimums = make([]float64, m.Columns)
	m.Maximums = make([]float64, m.Columns)

	for i := 0; i < m.Columns; i++ {
		rawColumn := m.Column(i)
		clearedColumn := clearNANs(rawColumn)

		m.Minimums[i] = floats.Min(clearedColumn)
		m.Maximums[i] = floats.Max(clearedColumn)

		if floats.HasNaN(rawColumn) {
			m.NaNs = true
		}
	}

	m.Minimum = floats.Min(m.Minimums)
	m.Maximum = floats.Max(m.Maximums)

	return m
}

// Column returns all values in a column.
func (m *Matrix) Column(col int) []float64 {
	out := make([]float64, m.Rows)

	for i, row := range m.Data {
		out[i] = row[col]
	}

	return out
}

// RandomRow returns a random row from the matrix.
func (m *Matrix) RandomRow() []float64 {
	return m.Data[rand.Intn(m.Rows)]
}

// SubMatrix returns a matrix that holds a subset of the current matrix.
func (m *Matrix) SubMatrix(start, length int) *Matrix {
	values := make([][]float64, m.Rows)

	for i, row := range m.Data {
		values[i] = make([]float64, length)
		copy(values[i], row[start:start+length])
	}

	return NewMatrix(values)
}

// LoadMatrixFromCSV reads CSV data and returns a new matrix.
func LoadMatrixFromCSV(source io.Reader) (*Matrix, error) {
	reader := csv.NewReader(source)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([][]float64, len(data))

	for i, row := range data {
		values[i] = make([]float64, len(row))

		for j, value := range row {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				values[i][j] = math.NaN()
			} else {
				values[i][j] = f
			}
		}
	}

	return NewMatrix(values), nil
}

// LoadMatrixFromJSON read JSON data and returns a new matrix.
func LoadMatrixFromJSON(source io.Reader) (*Matrix, error) {
	reader := json.NewDecoder(source)

	var data [][]interface{}

	err := reader.Decode(&data)
	if err != nil {
		return nil, err
	}

	values := make([][]float64, len(data))

	for i := 0; i < len(data); i++ {
		values[i] = make([]float64, len(data[i]))

		for j := 0; j < len(data[i]); j++ {
			if f, ok := data[i][j].(float64); ok {
				values[i][j] = f
			} else {
				values[i][j] = math.NaN()
			}
		}
	}

	return NewMatrix(values), nil
}
