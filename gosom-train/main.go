package main

import (
	"fmt"

	"github.com/256dpi/gosom"
	"github.com/llgcode/draw2d/draw2dimg"
)

var data = [][]float64{
	{ 0.0, 4.0 },
	{ 1.0, 3.0 },
	{ 2.0, 2.0 },
	{ 3.0, 1.0 },
	{ 4.0, 0.0 },
}

func main() {
	som := gosom.NewSOM(data, 64, 64)
	som.InitializeWithDataPoints()
	som.Train(10000, 0.5)

	fmt.Printf("3.5: %f\n", som.Classify([]float64{0.5})[1])
	fmt.Printf("2.5: %f\n", som.Classify([]float64{1.5})[1])
	fmt.Printf("1.5: %f\n", som.Classify([]float64{2.5})[1])
	fmt.Printf("0.5: %f\n\n", som.Classify([]float64{3.5})[1])

	fmt.Printf("3.5: %f\n", som.Interpolate([]float64{0.5}, 8)[1])
	fmt.Printf("2.5: %f\n", som.Interpolate([]float64{1.5}, 8)[1])
	fmt.Printf("1.5: %f\n", som.Interpolate([]float64{2.5}, 8)[1])
	fmt.Printf("0.5: %f\n\n", som.Interpolate([]float64{3.5}, 8)[1])

	fmt.Printf("0.75: %f\n", som.Classify([]float64{3.25})[1])
	fmt.Printf("0.75: %f\n", som.Interpolate([]float64{3.25}, 8)[1])
	fmt.Printf("0.75: %f\n", som.WeightedInterpolate([]float64{3.25}, 8)[1])

	images := som.DimensionImages(5)

	for i, img := range images {
		draw2dimg.SaveToPngFile(fmt.Sprintf("dimension%d.png", i), img)
	}
}
