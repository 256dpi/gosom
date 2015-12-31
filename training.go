package gosom

// A Training holds settings for a SOM training.
type Training struct {
	SOM                 *SOM
	Steps               int
	InitialLearningRate float64
	FinalLearningRate   float64
	InitialRadius       float64
	FinalRadius         float64
}

// NewTraining returns a new Training.
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

// Progress returns the current progress based on step and steps.
func (t *Training) Progress(step int) float64 {
	return float64(step) / float64(t.Steps)
}

// LearningRate calculates the current learning rate.
func (t *Training) LearningRate(step int) float64 {
	r := t.InitialLearningRate - t.FinalLearningRate
	return r * t.SOM.CF(t.Progress(step)) + t.FinalLearningRate
}

// Radius calculates the current radius.
func (t *Training) Radius(step int) float64 {
	r := t.InitialRadius - t.FinalRadius
	return r * t.SOM.CF(t.Progress(step)) + t.FinalRadius
}
