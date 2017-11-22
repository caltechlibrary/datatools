
# USAGE

## csvcols [OPTIONS] [ARGS_AS_COL_VALUES]

## SYNOPSIS

csvcols converts a set of command line args into columns output in CSV format.
It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).

## OPTIONS

	-col	output specified columns (e.g. -col 1,12:14,2,4))
	-cols	output specified columns (e.g. -col 1,12:14,2,4))
	-d	set delimiter character
	-delimiter	set delimiter character
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-skip-header-row	skip the header row
	-uuid	add a prefix row with generated UUID cell
	-v	display version
	-version	display version

## EXAMPLES

Simple usage of building a CSV file one row at a time.

```shell
    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv
```

Example parsing a pipe delimited string into a CSV line

```shell
    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
```

Filter a 10 column CSV file for columns 1,4,6 (left most column is one)

```shell
    cat 10col.csv | csvcols -col 1,4,6 > 3col.csv
```


Filter a 10 columns CSV file for columns 1,4,6 from file named "10col.csv"

```shell
    csvcols -i 10col.csv -col 1,4,6 > 3col.csv
```


csvcols v0.0.17
