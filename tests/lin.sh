#!/usr/bin/env bash

echo "---> cleaning and creating './tmp'"
rm -rf ./tmp
mkdir -p ./tmp
cd ./tmp

echo "---> copy data set"
cp ../lin.csv ./data.csv
cp ../lin.test.csv ./test.csv

echo "---> preparing SOM"
../gosom prepare som.json data.csv 100 100

echo "---> training SOM"
../gosom train som.json data.csv

echo "---> testing SOM"
../gosom test som.json test.csv -k 25

echo "---> plotting SOM"
../gosom plot som.json . -p lin

echo "---> opening folder"
open .
