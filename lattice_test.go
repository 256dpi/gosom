package gosom

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewLattice(t *testing.T) {
	l := NewLattice(10, 10, 2)

	for y:=0; y<10; y++ {
		for x:=0; x<10; x++ {
			i := y * 10 + x

			require.Equal(t, []float64{float64(x), float64(y)}, l[i].Position)
			require.Equal(t, []float64{0.0, 0.0}, l[i].Weights)
		}
	}
}

func TestLatticeSorting(t *testing.T) {
	l1 := NewLattice(2, 2, 1)

	for i:=0; i<4; i++ {
		l1[i].Weights[0] = float64(i)
	}

	l2 := l1.Sort(func(n1, n2 *Node) bool {
		return n1.Weights[0] > n2.Weights[0]
	})

	l3 := l1.Sort(func(n1, n2 *Node) bool {
		return n1.Weights[0] < n2.Weights[0]
	})

	require.Equal(t, 3.0, l2[0].Weights[0])
	require.Equal(t, 0.0, l2[3].Weights[0])
	require.Equal(t, 0.0, l3[0].Weights[0])
	require.Equal(t, 3.0, l3[3].Weights[0])
}
