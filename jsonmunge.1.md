%jsonmunge(1) user manual | version 1.3.0 f486d87
% R. S. Doiel
% 2025-01-31

# NAME

jsonmunge 

# SYNOPSIS

jsonmunge [OPTIONS] TEMPLATE_FILENAME

# DESCRIPTION

jsonmunge is a command line tool that takes a JSON document and
one or more Go templates rendering the results. Useful for
reshaping a JSON document, transforming into a new format,
or filter for specific content.

- TEMPLATE_FILENAME is the name of a Go text tempate file used to render
  the outbound JSON document

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-E, -expression
: use template expression as template

-i, -input
: input filename

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

If person.json contained

~~~
   {"name": "Doe, Jane", "email":"jd@example.org", "age": 42}
~~~

and the template, name.tmpl, contained

~~~
   {{- .name -}}
~~~

Getting just the name could be done with

~~~
    cat person.json | jsonmunge name.tmpl
~~~

This would yield

~~~
    "Doe, Jane"
~~~

jsonmunge 1.3.0

