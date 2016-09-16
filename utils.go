package gosom

import (
	"math"

	"github.com/gonum/floats"
)

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

func clearNANs(in []float64) []float64 {
	var out []float64

	for _, f := range in {
		if !math.IsNaN(f) {
			out = append(out, f)
		}
	}

	return out
}
