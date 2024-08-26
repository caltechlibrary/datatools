%csv2tab(1) user manual | version 1.2.10 55fe3e8
% R. S. Doiel
% 2024-08-26

# NAME

csv2tab 

# SYNOPSIS

csv2tab [OPTIONS]

# DESCRIPTION

csv2tab is a simple conversion utility to convert from CSV to tab separated values.
csv2tab reads from standard input and writes to standard out.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


# EXAMPLES

If my.tab contained

~~~
    "name","email","age"
	"Doe, Jane","jane.doe@example.org",42
~~~

Concert this to a tab separated values

~~~
    csv2tab < my.csv 
~~~

This would yield

~~~
    name	email	age
	Doe, Jane	jane.doe@example.org	42
~~~

csv2tab 1.2.10


