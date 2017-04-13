# USAGE

    csvcols [OPTIONS] ARGS_AS_COLS

## SYNOPSIS

csvcols converts a set of command line args into columns output in CSV format.
It can also be used to filter input CSV and rendering only the column numbers
listed on the commandline.

## OPTIONS

```
	-d	set delimiter for conversion
	-delimiter	set delimiter for conversion
	-f	filter CSV input for columns requested
	-filter-columns	filter CSV input for columns requested
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-v	display version
	-version	display version
```

## EXAMPLES

Simple usage of building a CSV file one row at a time.

```
    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv
```

Example parsing a pipe delimited string into a CSV line

```
    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
```

Filter a 10 column CSV file for columns 0,3,5 (left most column is number zero)

```
	cat 10col.csv | csvcols -f 0 3 5 > 3col.csv
```
