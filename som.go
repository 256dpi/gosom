package gosom

type SOM struct {
	Width int
	Height int

	Data [][]float64
	Nodes []Node

	CoolingFunction CoolingFunction
	DistanceFunction DistanceFunction
	NeighborhoodFunction NeighborhoodFunction
}

func NewSOM(data [][]float64, width, height int) *SOM {
	return &SOM{
		Width: width,
		Height: height,
		Data: data,
		Nodes: make([]Node, width * height),
	}
}

func (som *SOM) Initialize() {

}

func (som *SOM) Train(steps int, initialLearningRate float64) {

}

func (som *SOM) Predict(input []float64) []float64 {
	return []float64{ 0.0 }
}
