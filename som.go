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
	Nodes []*Node

	Data [][]float64
	Rows int
	Columns int

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
		DistanceFunction: EuclideanDistance,
		NeighborhoodFunction: ConeNeighborhood,
	}

	// create nodes
	for i:=0; i<som.Height; i++ {
		for j:=0; j<som.Width; j++ {
			k := i * som.Height + j
			som.Nodes[k] = NewNode(j, i, som.Columns)
		}
	}

	return som
}

func (som *SOM) InitializeWithRandomValues() {
	min := make([]float64, som.Columns)
	max := make([]float64, som.Columns)

	for j:=0; j<som.Rows; j++ {
		for i:=0; i<som.Columns; i++ {
			min[i] = math.Min(min[i], som.Data[j][i])
			max[i] = math.Max(max[i], som.Data[j][i])
		}
	}

	for _, node := range som.Nodes {
		node.Weights = make([]float64, som.Columns)
		for i:=0; i<som.Columns; i++ {
			node.Weights[i] = rand.Float64() * (max[i] - min[i]) + min[i]
		}
	}
}

func (som *SOM) InitializeWithRandomDataPoints() {
	for _, node := range som.Nodes {
		node.Weights = make([]float64, som.Columns)
		copy(node.Weights, som.Data[rand.Intn(som.Rows-1)])
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
			t = d;
			n = append([]*Node{}, node)
		} else if(d <= t) {
			// add winner
			n = append(n, node)
		}
	}

	if len(n) > 1 {
		// return random winner
		return n[rand.Intn(len(n)-1)]
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
		dataPoint := som.Data[rand.Intn(som.Rows-1)]

		// get closest node to input
		winningNode := som.Closest(dataPoint);

		for _, node := range som.Nodes {
			// calculate distance to winner
			distance := som.DistanceFunction(winningNode.Position, node.Position)

			// check inclusion in the radius (doubled to fit gaussian function)
			if(distance < radius * 2) {
				// calculate the influence
				i := som.NeighborhoodFunction(distance / radius)

				// adjust node
				node.Adjust(winningNode.Weights, i * learningRate);
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
			total[j] += neighbors[i].Weights[j];
		}
	}

	// calculate average
	for i:=0; i<som.Columns; i++ {
		total[i] = total[i] / float64(K);
	}

	return total;
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
