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

func (s *latticeSorter) Len() int {
	return len(s.lattice)
}

func (s *latticeSorter) Swap(i, j int) {
	s.lattice[i], s.lattice[j] = s.lattice[j], s.lattice[i]
}

func (s *latticeSorter) Less(i, j int) bool {
	return s.sortFunction(s.lattice[i], s.lattice[j])
}
