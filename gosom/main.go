package main

import (
	"fmt"
	"os"
	"encoding/json"
	"math"

	"github.com/256dpi/gosom"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/cheggaaa/pb"
	"github.com/gonum/floats"
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
	} else if c.test {
		doTest(c)
	} else if c.functions {
		doFunctions()
	}
}

func doPrepare(config *config) {
	som := gosom.NewSOM(config.width, config.height)
	data := loadData(config.data)

	switch config.initialization {
	case "random":
		som.InitializeWithRandomValues(data)
	case "datapoints":
		som.InitializeWithDataPoints(data)
	}

	storeSOM(config.file, som)
	fmt.Printf("Prepared new SOM and saved to '%s'.\n", config.file)
}

func doTrain(config *config) {
	som := loadSOM(config.file)
	data := loadData(config.data)

	som.DistanceFunction = config.distanceFunction
	som.NeighborhoodFunction = config.neighborhoodFunction
	som.CoolingFunction = config.coolingFunction

	bar := pb.StartNew(config.trainingSteps)

	for step:=0; step<config.trainingSteps; step++ {
		som.Step(data, step, config.trainingSteps, config.initialLearningRate)
		bar.Increment()
	}

	bar.Finish()

	storeSOM(config.file, som)
	fmt.Printf("Trained SOM and saved to '%s'.\n", config.file)
}

func doPlot(config *config) {
	som := loadSOM(config.file)

	dimensions := gosom.DrawDimensions(som, config.size)

	for i, dimension := range dimensions {
		file := fmt.Sprintf("%s/dimension-%d.png", config.directory, i)

		err := draw2dimg.SaveToPngFile(file, dimension)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Plotted dimension to '%s'.\n", file)
	}

	uMatrix := gosom.DrawUMatrix(som, config.size)
	file := fmt.Sprintf("%s/umatrix.png", config.directory)

	err := draw2dimg.SaveToPngFile(file, uMatrix)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Plotted U-Matrix to '%s'.\n", file)
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

func doTest(config *config) {
	som := loadSOM(config.file)
	data := loadData(config.data)
	test := data.SubMatrix(0, config.testDimensions)

	fmt.Println("Classification tests:")
	testHelper(data, test, func(input []float64)([]float64){
		return som.Classify(input)
	})

	fmt.Printf("\nInterpolation tests (K=%d):\n", config.nearestNeighbors)
	testHelper(data, test, func(input []float64)([]float64){
		return som.Interpolate(input, config.nearestNeighbors)
	})

	fmt.Printf("\nWeighted interpolation tests (K=%d):\n", config.nearestNeighbors)
	testHelper(data, test, func(input []float64)([]float64){
		return som.WeightedInterpolate(input, config.nearestNeighbors)
	})
}

func testHelper(data *gosom.Matrix, test *gosom.Matrix, tester func([]float64)([]float64)) {
	errors := make([]float64, data.Rows)

	for i:=0; i<data.Rows; i++ {
		output := tester(test.Data[i])
		errors[i] = gosom.Avg(getErrors(data.Data[i], output, test.Columns))
		fmt.Printf("  %.3f: %.3f (Error: %.2f%%)\n", data.Data[i], output, errors[i])
	}

	fmt.Printf("  Min: %.2f%%, Max: %.2f%%, Avg: %.2f%%\n", floats.Min(errors), floats.Max(errors), gosom.Avg(errors))
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

func getErrors(data, test []float64, offset int) []float64 {
	errors := make([]float64, 0)

	for i:=offset; i<len(data); i++ {
		errors = append(errors, math.Abs((test[i] - data[i]) / data[i] * 100))
	}

	return errors
}
