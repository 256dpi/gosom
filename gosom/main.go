package main

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/256dpi/gosom"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/cheggaaa/pb"
)

func main() {
	c := parseConfig()

	if c.prepare {
		doPrepare(c)
	} else if c.train {
		doTrain(c)
	} else if c.plot {
		doPlot(c)
	} else if c.classify {
		doClassification(c)
	} else if c.interpolate {
		doInterpolation(c)
	} else if c.functions {
		doFunctions()
	}
}

func doPrepare(config *config) {
	som := gosom.NewSOM(config.width, config.height)
	ds := loadData(config.data)

	switch config.initialization {
	case "random":
		som.InitializeWithRandomValues(ds)
	case "datapoints":
		som.InitializeWithDataPoints(ds)
	}

	storeSOM(config.file, som)
	fmt.Printf("Prepared new SOM and saved to '%s'.\n", config.file)
}

func doTrain(config *config) {
	som := loadSOM(config.file)
	ds := loadData(config.data)

	som.DistanceFunction = config.distanceFunction
	som.NeighborhoodFunction = config.neighborhoodFunction
	som.CoolingFunction = config.coolingFunction

	bar := pb.StartNew(config.trainingSteps)

	for step:=0; step<config.trainingSteps; step++ {
		som.Step(ds, step, config.trainingSteps, config.initialLearningRate)
		bar.Increment()
	}

	bar.Finish()

	storeSOM(config.file, som)
	fmt.Printf("Trained SOM and saved to '%s'.\n", config.file)
}

func doPlot(config *config) {
	som := loadSOM(config.file)
	ds := loadData(config.data)

	images := gosom.DrawDimensions(som, ds, config.size)

	for i, img := range images {
		file := fmt.Sprintf("%s/dimension-%d.png", config.directory, i)

		err := draw2dimg.SaveToPngFile(file, img)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Plotted dimension to '%s'.\n", file)
	}
}

func doClassification(config *config) {
	som := loadSOM(config.file)

	input := readInput(config.input)
	fmt.Printf("%f: %f", input, som.Classify(input))
}

func doInterpolation(config *config) {
	som := loadSOM(config.file)

	input := readInput(config.input)

	if config.weighted {
		fmt.Printf("%f: %f", input, som.WeightedInterpolate(input, config.nearestNeighbors))
	} else {
		fmt.Printf("%f: %f", input, som.Interpolate(input, config.nearestNeighbors))
	}
}

func doFunctions(){
	fmt.Printf("Plotting cooling functions to './cooling.png' ...\n")
	plotCoolingFunctions("cooling.png")

	fmt.Println("Plotting neighborhood functions to './neighborhood.png' ...")
	plotNeighborhoodFunctions("neighborhood.png")
}

func loadData(file string) *gosom.Matrix {
	handle, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	ds, err := gosom.LoadMatrixFromCSV(handle)
	if err != nil {
		panic(err)
	}

	return ds
}

func loadSOM(file string) *gosom.SOM {
	handle, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	som, err := gosom.LoadSOMFromJSON(handle)
	if err != nil {
		panic(err)
	}

	return som
}

func storeSOM(file string, som *gosom.SOM) {
	handle, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	err = som.SaveAsJSON(handle)
	if err != nil {
		panic(err)
	}
}

func readInput(input string) []float64 {
	floats := make([]float64, 0)

	err := json.Unmarshal([]byte(input), &floats)
	if err != nil {
		panic(err)
	}

	return floats
}
