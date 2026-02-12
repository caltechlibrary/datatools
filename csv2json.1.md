%csv2json(1) user manual | version 1.3.5 effbad2
% R. S. Doiel
% 2026-02-12

# NAME

csv2json

# SYNOPSIS

csv2json [OPTIIONS]

# DESCRIPTION

csv2json reads CSV from stdin and writes a JSON to stdout. JSON output
can be either an array of JSON blobs or one JSON blob (row as object)
per line.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-as-blobs
: output as one JSON blob per line

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

Convert data1.csv to data1.json using Unix pipes.

~~~
    cat data1.csv | csv2json > data1.json
~~~

Convert data1.csv to JSON blobs, one line per blob

~~~
    csv2json -as-blobs -i data1.csv
~~~


