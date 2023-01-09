---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [YAML_FILENAME] [JSON_NAME]

# DESCRIPTION

{app_name} is a tool that converts YAML into JSON output.

# OPTIONS


-help:
display help

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
    {app_name} my.yaml > my.json

    {app_name} my.yaml my.json

	cat my.yaml | {app_name} -i - > my.json
~~~

{app_name} {version}

