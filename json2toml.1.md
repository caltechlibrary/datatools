%json2toml(1) user manual | version 1.3.4 4312aaa
% R. S. Doiel
% 2025-05-15

# NAME

json2toml 

# SYNOPSIS

json2toml [OPTIONS] [JSON_FILENAME] [TOML_FILENAME]

# DESCRIPTION

json2toml is a tool that converts JSON objects into TOML output.

# OPTIONS

-help
: display help

-license
: display license

-version:
display version

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -pretty
: pretty print output

-quiet
: suppress error messages


# EXAMPLES

These would get the file named "my.json" and save it as my.toml

~~~
    json2toml my.json > my.toml

	json2toml my.json my.toml

	cat my.json | json2toml -i - > my.toml
~~~

json2toml 1.3.4


