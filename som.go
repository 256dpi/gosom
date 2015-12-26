package gosom

import "fmt"

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
	for i:=0; i<height; i++ {
		for j:=0; j<width; j++ {
			k := i * height + j
			som.Nodes[k] = &Node{ X: j, Y: i}
		}
	}

	return som
}

func (som *SOM) Prepare(initialization Initialization) {
	d := len(som.Data[0])

	switch initialization {
	case ZeroInitialization:
		for _, n := range som.Nodes {
			n.Weights = make([]float64, d)
		}
	case RandomInitialization:
		/*for _, n := range som.Nodes {

		}*/
	case RandomDataInitialization:

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
			s += fmt.Sprintf("%v ", som.Nodes[k].Weights)
		}

		s += "\n"
	}

	return s
}
