package main

import (
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/256dpi/gosom"
)

type config struct {
	prepare bool
	train bool
	classify bool
	interpolate bool
	plot bool
	functions bool

	file string
	directory string
	data string
	width int
	height int
	initialization string
	initialLearningRate float64
	trainingSteps int
	distanceFunction gosom.DistanceFunction
	neighborhoodFunction gosom.NeighborhoodFunction
	coolingFunction gosom.CoolingFunction
	input string
	weighted bool
	nearestNeighbors int
	size int
}

func parseConfig() *config {
	usage := `Self organizing maps for go.

Usage:
  gosom prepare <file> <data> <width> <height> [-i <im>]
  gosom train <file> <data> [-l <lr> -t <ts> -d <df> -n <nf> -c <cf>]
  gosom classify <file> <input> [-d <df>]
  gosom interpolate <file> <input> [-d <df> -n <nf> -w -k <nn>]
  gosom plot <file> <directory> [-s <ns>]
  gosom -h
  gosom -v
  gosom -f

Options:
  -h       Show help.
  -v       Show version.
  -f       Plot functions to current directoy.
  -i <im>  Initialization method (random, datapoints) [default: datapoints].
  -l <lr>  Initial learning rate [default: 0.5].
  -t <ts>  Number of training steps [default: 10000].
  -d <df>  Distance function (euclidean, manhattan) [default: euclidean].
  -n <nf>  Neighborhood function (bubble, cone, gaussian, mexicanhat) [default: cone].
  -c <cf>  Cooling function (linear, soft, medium, hard) [default: linear].
  -k <nn>  Number of nearest neighbors to consider [default: 5].
  -w       Use weighted interpolation.
  -s <ns>  Size of the individual nodes [default: 10].`

	a, err := docopt.Parse(usage, nil, true, "gosom 0.1", false)
	if err != nil {
		panic(err)
	}

	return &config{
		prepare: getBool(a["prepare"]),
		train: getBool(a["train"]),
		classify: getBool(a["classify"]),
		interpolate: getBool(a["interpolate"]),
		plot: getBool(a["plot"]),
		functions: getBool(a["-f"]),
		file: getString(a["<file>"]),
		directory: getString(a["<directory>"]),
		data: getString(a["<data>"]),
		width: getInt(a["<width>"]),
		height: getInt(a["<height>"]),
		initialization: getString(a["-i"]),
		initialLearningRate: getFloat(a["-l"]),
		trainingSteps: getInt(a["-t"]),
		distanceFunction: getDistanceFunction(a["-d"]),
		neighborhoodFunction: getNeighborhoodFunction(a["-n"]),
		coolingFunction: getCoolingFunction(a["-c"]),
		input: getString(a["<input>"]),
		weighted: getBool(a["-w"]),
		nearestNeighbors: getInt(a["-k"]),
		size: getInt(a["-s"]),
	}
}

func getBool(v interface{}) bool {
	b, ok := v.(bool)

	if !ok {
		return false
	} else {
		return b
	}
}

func getString(v interface{}) string {
	s, ok := v.(string)

	if !ok {
		return ""
	} else {
		return s
	}
}

func getInt(v interface{}) int {
	s := getString(v)

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func getFloat(v interface{}) float64 {
	s := getString(v)

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	} else {
		return f
	}
}

func getDistanceFunction(v interface{}) gosom.DistanceFunction {
	s := getString(v)

	switch s {
	case "euclidean":
		return gosom.EuclideanDistance
	case "manhattan":
		return gosom.ManhattanDistance
	}

	return gosom.EuclideanDistance
}

func getNeighborhoodFunction(v interface{}) gosom.NeighborhoodFunction {
	s := getString(v)

	switch s {
	case "bubble":
		return gosom.BubbleNeighborhood
	case "cone":
		return gosom.ConeNeighborhood
	case "gaussian":
		return gosom.GaussianNeighborhood
	case "mexicanhat":
		return gosom.MexicanHatNeighborhood
	}

	return gosom.ConeNeighborhood
}

func getCoolingFunction(v interface{}) gosom.CoolingFunction {
	s := getString(v)

	switch s {
	case "linear":
		return gosom.LinearCooling
	case "soft":
		return gosom.SoftCooling
	case "medium":
		return gosom.MediumCooling
	case "hard":
		return gosom.HardCooling
	}

	return gosom.LinearCooling
}
