#!/bin/bash

echo "Creating 2cols.csv"
cat mysheet.csv | csvcleaner -field-per-row=2 > 2cols.csv
echo "Creating 3cols.csv"
cat mysheet.csv | csvcleaner -field-per-row=3 > 3cols.csv
echo "Creating ltrim.csv"
cat mysheet.csv | csvcleaner -left-trim-spaces > ltrim.csv
echo "Creating rtrim.csv"
cat mysheet.csv | csvcleaner -right-trim-spaces > rtrim.csv
echo "Creating trim.csv"
cat mysheet.csv | csvcleaner -trim-spaces > trim.csv
