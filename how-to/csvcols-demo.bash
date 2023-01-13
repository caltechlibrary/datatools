#!/bin/bash

echo "Create three columns (words)"
csvcols one two three > 3col.csv
echo "Create three columns (numeric)"
csvcols 1 2 3 >> 3col.csv
echo "Displaying results"
cat 3col.csv
echo "Creating three columns (pipe delimited words)"
csvcols -d "|" "one|two|three" > 3col.csv
echo "Creating three columns (pipe delimited numbers)"
csvcols -delimiter "|" "1|2|3" >> 3col.csv
echo "Display results"
cat 3col.csv
echo "Piping columns 1,3 from 3col.csv to 2col.csv"
cat 3col.csv | csvcols -col 1,3 > 2col.csv
echo "Display results"
cat 2col.csv
echo "Using options, extract columns 1,3 from 3col.csv"
csvcols -i 3col.csv -col 1,3 > 2col.csv
echo "Display results"
cat 2col.csv
