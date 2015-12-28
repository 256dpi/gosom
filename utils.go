package gosom

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
	t := 0.0
	
	for _, f := range v {
		t += f
	}

	return t / float64(len(v))
}
