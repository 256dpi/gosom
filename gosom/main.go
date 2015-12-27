package main

import (
	"fmt"
	"os"
	"encoding/csv"
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
	som.LoadData(readData(config.data))

	switch config.initialization {
	case "random":
		som.InitializeWithRandomValues()
	case "datapoints":
		som.InitializeWithDataPoints()
	}

	err := gosom.Store(som, config.file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Prepared new SOM and saved to '%s'.\n", config.file)
}

func doTrain(config *config) {
	som, err := gosom.Load(config.file)
	if err != nil {
		panic(err)
	}

	som.LoadData(readData(config.data))
	bar := pb.StartNew(config.trainingSteps)

	for step:=0; step<config.trainingSteps; step++ {
		som.Step(step, config.trainingSteps, config.initialLearningRate)
		bar.Increment()
	}

	bar.FinishPrint("Done!")

	err = gosom.Store(som, config.file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Trained SOM and saved to '%s'.\n", config.file)
}

func doPlot(config *config) {
	som, err := gosom.Load(config.file)
	if err != nil {
		panic(err)
	}

	images := som.DimensionImages(config.size)

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
	som, err := gosom.Load(config.file)
	if err != nil {
		panic(err)
	}

	input := readInput(config.input)
	fmt.Printf("%f: %f", input, som.Classify(input))
}

func doInterpolation(config *config) {
	som, err := gosom.Load(config.file)
	if err != nil {
		panic(err)
	}

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

func readData(file string) [][]float64 {
	handle, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	reader := csv.NewReader(handle)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	floats := make([][]float64, len(data))

	for i, row := range data {
		floats[i] = make([]float64, len(row))

		for j, col := range row {
			floats[i][j] = getFloat(col)
		}
	}

	return floats
}

func readInput(input string) []float64 {
	floats := make([]float64, 0)

	err := json.Unmarshal([]byte(input), &floats)
	if err != nil {
		panic(err)
	}

	return floats
}
