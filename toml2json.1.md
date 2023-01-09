---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME
 
{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [TOML_FILENAME] [JSON_NAME]

# DESCRIPTION

{app_name} is a tool that converts TOML into JSON. It operates
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
    {app_name} my.toml > my.json

    {app_name} my.toml my.json

	cat my.toml | {app_name} -i - > my.json
~~~

{app_name} {version}
