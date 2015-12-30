package gosom

// A Node is a single neuron in a self organizing map.
type Node struct {
	Position []float64
	Weights  []float64
}

// NewNode returns a new Node.
func NewNode(x, y, dimensions int) *Node {
	return &Node{
		Position: []float64{float64(x), float64(y)},
		Weights:  make([]float64, dimensions),
	}
}

// X returns the x coordinate of the node.
func (n *Node) X() int {
	return int(n.Position[0])
}

// Y returns the y coordinate of the node.
func (n *Node) Y() int {
	return int(n.Position[1])
}

// Adjust makes the node more alike to the input based on the influence.
func (n *Node) Adjust(input []float64, influence float64) {
	l := min(len(input), len(n.Weights))

	for i := 0; i < l; i++ {
		n.Weights[i] += (input[i] - n.Weights[i]) * influence
	}
}
