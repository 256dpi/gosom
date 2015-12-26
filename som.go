package gosom

import (
	"fmt"
	"math"
	"math/rand"
)

type Initialization int

const(
	ZeroInitialization Initialization = iota
	RandomInitialization
	RandomDataInitialization
)

type SOM struct {
	Width int
	Height int

	Data [][]float64
	Nodes []*Node

	CoolingFunction CoolingFunction
	DistanceFunction DistanceFunction
	NeighborhoodFunction NeighborhoodFunction
}

func NewSOM(data [][]float64, width, height int) *SOM {
	som := &SOM{
		Width: width,
		Height: height,
		Data: data,
		Nodes: make([]*Node, width * height),
	}

	// create nodes
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Height + j
			som.Nodes[k] = NewNode(j, i)
		}
	}

	return som
}

// Prepare initializes the nodes using the selected Initialization method.
func (som *SOM) Prepare(initialization Initialization) {
	n := len(som.Data)
	d := len(som.Data[0])

	switch initialization {
	case ZeroInitialization:
		for _, n := range som.Nodes {
			n.Weights = make([]float64, d)
		}
	case RandomInitialization:
		min := make([]float64, d)
		max := make([]float64, d)

		for j:=0; j<n; j++ {
			for i:=0; i<d; i++ {
				min[i] = math.Min(min[i], som.Data[j][i])
				max[i] = math.Max(max[i], som.Data[j][i])
			}
		}

		for _, node := range som.Nodes {
			node.Weights = make([]float64, d)
			for i:=0; i<d; i++ {
				node.Weights[i] = rand.Float64() * (max[i] - min[i]) + min[i]
			}
		}
	case RandomDataInitialization:
		for _, node := range som.Nodes {
			node.Weights = make([]float64, d)
			copy(node.Weights, som.Data[rand.Intn(n-1)])
		}
	}
}

func (som *SOM) Train(steps int, initialLearningRate float64) {

}

func (som *SOM) Predict(input []float64) []float64 {
	return []float64{ 0.0 }
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
