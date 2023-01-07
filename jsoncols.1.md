---
title: "jsoncols (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-06
---

# NAME

jsoncols

# SYNOPSIS

jsoncols [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

jsoncols provides scripting flexibility for data extraction from JSON data
returning the results in columns.  This is helpful in flattening content
extracted from JSON blobs.  The default delimiter for each value
extracted is a comma. This can be overridden with an option.

- EXPRESSION can be an empty string or dot notation for an object's path
- INPUT_FILENAME is the filename to read or a dash "-" if you want to
  explicitly read from stdin
	- if not provided then jsoncols reads from stdin
- OUTPUT_FILENAME is the filename to write or a dash "-" if you want to
  explicitly write to stdout
	- if not provided then jsoncols write to stdout

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-csv
: output as CSV or other flat delimiter row

-d, -delimiter
: set the delimiter for multi-field csv output

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-p, -pretty
: pretty print JSON output

-quiet
: suppress error messages

-quote
: quote strings and JSON notation

-r, -repl
: run interactively


# EXAMPLES

If myblob.json contained

~~~
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
~~~

Getting just the name could be done with

~~~
    jsoncols -i myblob.json .name
~~~

This would yield

~~~
    "Doe, Jane"
~~~

Flipping .name and .age into pipe delimited columns is as
easy as listing each field in the expression inside a
space delimited string.

~~~
    jsoncols -i myblob.json -d\|  .name .age
~~~

This would yield

~~~
    Doe, Jane|42
~~~

You can also pipe JSON data in.

~~~
    cat myblob.json | jsoncols .name .email .age
~~~

Would yield

~~~
   "Doe, Jane","jane.doe@xample.org",42
~~~

jsoncols 1.2.1


