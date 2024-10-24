%string(1) user manual | version 1.2.11 ff4493f
% R. S. Doiel
% 2024-10-24

# NAME

string

# SYNOPSIS

string [OPTIONS] [VERB] [VERB PARAMETERS...]

# DESCRIPTION

string is a command line tool for transforming strings in common ways.

- string length
- changing cases
- checking for prefixes, suffixes 
- trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
- position, counting and replacing substrings
- splitting a string into a JSON array of strings, joining JSON a string arrays into a string

VERB refers to the operation to performed on the supplied string(s).
VER PARAMETERS are thsose additional terms need to complete the process
provided by the VERB.

# OPTIONS

Options always come before the VERB.

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
: has substrings: SUBSTRING [STRING] `string contains SUBSTRING [STRING]`

count
: count substrings: SUBSTRING [STRING] `string count SUBSTRING [STRING]`

englishtitle
: English style title case: [STRING] `string englishtitle [STRING]`

hasprefix
: true/false on prefix: PREFIX [STRING] `string hasprefix PREFIX [STRING]`

hassuffix
: true/false on suffix: SUFFIX [STRING] `string hassuffix SUFFIX [STRING]`

join
: join JSON array into string: DELIMITER [JSON_ARRAY] `string join DELIMITER [JSON_ARRAY]`

length
: length of string: [STRING] `string length [STRING]`

padleft
: left pad PADDING MAX_LENGTH [STRING] `string padleft PADDING MAX_LENGTH [STRING]`

padright
: right pad PADDING MAX_LENGTH [STRING] `string padright PADDING MAX_LENGTH [STRING]`

position
: position of substring: SUBSTRING [STRING] `string position SUBSTRING [STRING]`

replace
: replace: OLD NEW [STRING] `string replace OLD NEW [STRING]`

replacen
: replace n times: OLD NEW N [STRING] `string replacen OLD NEW N [STRING]`

slice
: copy a substring: START END [STRING] `string slice START END [STRING]`

split
: split into a JSON array: DELIMITER [STRING] `string split DELIMITER [STRING]`

splitn
: split into an N length JSON array: DELIMITER N [STRING] `string splitn DELIMITER N [STRING]`

tolower
: to lower case: [STRING] `string tolower [STRING]`

totitle
: to title case: [STRING] `string totitle [STRING]`

toupper
: to upper case: [STRING] `string toupper [STRING]`

trim
: trim (beginning and end), CUTSET [STRING] `string trim CURSET [STRING]`

trimleft
: left trim CUTSET [STRING] `string trimleft CUTSET [STRING]`

trimprefix
: trims prefix: PREFIX [STRING] `string trimprefix PREFIX [STRING]`

trimright
: right trim: CUTSET [STRING] `string trimright CUTSET [STRING]`

trimspace
: trim leading and trailing spaces: [STRING] `string trimspace [STRING]`

trimsuffix
: trim suffix: SUFFIX [STRING] `string trimsuffix SUFFIX [STRING]`

# EXAMPLES

Convert text to upper case

~~~
	string toupper "one"
~~~

Convert text to lower case

~~~
	string tolower "ONE"
~~~

Captialize an English phrase

~~~
	string englishtitle "one more thing to know"
~~~

Split a space newline delimited list of words into a JSON array

~~~
	string -i wordlist.txt split "\n"
~~~

Join a JSON array of strings into a newline delimited list

~~~
	string join '\n' '["one","two","three","four","five"]'
~~~

string 1.2.11

