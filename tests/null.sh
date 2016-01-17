#!/usr/bin/env bash
set -e

# This test trains the SOM using an advanced data set containing NULLs.
# The first rows try to train a binary switch while the other rows try to train
# a linear data set. The SOM is expected to linearly classify and interpolate
# and return 0 if the switch is on.

echo "---> cleaning and creating './tmp'"
rm -rf ./tmp
mkdir -p ./tmp
cd ./tmp

echo "---> write data set"
cat <<EOF > ./data.csv
NULL,1.00,1.00
NULL,0.00,0.00
0.00,NULL,0.00
0.25,NULL,0.25
0.50,NULL,0.50
0.75,NULL,0.75
1.00,NULL,1.00
EOF

echo "---> write test set"
cat <<EOF > ./test.csv
0.00,0.00,0.00
0.25,0.00,0.00
0.75,0.00,0.00
1.00,0.00,0.00
0.00,1.00,0.00
0.25,1.00,0.25
0.50,1.00,0.50
0.75,1.00,0.75
1.00,1.00,1.00
EOF

echo "---> preparing SOM"
../gosom prepare som.json data.csv 30 30 -n gaussian -c soft -i random

echo "---> training SOM"
../gosom train som.json data.csv -t 100000

echo "---> plotting trained SOM"
../gosom plot som.json . -p trained

echo "---> tuning SOM"
../gosom train som.json data.csv -t 100000 -l 0.05 -r 10

echo "---> plotting tuned SOM"
../gosom plot som.json . -p tuned

echo "---> testing SOM"
../gosom test som.json test.csv -k 10

echo "---> opening folder"
open .
