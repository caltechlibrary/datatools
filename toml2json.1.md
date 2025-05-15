%toml2json(1) user manual | version 1.3.4 4312aaa
% R. S. Doiel
% 2025-05-15

# NAME
 
toml2json

# SYNOPSIS

toml2json [OPTIONS] [TOML_FILENAME] [JSON_NAME]

# DESCRIPTION

toml2json is a tool that converts TOML into JSON. It operates
on standard input and writes to standard output.

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

-p, -pretty
: pretty print output

-quiet
: suppress error messages


# EXAMPLES

These would get the file named "my.toml" and save it as my.json

~~~
    toml2json my.toml > my.json

    toml2json my.toml my.json

	cat my.toml | toml2json -i - > my.json
~~~

toml2json 1.3.4

