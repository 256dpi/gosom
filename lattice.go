package gosom

import "sort"

// A Lattice is a collection of nodes arranged in a two dimensional space.
type Lattice []*Node

// NewLattice returns an initialized lattice.
func NewLattice(width, height, dimensions int) Lattice {
	lattice := make(Lattice, width*height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k := i*width + j
			lattice[k] = NewNode(j, i, dimensions)
		}
	}

	return lattice
}

// Sort returns a sorted lattice using the sorting function.
func (l Lattice) Sort(sortFunction func(n1, n2 *Node) bool) Lattice {
	cpy := make(Lattice, len(l))
	copy(cpy, l)

	ns := &latticeSorter{
		lattice:      cpy,
		sortFunction: sortFunction,
	}

	sort.Sort(ns)

	return cpy
}

type latticeSorter struct {
	lattice      Lattice
	sortFunction func(n1, n2 *Node) bool
}

func (ns *latticeSorter) Len() int {
	return len(ns.lattice)
}

func (ns *latticeSorter) Swap(i, j int) {
	ns.lattice[i], ns.lattice[j] = ns.lattice[j], ns.lattice[i]
}

func (ns *latticeSorter) Less(i, j int) bool {
	return ns.sortFunction(ns.lattice[i], ns.lattice[j])
}
