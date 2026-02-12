%json2jsonl(1) user manual | version 1.3.5 f86e208
% R. S. Doiel
% 2026-02-12

# NAME

json2jsonl

# SYNOPSIS

json2jsonl [OPTIIONS]

# DESCRIPTION

json2jsonl reads a JSON array document rending the results as a JSON lines document.
See https://jsonlines.org.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i, -input
: input filename

-o, -output
: output filename

-quiet
: suppress error output

-as-dataset ATTRIBUTE_NAME
: if ATTRIBUTE_NAME is not empty string, find the top level ATTRIBUTE_NAME and use as key value for 
when generating a dataset load compatible version of the JSON array contents.

# EXAMPLES

Convert data1.json containing an array of objects into data1.jsonl using Unix redirection.

~~~
    json2jsonl < data1.json > data1.jsonl
~~~


