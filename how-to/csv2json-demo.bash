#!/bin/bash

echo "as pipeline --> data1.json"
cat data1.csv | csv2json > data1.json
cat data1.json

echo "using options `-as-blobs -i=data1.csv`"
csv2json -as-blobs -i data1.csv
cat data1.json
