%tab2csv(1) user manual | version 1.3.4 4312aaa
% R. S. Doiel
% 2025-05-15

# NAME

tab2csv 

# SYNOPSIS

tab2csv [OPTIONS]

# DESCRIPTION

tab2csv is a simple conversion utility to convert from tabs to quoted CSV.

tab2csv reads from standard input and writes to standard out.


# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-fields-per-record
: (int) sets the number o fields expected in each row, -1 turns this off

-reuse-record
: re-uses the backing array on reader

-trim-leading-space
: trims leading space read

-use-lazy-quotes
: use lazy quoting for reader

-crlf
: use CRLF for end of line (EOL) on write, defaults to true for Windows

# EXAMPLES

If my.tab contained

~~~
    name	email	age
	Doe, Jane	jane.doe@example.org	42
~~~

Concert this to a CSV file format

~~~
    tab2csv < my.tab 
~~~

This would yield

~~~
    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42
~~~

tab2csv 1.3.4


