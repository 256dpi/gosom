package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"

	"github.com/256dpi/gosom"
	"io/ioutil"
)

func main() {
	c := parseConfig()

	if c.prepare {
		doPrepare(c)
	} else if c.train {
		doTrain(c)
	} else if c.plot {
		doPlot(c)
	} else if c.functions {
		doFunctions()
	}
}

func doPrepare(config *config) {
	data := readData(config.data)

	som := gosom.NewSOM(data, config.width, config.height)

	switch config.initialization {
	case "random":
		som.InitializeWithRandomValues()
	case "datapoints":
		som.InitializeWithDataPoints()
	}

	dump, err := json.MarshalIndent(som, "", "  ")
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(config.file, dump, 0644)
}

func doTrain(config *config) {

}

func doPlot(config *config) {

}

func doFunctions(){
	fmt.Printf("plotting cooling functions to './cooling.png' ...\n")
	plotCoolingFunctions("cooling.png")

	fmt.Println("plotting neighborhood functions to './neighborhood.png' ...")
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
