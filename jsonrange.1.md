---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS] [DOT_PATH_EXPRESSION]

# DESCRIPTION

{app_name} returns returns a range of values based on the JSON structure
being read and options applied.  Without options the JSON structure is
read from standard input and writes a list of keys to standard out. Keys
are either attribute names or for arrays the index position (counting
form zero).  If a DOT_PATH_EXPRESSION is included on the command line then
that is used to generate the results. Using options to can choose to read
the JSON data structure from a file, write the output to a file as well
as display values instead of keys. a list of "keys" of an index or map in
JSON.  

Using options it can also return a list of values.  The JSON object is
read from standard in and the resulting list is normally written to
standard out. There are options to read or write to files.  Additional
parameters are assumed to be a dot path notation select the parts of the
JSON data structure you want from the range. 

DOT_PATH_EXPRESSION is a dot path stale expression indicating what you
want range over.  E.g.

- . would indicate the whole JSON data structure read is used to range over
- .name would indicate to range over the value pointed at by the "name" attribute 
- ["name"] would indicate to range over the value pointed at by the "name" attribute
- [0] would indicate to range over the value held in the zero-th element of the array

The path can be chained together

- .name.family would point to the value heald by the "name" attributes' "family" attribute.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-d, -delimiter
: set delimiter for range output

-i, -input
: read JSON from file

-last
: return the index of the last element in list (e.g. length - 1)

-length
: return the number of keys or values

-limit
: limit the number of items output

-nl, -newline
: if true add a trailing newline

-o, -output
: write to output file

-quiet
: suppress error messages

-values
: return the values instead of the keys


# EXAMPLES

Working with a map

~~~
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
       | {app_name}
~~~

This would yield

~~~
    name
    email
    age
~~~

Using the -values option on a map

~~~
    echo '{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}' \
      | {app_name} -values
~~~

This would yield

~~~
    "Doe, Jane"
    "jane.doe@example.org"
    42
~~~


Working with an array

~~~
    echo '["one", 2, {"label":"three","value":3}]' | {app_name}
~~~

would yield

~~~
    0
    1
    2
~~~

Using the -values option on the same array

~~~
    echo '["one", 2, {"label":"three","value":3}]' | {app_name} -values
~~~

would yield

~~~
    one
    2
    {"label":"three","value":3}
~~~

Checking the length of a map or array or number of keys in map

~~~
    echo '["one","two","three"]' | {app_name} -length
~~~

would yield

~~~
    3
~~~

Check for the index of last element

~~~
    echo '["one","two","three"]' | {app_name} -last
~~~

would yield

~~~
    2
~~~

Check for the index value of last element

~~~
    echo '["one","two","three"]' | {app_name} -values -last
~~~

would yield

~~~
    "three"
~~~

Limitting the number of items returned

~~~
    echo '[10,20,30,40,50]' | %!s(MISSING) -limit 2
~~~

would yield

~~~
    1
    2
~~~

Limitting the number of values returned

~~~
    echo '[10,20,30,40,50]' | %!s(MISSING) -values -limit 2
~~~

would yield

~~~
    10
    20
~~~

{app_name} {version}
