---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [VERB] [VERB PARAMETERS...]

# DESCRIPTION

{app_name} is a command line tool for transforming strings in common ways.

- string length
- changing cases
- checking for prefixes, suffixes 
- trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
- position, counting and replacing substrings
- splitting a string into a JSON array of strings, joining JSON a string arrays into a string

# OPTIONS

-help
: display help

-license
:display license

-version
: display version

-d, -delimiter
: set the delimiter

-do, -output-delimiter
: set the output delimiter

-i, -input
: input file name

-nl, -newline
: if true add a trailing newline

-o, -output
: output file name

-quiet
: suppress error messages


## VERBS

contains
: has substrings: SUBSTRING [STRING] `{app_name} contains SUBSTRING [STRING]`

count
: count substrings: SUBSTRING [STRING] `{app_name} count SUBSTRING [STRING]`

englishtitle
: English style title case: [STRING] `{app_name} englishtitle [STRING]`

hasprefix
: true/false on prefix: PREFIX [STRING] `{app_name} hasprefix PREFIX [STRING]`

hassuffix
: true/false on suffix: SUFFIX [STRING] `{app_name} hassuffix SUFFIX [STRING]`

join
: join JSON array into string: DELIMITER [JSON_ARRAY] `{app_name} join DELIMITER [JSON_ARRAY]`

length
: length of string: [STRING] `{app_name} length [STRING]`

padleft
: left pad PADDING MAX_LENGTH [STRING] `{app_name} padleft PADDING MAX_LENGTH [STRING]`

padright
: right pad PADDING MAX_LENGTH [STRING] `{app_name} padright PADDING MAX_LENGTH [STRING]`

position
: position of substring: SUBSTRING [STRING] `{app_name} position SUBSTRING [STRING]`

replace
: replace: OLD NEW [STRING] `{app_name} replace OLD NEW [STRING]`

replacen
: replace n times: OLD NEW N [STRING] `{app_name} replacen OLD NEW N [STRING]`

slice
: copy a substring: START END [STRING] `{app_name} slice START END [STRING]`

split
: split into a JSON array: DELIMITER [STRING] `{app_name} split DELIMITER [STRING]`

splitn
: split into an N length JSON array: DELIMITER N [STRING] `{app_name} splitn DELIMITER N [STRING]`

tolower
: to lower case: [STRING] `{app_name} tolower [STRING]`

totitle
: to title case: [STRING] `{app_name} totitle [STRING]`

toupper
: to upper case: [STRING] `{app_name} toupper [STRING]`

trim
: trim (beginning and end), CUTSET [STRING] `{app_name} trim CURSET [STRING]`

trimleft
: left trim CUTSET [STRING] `{app_name} trimleft CUTSET [STRING]`

trimprefix
: trims prefix: PREFIX [STRING] `{app_name} trimprefix PREFIX [STRING]`

trimright
: right trim: CUTSET [STRING] `{app_name} trimright CUTSET [STRING]`

trimspace
: trim leading and trailing spaces: [STRING] `{app_name} trimspace [STRING]`

trimsuffix
: trim suffix: SUFFIX [STRING] `{app_name} trimsuffix SUFFIX [STRING]`

# EXAMPLES

Convert text to upper case

~~~
	{app_name} toupper "one"
~~~

Convert text to lower case

~~~
	{app_name} tolower "ONE"
~~~

Captialize an English phrase

~~~
	{app_name} englishtitle "one more thing to know"
~~~

Split a space newline delimited list of words into a JSON array

~~~
	{app_name} -i wordlist.txt split "\n"
~~~

Join a JSON array of strings into a newline delimited list

~~~
	{app_name} join '\n' '["one","two","three","four","five"]'
~~~

{app_name} {version}
