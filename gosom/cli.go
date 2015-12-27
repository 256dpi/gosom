package main

import (
	"github.com/docopt/docopt-go"
	"strconv"
)

type config struct {
	prepare bool
	train bool
	plot bool
	functions bool

	file string
	data string
	width int
	height int
	initialization string
	initialLearningRate float64
	trainingSteps int
}

func parseConfig() *config {
	usage := `Self organizing maps for go.

Usage:
  gosom prepare <file> <data> <width> <height> [-i <im>]
  gosom train <file> <data> [-l <lr> -s <ts>]
  gosom plot <file>
  gosom -h
  gosom -v
  gosom -f

Options:
  -h       Show help.
  -v       Show version.
  -f       Plot functions to current directoy.
  -i <im>  Initialization method (random, datapoints) [default: datapoints].
  -l <lr>  Initial learning rate [default: 0.5].
  -s <ts>  Number of training steps [default: 10000].`

	a, err := docopt.Parse(usage, nil, true, "gosom 0.1", false)
	if err != nil {
		panic(err)
	}

	return &config{
		prepare: getBool(a["prepare"]),
		train: getBool(a["train"]),
		plot: getBool(a["plot"]),
		functions: getBool(a["-f"]),
		file: getString(a["<file>"]),
		data: getString(a["<data>"]),
		width: getInt(a["<width>"]),
		height: getInt(a["<height>"]),
		initialization: getString(a["-i"]),
		initialLearningRate: getFloat(a["-l"]),
		trainingSteps: getInt(a["-s"]),
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
