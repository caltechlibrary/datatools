---
title: "yaml2json (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

yaml2json

# SYNOPSIS

yaml2json [OPTIONS] [YAML_FILENAME] [JSON_NAME]

# DESCRIPTION

yaml2json is a tool that converts YAML into JSON. The
program reads from standard input and writes to standard out.

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

These would get the file named "my.yaml" and save it as my.json

~~~
    yaml2json my.yaml > my.json

    yaml2json my.yaml my.json

	cat my.yaml | yaml2json -i - > my.json
~~~

yaml2json 1.2.2
