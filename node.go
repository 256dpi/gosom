package gosom

type Node struct {
	X, Y int
	Weights []float64
}

func NewNode(x, y int) *Node {
	return &Node{
		X: x,
		Y: y,
	}
}

func (n *Node) Adjust(input []float64, influence float64) {
	l := Min(len(input), len(n.Weights))

	for i:=0; i<l; i++ {
		n.Weights[i] += (input[i] - n.Weights[i]) * influence
	}
}
