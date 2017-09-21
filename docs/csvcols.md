
# csvcols

## USAGE

    csvcols [OPTIONS] ARGS_AS_COLS

## SYNOPSIS

csvcols converts a set of command line args into columns output in CSV format.
It can also be used to filter input CSV and rendering only the column numbers
listed on the commandline (first column is 1 not 0)

## OPTIONS

```
    -col        filter CSV input for columns requested
    -cols       filter CSV input for columns requested
    -d          set delimiter character
    -delimiter  set delimiter character
    -h          display help
    -help       display help
    -i          input filename
    -input      input filename
    -l          display license
    -license    display license
    -o          output filename
    -output     output filename
    -v          display version
    -version    display version
```

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

Filter a 10 column CSV file for columns 1,4,6 (left most column is number zero)

```shell
    cat 10col.csv | csvcols -col 1,4,6 > 3col.csv
```

Filter a 10 columns CSV file for columns 1,4,6 from input file

```shell
    csvcols -i 10col.csv -col 1,4,6 > 3col.csv
```


