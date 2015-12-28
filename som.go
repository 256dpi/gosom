package gosom

import (
	"io"
	"fmt"
	"math/rand"
	"encoding/json"
)

// SOM holds an instance of a self organizing map.
type SOM struct {
	Width int
	Height int
	Nodes []*Node
	CoolingFunction string
	DistanceFunction string
	NeighborhoodFunction string
}

// NewSOM creates and returns a new self organizing map.
func NewSOM(width, height int) *SOM {
	return &SOM{
		Width: width,
		Height: height,
		Nodes: make([]*Node, width * height),
		CoolingFunction: "linear",
		DistanceFunction: "euclidean",
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

func (som *SOM) createNodes(dimensions int) {
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Width + j
			som.Nodes[k] = NewNode(j, i, dimensions)
		}
	}
}

func (som *SOM) InitializeWithRandomValues(matrix *Matrix) {
	som.createNodes(matrix.Columns)

	for _, node := range som.Nodes {
		for i:=0; i< matrix.Columns; i++ {
			r := (matrix.Maximums[i] - matrix.Minimums[i]) + matrix.Minimums[i]
			node.Weights[i] = r * rand.Float64()
		}
	}
}

func (som *SOM) InitializeWithDataPoints(matrix *Matrix) {
	som.createNodes(matrix.Columns)

	for _, node := range som.Nodes {
		copy(node.Weights, matrix.RandomRow())
	}
}

func (som *SOM) Closest(input []float64) *Node {
	nodes := make([]*Node, 0)

	// get initial distance
	t := som.Distance(input, som.Nodes[0].Weights)

	for _, node := range som.Nodes {
		// calculate distance
		d := som.Distance(input, node.Weights)

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
		d1 := som.Distance(input, n1.Weights)
		d2 := som.Distance(input, n2.Weights)

		return d1 < d2
	})

	neighbors := make([]*Node, 0, K)

	for i:=0; i<K; i++ {
		neighbors = append(neighbors, nodes[i])
	}

	return neighbors
}

func (som *SOM) Step(matrix *Matrix, step, steps int, initialLearningRate float64) {
	// calculate position
	pos := float64(step) / float64(steps)

	// calculate learning rate
	learningRate := initialLearningRate * som.CoolingFactor(pos)

	// calculate neighborhood radius
	initialRadius := float64(Max(som.Width, som.Height)) / 2.0
	radius := initialRadius * som.CoolingFactor(pos)

	// get closest node to input
	winningNode := som.Closest(matrix.RandomRow())

	for _, node := range som.Nodes {
		// calculate distance to winner
		distance := som.Distance(winningNode.Position, node.Position)

		// check inclusion in the radius (doubled to fit gaussian function)
		if(distance < radius * 2) {
			// calculate the influence
			i := som.NeighborhoodInfluenceFactor(distance / radius)

			// adjust node
			node.Adjust(winningNode.Weights, i * learningRate)
		}
	}
}

func (som *SOM) Train(matrix *Matrix, steps int, initialLearningRate float64) {
	for step:=0; step<steps; step++ {
		som.Step(matrix, step, steps, initialLearningRate)
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
	radius := som.Distance(input, neighbors[K-1].Weights)
	for i, n := range neighbors {
		distance := som.Distance(input, n.Weights)
		neighborWeights[i] = som.NeighborhoodInfluenceFactor(distance / radius)
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

func (som *SOM) CoolingFactor(input float64) float64 {
	switch som.CoolingFunction {
	case "linear":
		return LinearCooling(input)
	case "soft":
		return SoftCooling(input)
	case "medium":
		return MediumCooling(input)
	case "hard":
		return HardCooling(input)
	}

	return 0.0
}

func (som *SOM) Distance(from, to []float64) float64 {
	switch som.DistanceFunction {
	case "euclidean":
		return EuclideanDistance(from, to)
	case "manhattan":
		return ManhattanDistance(from, to)
	}

	return 0.0
}

func (som *SOM) NeighborhoodInfluenceFactor(distance float64) float64 {
	switch som.NeighborhoodFunction {
	case "bubble":
		return BubbleNeighborhood(distance)
	case "cone":
		return ConeNeighborhood(distance)
	case "gaussian":
		return GaussianNeighborhood(distance)
	case "mexicanthat":
		return MexicanHatNeighborhood(distance)
	}

	return 0.0
}

func (som *SOM) Dimensions() int {
	return len(som.Nodes[0].Weights)
}

func (som *SOM) SaveAsJSON(destination io.Writer) (error) {
	writer := json.NewEncoder(destination)
	return writer.Encode(som)
}
