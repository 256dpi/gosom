package gosom

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawDimensions(som *SOM, matrix *Matrix, nodeWidth int) []image.Image {
	images := make([]image.Image, som.Dimensions())

	for i:=0; i<som.Dimensions(); i++ {
		img := image.NewRGBA(image.Rect(0, 0, som.Width*nodeWidth, som.Height*nodeWidth))
		gc := draw2dimg.NewGraphicContext(img)

		for _, node := range som.Nodes {
			r := matrix.Maximums[i] - matrix.Minimums[i]
			g := uint8(((node.Weights[i] - matrix.Minimums[i]) / r) * 255)
			gc.SetFillColor(&color.Gray{ Y: g })

			x := node.X() * nodeWidth
			y := node.Y() * nodeWidth
			draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+nodeWidth), float64(y+nodeWidth))
			gc.Fill()
		}

		images[i] = img
	}

	return images
}
