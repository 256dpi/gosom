package gosom

type Node struct {
	Position []float64
	Weights []float64
}

func NewNode(x, y, dimensions int) *Node {
	return &Node{
		Position: []float64{ float64(x), float64(y) },
		Weights: make([]float64, dimensions),
	}
}

func (n *Node) Adjust(input []float64, influence float64) {
	l := Min(len(input), len(n.Weights))

	for i:=0; i<l; i++ {
		n.Weights[i] += (input[i] - n.Weights[i]) * influence
	}
}
