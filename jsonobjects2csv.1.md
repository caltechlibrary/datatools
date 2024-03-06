%jsonobjects2csv(1) user manual | version 1.2.9 0d7364a
% R. S. Doiel
% 2024-03-06

# NAME

jsonobjects2csv

# SYNOPSIS

jsonobjects2csv [OPTIONS] [JSON_FILENAME] [YAML_FILENAME]

# DESCRIPTION

jsonobjects2csv is a tool that converts a JSON list of objects into CSV output.

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

These would get the file named "my_list.json" and save it as my.csv

~~~
    jsonobjects2csv my_list.json > my.csv

	jsonobjects2csv my_list.json my.csv

	cat my_list.json | jsonobjects2csv -i - > my.csv
~~~

jsonobjects2csv 1.2.9


