
# USAGE

	string [OPTIONS] [ACTION] [ACTION PARAMETERS...]

## SYNOPSIS


string is a command line tool for transforming strings in common ways.

+ changing cases
+ checking for prefixes, suffixes or substrings
+ trimming prefixes, suffixes or specific characters (i.e. cutsets)
+ locating, counting and replacing substrings
+ splitting and joining JSON string arrays
+ formatting and padding strings and numbers


## OPTIONS

Options are shared between all actions and must precede the action on the command line.

```
    -e, -examples             display examples
    -generate-markdown-docs   output documentation in Markdown
    -h, -help                 display help
    -i, -input                input file name
    -l, -license              display license
    -nl, -newline             output a trailing newline
    -o, -output               output file name
    -quiet                    suppress error messages
    -t, -text                 handle arrays as plain text
    -v, -version              display version
```


## ACTIONS

```
    hasprefix    output true if string(s) have prefix otherwise false, first parameter is prefix
    hassuffix    output true if string(s) have suffix otherwise false, first parameter is suffix
    join         join JSON array(s) of strings or join delimited input, first parameter is delimiter
    lower        converts a string(s) to lower case
    split        splits a string into a JSON array or delimiter output, first parameter is delimiter
    title        converts a string(s) to title case
    trim         trims the cutset from beginning and end of string(s), first parameter is cutset
    trimleft     left trim the cutset from a string(s), first parameter is cutset
    trimprefix   trims the prefix from a string(s), first parameter is prefix
    trimright    right trim the cutset from a string(s), first parameter is cutset
    trimsuffix   trims the suffix from a string(s), first parameter is suffix
    upper        converts a string(s) to upper case
```


## EXAMPLES




Related: [hasprefix](hasprefix.html), [hassuffix](hassuffix.html), [join](join.html), [lower](lower.html), [split](split.html), [title](title.html), [trim](trim.html), [trimleft](trimleft.html), [trimprefix](trimprefix.html), [trimright](trimright.html), [trimsuffix](trimsuffix.html), [upper](upper.html)

string v0.0.19-dev
