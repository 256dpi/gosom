package gosom

import (
	"image"
	"image/color"

	"github.com/gonum/floats"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func DrawDimensions(som *SOM, nodeWidth int) []image.Image {
	matrix := som.WeightMatrix()
	images := make([]image.Image, som.Dimensions())

	for i := 0; i < som.Dimensions(); i++ {
		img := image.NewRGBA(image.Rect(0, 0, som.Width*nodeWidth, som.Height*nodeWidth))
		gc := draw2dimg.NewGraphicContext(img)

		for _, node := range som.Nodes {
			r := matrix.Maximums[i] - matrix.Minimums[i]
			g := uint8(((node.Weights[i] - matrix.Minimums[i]) / r) * 255)
			gc.SetFillColor(&color.Gray{Y: g})

			x := node.X() * nodeWidth
			y := node.Y() * nodeWidth
			draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+nodeWidth), float64(y+nodeWidth))
			gc.Fill()
		}

		images[i] = img
	}

	return images
}

func DrawUMatrix(som *SOM, nodeWidth int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, som.Width*nodeWidth, som.Height*nodeWidth))
	gc := draw2dimg.NewGraphicContext(img)

	values := make([]float64, len(som.Nodes))

	for i, node := range som.Nodes {
		distances := make([]float64, 0)

		if node.X() > 1 {
			distances = append(distances, som.D(node.Weights, som.Node(node.X()-1, node.Y()).Weights))
		}
		if node.X()+1 < som.Width {
			distances = append(distances, som.D(node.Weights, som.Node(node.X()+1, node.Y()).Weights))
		}
		if node.Y() > 1 {
			distances = append(distances, som.D(node.Weights, som.Node(node.X(), node.Y()-1).Weights))
		}
		if node.Y()+1 < som.Height {
			distances = append(distances, som.D(node.Weights, som.Node(node.X(), node.Y()+1).Weights))
		}

		values[i] = avg(distances)
	}

	min := floats.Min(values)
	max := floats.Max(values)

	for i, node := range som.Nodes {
		g := 255 - uint8(((values[i]-min)/max-min)*255)
		gc.SetFillColor(&color.Gray{Y: g})

		x := node.X() * nodeWidth
		y := node.Y() * nodeWidth
		draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+nodeWidth), float64(y+nodeWidth))
		gc.Fill()
	}

	return img
}
