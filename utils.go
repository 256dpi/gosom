package gosom

import "github.com/gonum/floats"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Avg(v []float64) float64 {
	return floats.Sum(v) / float64(len(v))
}
