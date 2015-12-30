// Package gosom implements the self organizing map algorithm.
package gosom

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"github.com/256dpi/gosom/functions"
)

// SOM holds an instance of a self organizing map.
type SOM struct {
	Width                int
	Height               int
	Nodes                Lattice
	CoolingFunction      string
	DistanceFunction     string
	NeighborhoodFunction string
}

// NewSOM creates and returns a new self organizing map.
func NewSOM(width, height int) *SOM {
	return &SOM{
		Width:                width,
		Height:               height,
		CoolingFunction:      "linear",
		DistanceFunction:     "euclidean",
		NeighborhoodFunction: "cone",
	}
}

func LoadSOMFromJSON(source io.Reader) (*SOM, error) {
	reader := json.NewDecoder(source)
	som := NewSOM(0, 0)

	err := reader.Decode(som)
	if err != nil {
		return nil, err
	}

	return som, nil
}

func (som *SOM) InitializeWithRandomValues(data *Matrix) {
	som.Nodes = NewLattice(som.Width, som.Height, data.Columns)

	for _, node := range som.Nodes {
		for i := 0; i < data.Columns; i++ {
			r := (data.Maximums[i] - data.Minimums[i]) + data.Minimums[i]
			node.Weights[i] = r * rand.Float64()
		}
	}
}

func (som *SOM) InitializeWithDataPoints(data *Matrix) {
	som.Nodes = NewLattice(som.Width, som.Height, data.Columns)

	for _, node := range som.Nodes {
		copy(node.Weights, data.RandomRow())
	}
}

func (som *SOM) Closest(input []float64) *Node {
	nodes := make([]*Node, 0)

	// get initial distance
	t := som.D(input, som.Nodes[0].Weights)

	for _, node := range som.Nodes {
		// calculate distance
		d := som.D(input, node.Weights)

		if d < t {
			// save distance, clear array and add winner
			t = d
			nodes = append([]*Node{}, node)
		} else if d == t {
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
	lat := som.Nodes.Sort(func(n1, n2 *Node) bool {
		d1 := som.D(input, n1.Weights)
		d2 := som.D(input, n2.Weights)

		return d1 < d2
	})

	neighbors := make([]*Node, 0, K)

	for i := 0; i < K; i++ {
		neighbors = append(neighbors, lat[i])
	}

	return neighbors
}

func (som *SOM) Step(data *Matrix, step, steps int, initialLearningRate float64) {
	// calculate position
	progress := float64(step) / float64(steps)

	// calculate learning rate
	learningRate := initialLearningRate * som.CF(progress)

	// calculate neighborhood radius
	initialRadius := float64(Max(som.Width, som.Height)) / 2.0
	radius := initialRadius * som.CF(progress)

	// get random input
	input := data.RandomRow()

	// get closest node to input
	winningNode := som.Closest(input)

	for _, node := range som.Nodes {
		// calculate distance to winner
		distance := som.D(winningNode.Position, node.Position)

		// check inclusion in the radius (doubled to fit gaussian function)
		if distance < radius*2 {
			// calculate the influence
			i := som.NIF(distance / radius)

			// adjust node
			node.Adjust(input, i*learningRate)
		}
	}
}

func (som *SOM) Train(data *Matrix, steps int, initialLearningRate float64) {
	for step := 0; step < steps; step++ {
		som.Step(data, step, steps, initialLearningRate)
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
	for i := 0; i < len(neighbors); i++ {
		for j := 0; j < som.Dimensions(); j++ {
			total[j] += neighbors[i].Weights[j]
		}
	}

	// calculate average
	for i := 0; i < som.Dimensions(); i++ {
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
	radius := som.D(input, neighbors[K-1].Weights)
	for i, n := range neighbors {
		distance := som.D(input, n.Weights)
		neighborWeights[i] = som.NIF(distance / radius)
	}

	// add up all values
	for i := 0; i < len(neighbors); i++ {
		for j := 0; j < som.Dimensions(); j++ {
			total[j] += neighbors[i].Weights[j] * neighborWeights[i]
			sumWeights[j] += neighborWeights[i]
		}
	}

	// calculate average
	for i := 0; i < som.Dimensions(); i++ {
		total[i] = total[i] / sumWeights[i]
	}

	return total
}

// String returns a string matrix of all nodes and weights
func (som *SOM) String() string {
	s := ""

	for i := 0; i < som.Height; i++ {
		for j := 0; j < som.Width; j++ {
			k := i*som.Height + j
			s += fmt.Sprintf("%.2f ", som.Nodes[k].Weights)
		}

		s += "\n"
	}

	return s
}

func (som *SOM) CF(progress float64) float64 {
	return functions.CoolingFactor(som.CoolingFunction, progress)
}

func (som *SOM) D(from, to []float64) float64 {
	return functions.Distance(som.DistanceFunction, from, to)
}

func (som *SOM) NIF(distance float64) float64 {
	return functions.NeighborhoodInfluenceFactor(som.NeighborhoodFunction, distance)
}

func (som *SOM) Dimensions() int {
	return len(som.Nodes[0].Weights)
}

func (som *SOM) WeightMatrix() *Matrix {
	data := make([][]float64, len(som.Nodes))

	for i, node := range som.Nodes {
		data[i] = make([]float64, som.Dimensions())
		copy(data[i], node.Weights)
	}

	return NewMatrix(data)
}

func (som *SOM) Node(x, y int) *Node {
	return som.Nodes[y*som.Width+x]
}

func (som *SOM) SaveAsJSON(destination io.Writer) error {
	writer := json.NewEncoder(destination)
	return writer.Encode(som)
}
