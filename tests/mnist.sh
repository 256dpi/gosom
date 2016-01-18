#!/usr/bin/env bash
set -e

# This test trains the SOM using the well known MNIST data set.

echo "---> cleaning and creating './tmp'"
rm -rf ./tmp
mkdir -p ./tmp
cd ./tmp

echo "---> downloading data set"
wget http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz
wget http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz
wget http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz
wget http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz

echo "---> unzip data set"
gunzip ./*

echo "---> convert data set"
python ../mnist.py

echo "---> remove intermediary files"
rm ./*-ubyte

echo "---> preparing SOM"
../gosom prepare som.json train.csv 28 28 -n gaussian -c soft

echo "---> training SOM"
../gosom train som.json train.csv -t 10000

echo "---> plotting trained SOM"
../gosom plot som.json .

echo "---> testing SOM"
../gosom test som.json test.csv -k 5 -q

echo "---> opening folder"
open .
