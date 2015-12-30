package main

import (
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/256dpi/gosom/functions"
)

func plotCoolingFunctions(file string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "CoolingFunctions"
	p.X.Label.Text = "Input"
	p.Y.Label.Text = "Output"

	linear := plotter.NewFunction(func(x float64) float64 {
		return functions.LinearCooling(x)
	})
	linear.Color = color.RGBA{B: 255, A: 255}
	linear.Samples = 100

	soft := plotter.NewFunction(func(x float64) float64 {
		return functions.SoftCooling(x)
	})
	soft.Color = color.RGBA{G: 255, A: 255}
	soft.Samples = 100

	medium := plotter.NewFunction(func(x float64) float64 {
		return functions.MediumCooling(x)
	})
	medium.Color = color.RGBA{R: 255, B: 255, A: 255}
	medium.Samples = 100

	hard := plotter.NewFunction(func(x float64) float64 {
		return functions.HardCooling(x)
	})
	hard.Color = color.RGBA{G: 255, B: 255, A: 255}
	hard.Samples = 100

	p.Add(linear, soft, medium, hard)
	p.Legend.Add("LinearCooling", linear)
	p.Legend.Add("SoftCooling", soft)
	p.Legend.Add("MediumCooling", medium)
	p.Legend.Add("HardCooling", hard)

	p.X.Min = 0.0
	p.X.Max = 1.0
	p.Y.Min = 0.0
	p.Y.Max = 1.0

	if err := p.Save(500, 500, file); err != nil {
		panic(err)
	}
}

func plotNeighborhoodFunctions(file string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "NeighborhoodFunctions"
	p.X.Label.Text = "Distance"
	p.Y.Label.Text = "Influence"

	bubble := plotter.NewFunction(func(x float64) float64 {
		return functions.BubbleNeighborhood(x)
	})
	bubble.Color = color.RGBA{B: 255, A: 255}
	bubble.Samples = 200

	cone := plotter.NewFunction(func(x float64) float64 {
		return functions.ConeNeighborhood(x)
	})
	cone.Color = color.RGBA{G: 255, A: 255}
	cone.Samples = 200

	gaussian := plotter.NewFunction(func(x float64) float64 {
		return functions.GaussianNeighborhood(x)
	})
	gaussian.Color = color.RGBA{R: 255, B: 255, A: 255}
	gaussian.Samples = 200

	mexicanHat := plotter.NewFunction(func(x float64) float64 {
		return functions.MexicanHatNeighborhood(x)
	})
	mexicanHat.Color = color.RGBA{G: 255, B: 255, A: 255}
	mexicanHat.Samples = 200

	p.Add(bubble, cone, gaussian, mexicanHat)
	p.Legend.Add("BubbleNeighborhood", bubble)
	p.Legend.Add("ConeNeighborhood", cone)
	p.Legend.Add("GaussianNeighborhood", gaussian)
	p.Legend.Add("MexicanHatNeighborhood", mexicanHat)

	p.X.Min = -2.0
	p.X.Max = 2.0
	p.Y.Min = -0.2
	p.Y.Max = 1.0

	if err := p.Save(500, 500, file); err != nil {
		panic(err)
	}
}
