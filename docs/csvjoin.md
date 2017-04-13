
# USAGE

    csvjoin [OPTIONS] CSV1 CSV2 COL1 COL2

## SYNOPSIS

csvjoin outputs CSV content based on two CSV files with match column values.
Each CSV input file has a designated column to match on. The values are
compared as strings.

## OPTIONS

```
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-v	display version
	-version	display version
```

## EXAMPLES

Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv.

```
    csvjoin data1.csv data2.csv 1 3 > merged-data.csv
```

