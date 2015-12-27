package gosom

import (
	"fmt"
	"math"
	"math/rand"
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// SOM holds an instance of a self organizing map.
type SOM struct {
	Width int
	Height int
	Nodes []*Node

	Data [][]float64
	Rows int
	Columns int
	Min []float64
	Max []float64

	CoolingFunction CoolingFunction
	DistanceFunction DistanceFunction
	NeighborhoodFunction NeighborhoodFunction
}

// NewSOM creates and returns a new self organizing map.
func NewSOM(data [][]float64, width, height int) *SOM {
	som := &SOM{
		Width: width,
		Height: height,
		Nodes: make([]*Node, width * height),
		Data: data,
		Rows: len(data),
		Columns: len(data[0]),
		CoolingFunction: LinearCooling,
		DistanceFunction: ManhattanDistance,
		NeighborhoodFunction: ConeNeighborhood,
	}

	// create nodes
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Height + j
			som.Nodes[k] = NewNode(j, i, som.Columns)
		}
	}

	som.Min = make([]float64, som.Columns)
	som.Max = make([]float64, som.Columns)

	// find min and max
	for j:=0; j<som.Rows; j++ {
		for i:=0; i<som.Columns; i++ {
			som.Min[i] = math.Min(som.Min[i], som.Data[j][i])
			som.Max[i] = math.Max(som.Max[i], som.Data[j][i])
		}
	}

	return som
}

func (som *SOM) InitializeWithRandomValues() {
	for _, node := range som.Nodes {
		node.Weights = make([]float64, som.Columns)
		for i:=0; i<som.Columns; i++ {
			node.Weights[i] = rand.Float64() * (som.Max[i] - som.Min[i]) + som.Min[i]
		}
	}
}

func (som *SOM) InitializeWithRandomDataPoints() {
	for _, node := range som.Nodes {
		node.Weights = make([]float64, som.Columns)
		copy(node.Weights, som.Data[rand.Intn(som.Rows)])
	}
}

func (som *SOM) Closest(input []float64) *Node {
	n := make([]*Node, 0, 1)

	// get initial distance
	t := som.DistanceFunction(input, som.Nodes[0].Weights)

	for _, node := range som.Nodes {
		// calculate distance
		d := som.DistanceFunction(input, node.Weights)

		if(d < t) {
			// save distance, clear array and add winner
			t = d
			n = append([]*Node{}, node)
		} else if(d <= t) {
			// add winner
			n = append(n, node)
		}
	}

	if len(n) > 1 {
		// return random winner
		return n[rand.Intn(len(n))]
	}

	return n[0]
}

func (som *SOM) Neighbors(input []float64, K int) []*Node {
	nodes := make([]*Node, len(som.Nodes))
	copy(nodes, som.Nodes)

	SortNodes(nodes, func(n1, n2 *Node)(bool){
		d1 := som.DistanceFunction(input, n1.Weights)
		d2 := som.DistanceFunction(input, n2.Weights)

		return d1 < d2
	})

	neighbors := make([]*Node, 0, K)

	for i:=0; i<K; i++ {
		neighbors = append(neighbors, nodes[i])
	}

	return neighbors
}

func (som *SOM) Train(steps int, initialLearningRate float64) {
	initialRadius := float64(Max(som.Width, som.Height)) / 2.0

	for step:=0; step<steps; step++ {
		s := float64(step) / float64(steps)

		// calculate learning rate
		learningRate := initialLearningRate * som.CoolingFunction(s)

		// calculate neighborhood radius
		radius := initialRadius * som.CoolingFunction(s)

		// pick random input point
		dataPoint := som.Data[rand.Intn(som.Rows)]

		// get closest node to input
		winningNode := som.Closest(dataPoint)

		for _, node := range som.Nodes {
			// calculate distance to winner
			distance := som.DistanceFunction(winningNode.Position, node.Position)

			// check inclusion in the radius (doubled to fit gaussian function)
			if(distance < radius * 2) {
				// calculate the influence
				i := som.NeighborhoodFunction(distance / radius)

				// adjust node
				node.Adjust(winningNode.Weights, i * learningRate)
			}
		}
	}
}

func (som *SOM) Classify(input []float64) []float64 {
	o := make([]float64, som.Columns)
	copy(o, som.Closest(input).Weights)
	return o
}

func (som *SOM) Interpolate(input []float64, K int) []float64 {
	neighbors := som.Neighbors(input, K)
	total := make([]float64, som.Columns)

	// add up all values
	for i:=0; i<len(neighbors); i++ {
		for j:=0; j<som.Columns; j++ {
			total[j] += neighbors[i].Weights[j]
		}
	}

	// calculate average
	for i:=0; i<som.Columns; i++ {
		total[i] = total[i] / float64(K)
	}

	return total
}

func (som *SOM) WeightedInterpolate(input []float64, K int) []float64 {
	neighbors := som.Neighbors(input, K)
	neighborWeights := make([]float64, K)
	total := make([]float64, som.Columns)
	sumWeights := make([]float64, som.Columns)

	// calculate weights for neighbors
	radius := som.DistanceFunction(input, neighbors[K-1].Weights)
	for i, n := range neighbors {
		distance := som.DistanceFunction(input, n.Weights)
		neighborWeights[i] = som.NeighborhoodFunction(distance / radius)
	}

	// add up all values
	for i:=0; i<len(neighbors); i++ {
		for j:=0; j<som.Columns; j++ {
			total[j] += neighbors[i].Weights[j] * neighborWeights[i]
			sumWeights[j] += neighborWeights[i]
		}
	}

	// calculate average
	for i:=0; i<som.Columns; i++ {
		total[i] = total[i] / sumWeights[i]
	}

	return total
}

// String returns a string matrix of all nodes and weights
func (som *SOM) String() string {
	s := ""

	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Height + j
			s += fmt.Sprintf("%.2f ", som.Nodes[k].Weights)
		}

		s += "\n"
	}

	return s
}

func (som *SOM) DimensionImages(nodeWidth int) []image.Image {
	images := make([]image.Image, som.Columns)

	for i:=0; i<som.Columns; i++ {
		img := image.NewRGBA(image.Rect(0, 0, som.Width*nodeWidth, som.Height*nodeWidth))
		gc := draw2dimg.NewGraphicContext(img)

		for _, node := range som.Nodes {
			g := uint8(((node.Weights[i] - som.Min[i]) / (som.Max[i] - som.Min[i])) * 255)
			gc.SetFillColor(&color.Gray{ Y: g })

			x := node.Position[0] * float64(nodeWidth)
			y := node.Position[1] * float64(nodeWidth)
			draw2dkit.Rectangle(gc, x, y, x+float64(nodeWidth), y+float64(nodeWidth))
			gc.Fill()
		}

		images[i] = img
	}

	return images
}
