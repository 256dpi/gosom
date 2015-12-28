package gosom

import "sort"

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

func (n *Node) X() int {
	return int(n.Position[0])
}

func (n *Node) Y() int {
	return int(n.Position[1])
}

func (n *Node) Adjust(input []float64, influence float64) {
	l := Min(len(input), len(n.Weights))

	for i:=0; i<l; i++ {
		n.Weights[i] += (input[i] - n.Weights[i]) * influence
	}
}

type NodeSortFunction func(n1, n2 *Node) bool

type nodeSorter struct {
	nodes []*Node
	sortFunction NodeSortFunction
}

func (ns *nodeSorter) Len() int {
	return len(ns.nodes)
}

func (ns *nodeSorter) Swap(i, j int) {
	ns.nodes[i], ns.nodes[j] = ns.nodes[j], ns.nodes[i]
}

func (ns *nodeSorter) Less(i, j int) bool {
	return ns.sortFunction(ns.nodes[i], ns.nodes[j])
}

func SortNodes(nodes []*Node, sortFunction NodeSortFunction) {
	ns := &nodeSorter{
		nodes: nodes,
		sortFunction: sortFunction,
	}

	sort.Sort(ns)
}
