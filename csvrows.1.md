%csvrows(1) user manual | version 1.2.12 1128bff
% R. S. Doiel
% 2024-11-07

# NAME

csvrows 

# SYNOPSIS

csvrows [OPTIONS] [ARGS_AS_ROW_VALUES]

# DESCRIPTION

csvrows converts a set of command line args into rows of CSV
formated output.  It can also be used to filter or list specific rows
of CSV input The first row is 1 not 0. Often row 1 is the header row 
and csvrows makes it easy to output only the data rows.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter character

-header
: display the header row (alias for '-rows 1')

-i, -input
: input filename

-o, -output
: output filename

-quiet
: suppress error messages

-random
: return N randomly selected rows

-row, -rows
: output specified rows in order (e.g. -row 1,5,2-4))

-skip-header-row
: skip the header row (alias for -row 2-

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: use lazy quotes for CSV input


# EXAMPLES

Simple usage of building a CSV file one rows at a time.

~~~
    csvrows "First,Second,Third" "one,two,three" > 4rows.csv
    csvrows "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv
~~~

Example parsing a pipe delimited string into a CSV line

~~~
    csvrows -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    csvrows -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv
~~~

Filter a 10 row CSV file for rows 1,4,6 (top most row is one)

~~~
    cat 10row.csv | csvrows -row 1,4,6 > 3rows.csv
~~~

Filter a 10 row CSV file for rows 1,4,6 from file named "10row.csv"

~~~
    csvrows -i 10row.csv -row 1,4,6 > 3rows.csv
~~~

Filter 3 randomly selected rows from 10row.csv rendering new CSV with
a header row from 10row.csv.

~~~
	csvrows -i 10row.csv -header=true -random=3
~~~

csvrows 1.2.12


