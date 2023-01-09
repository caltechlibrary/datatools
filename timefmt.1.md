---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] TIME_STRING_TO_CONVERT

# DESCRIPTION

{app_name} formats the current date or INPUT_DATE based on the output
format provided in options. The default input and  output format is
RFC3339.  Formats are specified based on Golang's time package including
the common constants (e.g. RFC822, RFC1123). 

For details see https://golang.org/pkg/time/#Time.Format.

One additional time layout provided by {app_name} 
 
- mysql "2006-01-02 15:04:05 -0700"

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-if, -input-format
: Set format for input

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-of, -output-format
: Set format for output

-quiet
: suppress error messages

-utc
: timestamps in UTC


# EXAMPLES

Format the date July, 7, 2016 in YYYY-MM-DD format

~~~
    {app_name} -if "2006-01-02" -of "01/02/2006" "2017-12-02"
~~~

Yields "12/02/2017"

Format the MySQL date/time of 8:08am, July 2, 2016

~~~
    {app_name} -input-format mysql -output-format RFC822  "2017-12-02 08:08:08"
~~~

Yields "02 Dec 17 08:08 UTC"

{app_name} {version}
