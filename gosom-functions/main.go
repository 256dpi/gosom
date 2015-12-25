package main

import (
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/256dpi/gosom"
)

func main() {
	plotCoolingFunctions()
	plotNeighborhoodFunctions()
}

func plotCoolingFunctions() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "CoolingFunctions"
	p.X.Label.Text = "Step"
	p.Y.Label.Text = "Value"

	linear := plotter.NewFunction(func(x float64) float64 { return gosom.LinearCooling(1.0, 0.0, 100, x) })
	linear.Color = color.RGBA{B: 255, A: 255}

	exponential := plotter.NewFunction(func(x float64) float64 { return gosom.ExponentialCooling(1.0, 0.0, 100, x) })
	exponential.Color = color.RGBA{G: 255, A: 255}

	p.Add(linear, exponential)
	p.Legend.Add("LinearCooling", linear)
	p.Legend.Add("ExponentialCooling", exponential)

	p.X.Min = 0
	p.X.Max = 100
	p.Y.Min = 0
	p.Y.Max = 1

	if err := p.Save(500, 500, "cooling.png"); err != nil {
		panic(err)
	}
}

func plotNeighborhoodFunctions() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "NeighborhoodFunctions"
	p.X.Label.Text = "Distance"
	p.Y.Label.Text = "Influence"

	bubble := plotter.NewFunction(func(x float64) float64 { return gosom.BubbleNeighborhood(x, 100) })
	bubble.Color = color.RGBA{B: 255, A: 255}

	cone := plotter.NewFunction(func(x float64) float64 { return gosom.ConeNeighborhood(x, 100) })
	cone.Color = color.RGBA{G: 255, A: 255}

	gauss := plotter.NewFunction(func(x float64) float64 { return gosom.GaussianNeighborhood(x, 100) })
	gauss.Color = color.RGBA{R: 255, A: 255}

	p.Add(bubble, cone, gauss)
	p.Legend.Add("BubbleNeighborhood", bubble)
	p.Legend.Add("ConeNeighborhood", cone)
	p.Legend.Add("GaussianNeighborhood", gauss)

	p.X.Min = -200
	p.X.Max = 200
	p.Y.Min = 0
	p.Y.Max = 1

	if err := p.Save(500, 500, "neighborhood.png"); err != nil {
		panic(err)
	}
}
