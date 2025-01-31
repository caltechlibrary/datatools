%jsonobjects2csv(1) user manual | version 1.2.12 eb5bc10
% R. S. Doiel
% 2025-01-31

# NAME

jsonobjects2csv

# SYNOPSIS

jsonobjects2csv [OPTIONS] [JSON_FILENAME] [YAML_FILENAME]

# DESCRIPTION

jsonobjects2csv is a tool that converts a JSON list of objects into CSV output.

jsonobjects2csv will take JSON expressing a list of objects and turn them into a CSV
representation. If the object's attributes include other objects or arrays they
are rendered as YAML in the cell of the csv output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages

-delimiter
: set the CSV column delimiter for output

-show-header
: set whether or not to output a header row at start of outout.

-i FILENAME
: Use  FILENAME for input, "-" will be interpreted as standard input

-o FILENAME
: Use FILENAME for ouput, "-" will be interpreted as standard output


# EXAMPLES

Used by typing into standard in (press Ctrl-d to end your input).

~~~shell
	jsonobjects2csv
	[
	  {"one": 1, "two": 2},
	  {"one": 10, "two": 20},
    ]
	^D
~~~

This should yield the following.

~~~text
one,two
1,2
10,20
~~~

These would get the file named "my_list.json" and save it as my.csv

~~~shell
    jsonobjects2csv my_list.json > my.csv

	jsonobjects2csv my_list.json my.csv

	cat my_list.json | jsonobjects2csv -i - > my.csv
~~~

jsonobjects2csv 1.2.12


