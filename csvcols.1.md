%csvcols(1) user manual | version 1.2.9 1b11c42
% R. S. Doiel
% 2024-07-09

# NAME

csvcols

# SYNOPSIS

csvcols [OPTIONS] [ARGS_AS_COL_VALUES]

# DESCRIPTION

csvcols converts a set of command line args into columns output in CSV format.

It can also be used CSV input rows and rendering only the column numbers
listed on the commandline (first column is 1 not 0).

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-col, -cols
: output specified columns (e.g. -col 1,12:14,2,4))

-d, -delimiter
: set the input delimiter character

-i, -input
: input filename

-o, -output
: output filename

-od, -output-delimiter
: set the output delimiter character

-quiet
: suppress error messages

-skip-header-row
: skip the header row

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: use lazy quotes on CSV input

-uuid
: add a prefix row with generated UUID cell


# EXAMPLES

Simple usage of building a CSV file one row at a time.

~~~
    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv
~~~

Example parsing a pipe delimited string into a CSV line

~~~
    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
~~~

Using a pipe filter a 3 column CSV for columns 1 and 3 into 2col.csv

~~~
    cat 3col.csv | csvcols -col 1,3 > 2col.csv
~~~


Using options filter a 3 column CSV file for columns 1,3 into 2col.csv

~~~
    csvcols -i 3col.csv -col 1,3 -o 2col.csv
~~~

csvcols 1.2.9


