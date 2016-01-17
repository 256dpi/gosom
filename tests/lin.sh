#!/usr/bin/env bash

echo "---> cleaning and creating './tmp'"
rm -rf ./tmp
mkdir -p ./tmp
cd ./tmp

echo "---> write data set"
cat <<EOF > ./data.csv
0.0,4.0
1.0,3.0
2.0,2.0
3.0,1.0
4.0,0.0
EOF

echo "---> write test set"
cat <<EOF > ./test.csv
0.5,3.5
1.5,2.5
2.5,1.5
3.5,0.5
EOF

echo "---> preparing SOM"
../gosom prepare som.json data.csv 100 100

echo "---> training SOM"
../gosom train som.json data.csv

echo "---> plotting SOM"
../gosom plot som.json . -p lin

echo "---> testing SOM"
../gosom test som.json test.csv -k 25

echo "---> opening folder"
open .
