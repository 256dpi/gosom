package gosom

import (
	"fmt"
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

	CoolingFunction CoolingFunction `json:"-"`
	DistanceFunction DistanceFunction `json:"-"`
	NeighborhoodFunction NeighborhoodFunction `json:"-"`
}

// NewSOM creates and returns a new self organizing map.
func NewSOM(width, height int) *SOM {
	return &SOM{
		Width: width,
		Height: height,
		Nodes: make([]*Node, width * height),
		CoolingFunction: LinearCooling,
		DistanceFunction: EuclideanDistance,
		NeighborhoodFunction: ConeNeighborhood,
	}
}

func (som *SOM) createNodes(dimensions int) {
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Width + j
			som.Nodes[k] = NewNode(j, i, dimensions)
		}
	}
}

func (som *SOM) InitializeWithRandomValues(dataSet *DataSet) {
	som.createNodes(dataSet.Dimensions)

	for _, node := range som.Nodes {
		for i:=0; i<dataSet.Dimensions; i++ {
			r := (dataSet.Maximums[i] - dataSet.Minimums[i]) + dataSet.Minimums[i]
			node.Weights[i] = r * rand.Float64()
		}
	}
}

func (som *SOM) InitializeWithDataPoints(dataSet *DataSet) {
	som.createNodes(dataSet.Dimensions)

	for _, node := range som.Nodes {
		copy(node.Weights, dataSet.RandomDataPoint())
	}
}

func (som *SOM) Closest(input []float64) *Node {
	nodes := make([]*Node, 0)

	// get initial distance
	t := som.DistanceFunction(input, som.Nodes[0].Weights)

	for _, node := range som.Nodes {
		// calculate distance
		d := som.DistanceFunction(input, node.Weights)

		if(d < t) {
			// save distance, clear array and add winner
			t = d
			nodes = append([]*Node{}, node)
		} else if(d == t) {
			// add winner
			nodes = append(nodes, node)
		}
	}

	if len(nodes) > 1 {
		// return random winner
		return nodes[rand.Intn(len(nodes))]
	}

	return nodes[0]
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

func (som *SOM) Step(dataSet *DataSet, step, steps int, initialLearningRate float64) {
	// calculate position
	pos := float64(step) / float64(steps)

	// calculate learning rate
	learningRate := initialLearningRate * som.CoolingFunction(pos)

	// calculate neighborhood radius
	initialRadius := float64(Max(som.Width, som.Height)) / 2.0
	radius := initialRadius * som.CoolingFunction(pos)

	// get closest node to input
	winningNode := som.Closest(dataSet.RandomDataPoint())

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

func (som *SOM) Train(dataSet *DataSet, steps int, initialLearningRate float64) {
	for step:=0; step<steps; step++ {
		som.Step(dataSet, step, steps, initialLearningRate)
	}
}

func (som *SOM) Classify(input []float64) []float64 {
	o := make([]float64, som.Dimensions())
	copy(o, som.Closest(input).Weights)
	return o
}

func (som *SOM) Interpolate(input []float64, K int) []float64 {
	neighbors := som.Neighbors(input, K)
	total := make([]float64, som.Dimensions())

	// add up all values
	for i:=0; i<len(neighbors); i++ {
		for j:=0; j<som.Dimensions(); j++ {
			total[j] += neighbors[i].Weights[j]
		}
	}

	// calculate average
	for i:=0; i<som.Dimensions(); i++ {
		total[i] = total[i] / float64(K)
	}

	return total
}

func (som *SOM) WeightedInterpolate(input []float64, K int) []float64 {
	neighbors := som.Neighbors(input, K)
	neighborWeights := make([]float64, K)
	total := make([]float64, som.Dimensions())
	sumWeights := make([]float64, som.Dimensions())

	// calculate weights for neighbors
	radius := som.DistanceFunction(input, neighbors[K-1].Weights)
	for i, n := range neighbors {
		distance := som.DistanceFunction(input, n.Weights)
		neighborWeights[i] = som.NeighborhoodFunction(distance / radius)
	}

	// add up all values
	for i:=0; i<len(neighbors); i++ {
		for j:=0; j<som.Dimensions(); j++ {
			total[j] += neighbors[i].Weights[j] * neighborWeights[i]
			sumWeights[j] += neighborWeights[i]
		}
	}

	// calculate average
	for i:=0; i<som.Dimensions(); i++ {
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

func (som *SOM) DimensionImages(dataSet *DataSet, nodeWidth int) []image.Image {
	images := make([]image.Image, som.Dimensions())

	for i:=0; i<som.Dimensions(); i++ {
		img := image.NewRGBA(image.Rect(0, 0, som.Width*nodeWidth, som.Height*nodeWidth))
		gc := draw2dimg.NewGraphicContext(img)

		for _, node := range som.Nodes {
			r := dataSet.Maximums[i] - dataSet.Minimums[i]
			g := uint8(((node.Weights[i] - dataSet.Minimums[i]) / r) * 255)
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

func (som *SOM) Dimensions() int {
	return len(som.Nodes[0].Weights)
}
