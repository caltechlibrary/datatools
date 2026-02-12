%jsonjoin(1) user manual | version 1.3.5 effbad2
% R. S. Doiel
% 2026-02-12

# NAME

jsonjoin 

# SYNOPSIS

jsonjoin [OPTIONS] JSON_FILENAME [JSON_FILENAME ...]

# DESCRIPTION

jsonjoin joins one or more JSON objects. By default the
objects are each assigned to an attribute corresponding with their
filenames minus the ".json" extension. If the object is read from
standard input then "_" is used as it's attribute name.

If you use the update or overwrite options you will create a merged
object. The update option keeps the attribute value first encountered
and overwrite takes the last attribute value encountered.

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

-create
: Create a root object placing each joined objects under their own attribute

-update
: update first object with the second object, ignore existing attributes

-overwrite
: update first object with the second object, overwriting existing attributes

# EXAMPLES

This is an example of take "my1.json" and "my2.json"
render "my.json"

~~~
    jsonjoin my1.json my2.json >my.json
~~~

my.json would have two attributes, "my1" and "my2" each
with their complete attributes.

Using the update option you can merge my1.json with any additional attribute
values found in m2.json.

~~~
    jsonjoin -update my1.json my2.json >my.json
~~~

Using the overwrite option you can merge my1.json with my2.json accepted
as replacement values.

~~~
    jsonjoin -overwrite my1.json my2.json >my.json
~~~







