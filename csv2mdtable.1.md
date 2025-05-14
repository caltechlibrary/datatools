%csv2mdtable(1) user manual | version 1.3.3 24eb061
% R. S. Doiel
% 2025-05-14

# NAME

csv2mdtable

# SYNOPSIS

csv2mdtable [OPTIONS]

# DESCRIPTION

csv2mdtable reads CSV from stdin and writes a Github Flavored Markdown
table to stdout.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter character

-i, -input
: input filename

-nl, -newline
: if true include leading/trailing newline

-o, -output
: output filename

-quiet
: suppress error message

-trim-leading-space
: trim leading space in field(s) for CSV input

-use-lazy-quotes
: using lazy quotes for CSV input


# EXAMPLES

Convert data1.csv to data1.md using Unix pipes.

~~~
    cat data1.csv | csv2mdtable > data1.md
~~~

Convert data1.csv to data1.md using options.

~~~
    csv2mdtable -i data1.csv -o data1.md
~~~

csv2mdtable 1.3.3


