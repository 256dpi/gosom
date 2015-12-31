#!/usr/bin/env bash

echo "---> cleaning and creating './tmp'"
rm -rf ./tmp
mkdir -p ./tmp
cd ./tmp

echo "---> downloading data set"
wget https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data -O data.csv

echo "---> preprocessing data set"
sed -i.bak s/Iris-setosa/1/g data.csv
sed -i.bak s/Iris-versicolor/2/g data.csv
sed -i.bak s/Iris-virginica/3/g data.csv
rm *.bak

echo "---> preparing SOM"
../gosom prepare som.json data.csv 50 50

echo "---> training SOM"
../gosom train som.json data.csv

echo "---> testing SOM"
../gosom test som.json data.csv -k 25

echo "---> plotting SOM"
../gosom plot som.json . -p iris

echo "---> opening folder"
open .
