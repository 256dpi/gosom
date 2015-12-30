package gosom

import "github.com/gonum/floats"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func avg(v []float64) float64 {
	return floats.Sum(v) / float64(len(v))
}
