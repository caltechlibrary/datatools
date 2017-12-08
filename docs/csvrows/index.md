
# USAGE

	csvrows [OPTIONS] [ARGS_AS_ROW_VALUES]

## SYNOPSIS


csvrows converts a set of command line args into rows of CSV formated output.
It can also be used to filter or list specific rows of CSV input
The first row is 1 not 0. Often row 1 is the header row and %!s(MISSING) makes it
easy to output only the data rows.


## OPTIONS

```
    -d, -delimiter            set delimiter character
    -examples                 display example(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -header                   display the header row (alias for '-rows 1')
    -i, -input                input filename
    -l, -license              display license
    -o, -output               output filename
    -quiet                    suppress error messages
    -row, -rows               output specified rows in order (e.g. -row 1,5,2:4))
    -skip-header-row          skip the header row (alias for -row 2:
    -v, -version              display version
```


## EXAMPLES


Simple usage of building a CSV file one rows at a time.

    csvrows "First,Second,Third" "one,two,three" > 4rows.csv
    csvrows "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv

Example parsing a pipe delimited string into a CSV line

    csvrows -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    csvrows -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv

Filter a 10 row CSV file for rows 1,4,6 (top most row is one)

    cat 10row.csv | csvrows -row 1,4,6 > 3rows.csv

Filter a 10 row CSV file for rows 1,4,6 from file named "10row.csv"

    csvrows -i 10row.csv -row 1,4,6 > 3rows.csv


csvrows v0.0.20-pre
