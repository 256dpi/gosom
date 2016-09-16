package gosom

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTraining(t *testing.T) {
	som := NewSOM(5, 5)
	tr := NewTraining(som, 10, 0.5, 0.0, 10.0, 0.0)

	require.Equal(t, 10, tr.Steps)
	require.Equal(t, 0.5, tr.InitialLearningRate)
	require.Equal(t, 0.0, tr.FinalLearningRate)
	require.Equal(t, 10.0, tr.InitialRadius)
	require.Equal(t, 0.0, tr.FinalRadius)

	require.Equal(t, 0.5, tr.Progress(5))
	require.Equal(t, 0.25, tr.LearningRate(5))
	require.Equal(t, 5.0, tr.Radius(5))
}
