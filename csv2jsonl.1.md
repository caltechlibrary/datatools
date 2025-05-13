%csv2jsonl(1) user manual | version 1.3.2 1ee0728
% R. S. Doiel
% 2025-05-13

# NAME

csv2jsonl

# SYNOPSIS

csv2jsonl [OPTIIONS]

# DESCRIPTION

csv2jsonl reads CSV from stdin and writes a JSON-L to stdout. JSON output
is one object per line. See https://jsonlines.org.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set the delimter character

-examples
: display example(s)

-fields-per-record
: Set the number of fields expected in the CSV read, -1 to turn off

-i, -input
: input filename

-nl, -newline
: include trailing newline in output

-o, -output
: output filename

-quiet
: suppress error output

-reuse-record
: reuse the backing array

-trim-leading-space
: trim leading space in fields for CSV input

-use-header
: treat the first row as field names

-use-lazy-quotes
: use lazy quotes for for CSV input


# EXAMPLES

Convert data1.csv to data1.jsonl using Unix pipes.

~~~
    cat data1.csv | csv2jsonl > data1.jsonl
~~~

Convert data1.csv to JSON line (one object line per blob)

~~~
    csv2jsonl data1.csv
~~~


