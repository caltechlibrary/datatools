%json2yaml(1) user manual | version 1.2.10 9821b18
% R. S. Doiel
% 2024-08-26

# NAME

json2yaml

# SYNOPSIS

json2yaml [OPTIONS] [JSON_FILENAME] [YAML_FILENAME]

# DESCRIPTION

json2yaml is a tool that converts JSON objects into YAML output.

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


# EXAMPLES

These would get the file named "my.json" and save it as my.yaml

~~~
    json2yaml my.json > my.yaml

	json2yaml my.json my.taml

	cat my.json | json2yaml -i - > my.taml
~~~

json2yaml 1.2.10


