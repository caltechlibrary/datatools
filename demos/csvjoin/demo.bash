#!/bin/bash

csvjoin -csv1=data1.csv -col1=2 \
   -csv2=data2.csv -col2=4 \
   -output=merged-data.csv
