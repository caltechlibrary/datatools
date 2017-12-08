
# USAGE

	string [OPTIONS] [ACTION] [ACTION PARAMETERS...]

## SYNOPSIS


string is a command line tool for transforming strings in common ways.

+ string length
+ changing cases
+ checking for prefixes, suffixes 
+ trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
+ position, counting and replacing substrings
+ splitting a string into a JSON array of strings, joining JSON a string arrays into a string


## OPTIONS

Options are shared between all actions and must precede the action on the command line.

```
    -e, -examples             display examples
    -generate-markdown-docs   output documentation in Markdown
    -h, -help                 display help
    -i, -input                input file name
    -l, -license              display license
    -nl, -newline             if true add a trailing newline
    -o, -output               output file name
    -quiet                    suppress error messages
    -v, -version              display version
```


## ACTIONS

```
    contains       has substrings: SUBSTRING [STRING]
    count          count substrings: SUBSTRING [STRING]
    englishtitle   English style title case: [STRING]
    hasprefix      true/false on prefix: PREFIX [STRING]
    hassuffix      true/false on suffix: SUFFIX [STRING]
    join           join JSON array into string: DELIMITER [JSON_ARRAY]
    length         length of string: [STRING]
    padleft        left pad: PADDING MAX_LENGTH [STRING]
    padright       right pad: PADDING MAX_LENGTH [STRING]
    position       position of substring: SUBSTRING [STRING]
    replace        replace: OLD NEW [STRING]
    replacen       replace n times: OLD NEW N [STRING]
    slice          copy a substring: START END [STRING]
    split          split into a JSON array: DELIMITER [STRING]
    splitn         split into an N length JSON array: DELIMITER N [STRING]
    tolower        to lower case: [STRING]
    totitle        to title case: [STRING]
    toupper        to upper case: [STRING]
    trim           trim (beginning and end), CUTSET [STRING]
    trimleft       left trim: CUTSET [STRING]
    trimprefix     trims prefix: PREFIX [STRING]
    trimright      right trim: CUTSET [STRING]
    trimsuffix     trim suffix: SUFFIX [STRING]
```


## EXAMPLES




Related: [contains](contains.html), [count](count.html), [englishtitle](englishtitle.html), [hasprefix](hasprefix.html), [hassuffix](hassuffix.html), [join](join.html), [length](length.html), [padleft](padleft.html), [padright](padright.html), [position](position.html), [replace](replace.html), [replacen](replacen.html), [slice](slice.html), [split](split.html), [splitn](splitn.html), [tolower](tolower.html), [totitle](totitle.html), [toupper](toupper.html), [trim](trim.html), [trimleft](trimleft.html), [trimprefix](trimprefix.html), [trimright](trimright.html), [trimsuffix](trimsuffix.html)

string v0.0.20-dev
