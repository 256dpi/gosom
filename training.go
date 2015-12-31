package gosom

type Training struct {
	Steps int
	InitialLearningRate float64
	FinalLearningRate float64
	InitialRadius float64
	FinalRadius float64
}

func NewTraining(steps int, ilr, flr, ir, fr float64) *Training {
	return &Training{
		Steps: steps,
		InitialLearningRate: ilr,
		FinalLearningRate: flr,
		InitialRadius: ir,
		FinalRadius: fr,
	}
}

func (t *Training) Progress(step int) float64 {
	return float64(step) / float64(t.Steps)
}
