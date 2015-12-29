package gosom

import (
	"math"
	"encoding/csv"
	"io"
	"strconv"
	"encoding/json"
	"math/rand"
)

// A Matrix holds and extends a two dimensional float slice.
type Matrix struct {
	// The float slice
	Data     [][]float64

	// The number of rows in the matrix.
	Rows     int

	// The number of columns in the matrix.
	Columns  int

	// The minimums of the values per column.
	Minimums []float64

	// The maximums of the values per column.
	Maximums []float64

	// The minimum of all values.
	Minimum  float64

	// The maximum of all values.
	Maximum  float64
}

// NewMatrix will create a new Matrix and work out the meta information.
// The function expects the float slice to be consistent.
func NewMatrix(data [][]float64) *Matrix {
	ds := &Matrix{
		Data: data,
		Rows: len(data),
		Columns: len(data[0]),
		Minimums: make([]float64, len(data[0])),
		Maximums: make([]float64, len(data[0])),
	}

	copy(ds.Minimums, ds.Data[0])
	copy(ds.Maximums, ds.Data[0])

	ds.Minimum = ds.Data[0][0]
	ds.Maximum = ds.Data[0][0]

	for j:=0; j<ds.Rows; j++ {
		for i:=0; i<ds.Columns; i++ {
			ds.Minimums[i] = math.Min(ds.Minimums[i], ds.Data[j][i])
			ds.Maximums[i] = math.Max(ds.Maximums[i], ds.Data[j][i])
			ds.Minimum = math.Min(ds.Minimum, ds.Data[j][i])
			ds.Maximum = math.Max(ds.Maximum, ds.Data[j][i])
		}
	}

	return ds
}

func (m *Matrix) RandomRow() []float64 {
	return m.Data[rand.Intn(m.Rows)]
}

func (m *Matrix) SubMatrix(start, length int) *Matrix {
	floats := make([][]float64, m.Rows)

	for i, row := range m.Data {
		floats[i] = make([]float64, length)
		copy(floats[i], row[start:start + length])
	}

	return NewMatrix(floats)
}

func LoadMatrixFromCSV(source io.Reader) (*Matrix, error) {
	reader := csv.NewReader(source)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	floats := make([][]float64, len(data))

	for i, row := range data {
		floats[i] = make([]float64, len(row))

		for j, value := range row {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			} else {
				floats[i][j] = f
			}

		}
	}

	return NewMatrix(floats), nil
}

func LoadMatrixFromJSON(source io.Reader) (*Matrix, error) {
	reader := json.NewDecoder(source)

	floats := make([][]float64, 0)

	err := reader.Decode(&floats)
	if err != nil {
		return nil, err
	}

	return NewMatrix(floats), nil
}
