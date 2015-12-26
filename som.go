package gosom

import (
	"fmt"
	"math"
	"math/rand"
)

// SOM holds an instance of a self organizing map.
type SOM struct {
	Width int
	Height int

	Data [][]float64
	Nodes []*Node

	CoolingFunction CoolingFunction
	DistanceFunction DistanceFunction
	NeighborhoodFunction NeighborhoodFunction
}

// NewSOM creates and returns a new self organizing map.
func NewSOM(data [][]float64, width, height int) *SOM {
	som := &SOM{
		Width: width,
		Height: height,
		Data: data,
		Nodes: make([]*Node, width * height),
	}

	d := len(som.Data[0])

	// create nodes
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Height + j
			som.Nodes[k] = NewNode(j, i, d)
		}
	}

	return som
}

func (som *SOM) InitializeWithRandomValues() {
	n := len(som.Data)
	d := len(som.Data[0])

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
}

func (som *SOM) InitializeWithRandomDataPoints() {
	n := len(som.Data)
	d := len(som.Data[0])

	for _, node := range som.Nodes {
		node.Weights = make([]float64, d)
		copy(node.Weights, som.Data[rand.Intn(n-1)])
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
