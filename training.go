package gosom

type Training struct {
	SOM                 *SOM
	Steps               int
	InitialLearningRate float64
	FinalLearningRate   float64
	InitialRadius       float64
	FinalRadius         float64
}

func NewTraining(som *SOM, steps int, ilr, flr, ir, fr float64) *Training {
	return &Training{
		SOM:                 som,
		Steps:               steps,
		InitialLearningRate: ilr,
		FinalLearningRate:   flr,
		InitialRadius:       ir,
		FinalRadius:         fr,
	}
}

func (t *Training) Progress(step int) float64 {
	return float64(step) / float64(t.Steps)
}

func (t *Training) LearningRate(step int) float64 {
	return t.InitialLearningRate * t.SOM.CF(t.Progress(step))
}

func (t *Training) Radius(step int) float64 {
	return t.InitialRadius * t.SOM.CF(t.Progress(step))
}
