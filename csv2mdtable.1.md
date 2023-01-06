---
title: "csv2mdtable (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-06
---

#NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} reads CSV from stdin and writes a Github Flavored Markdown
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
    cat data1.csv | {app_name} > data1.md
~~~

Convert data1.csv to data1.md using options.

~~~
    {app_name} -i data1.csv -o data1.md
~~~

{app_name} {version}

