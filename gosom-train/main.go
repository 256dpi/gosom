package main

import (
	"github.com/256dpi/gosom"
	"fmt"
)

var data = [][]float64{
	{ 0.0, 4.0 },
	{ 1.0, 3.0 },
	{ 2.0, 2.0 },
	{ 3.0, 1.0 },
	{ 4.0, 0.0 },
}

func main() {
	som := gosom.NewSOM(data, 8, 8)
	som.Initialize()
	som.Train(1000, 0.5)

	fmt.Printf("3.5: %f\n", som.Predict([]float64{0.5}))
	fmt.Printf("2.5: %f\n", som.Predict([]float64{1.5}))
	fmt.Printf("1.5: %f\n", som.Predict([]float64{2.5}))
	fmt.Printf("0.5: %f\n", som.Predict([]float64{3.5}))
}
