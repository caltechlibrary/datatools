
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
    -nl, -newline             output a trailing newline
    -o, -output               output file name
    -quiet                    suppress error messages
    -t, -text                 handle arrays as plain text
    -v, -version              display version
```


## ACTIONS

```
    contains       has substrings: SUBSTRING [STRINGS]
    count          count substrings: SUBSTRING [STRINGS]
    englishtitle   English style title case: [STRINGS]
    hasprefix      true/false on prefix: PREFIX [STRINGS]
    hassuffix      true/false on suffix: SUFFIX [STRINGS]
    join           join JSON array into string: DELIMITER [STRINS]
    length         length of string: [STRINGS]
    pad            pad (beginning and end): PADDING MAX_LENGTH [STRINGS]
    padleft        left pad: PADDING MAX_LENGTH [STRINGS]
    position       position of substring: SUBSTRING [STRINGS]
    replace        replace: TARGET REPLACEMENT [STRINGS]
    replacen       replace n times: TARGET REPLACEMENT COUNT [STRINGS]
    split          split into a JSON array: DELIMITER [STRINGS]
    splitn         split into an N length JSON array: DELIMITER N [STRINGS]
    tolower        to lower case: [STRINGS]
    totitle        to title case: [STRINGS]
    toupper        to upper case: [STRINGS]
    trim           trim (beginning and end), CUTSET [STRINGS]
    trimleft       left trim: CUTSET [STRINGS]
    trimprefix     trims prefix: PREFIX [STRINGS]
    trimright      right trim: CUTSET [STRINGS]
    trimsuffix     trim suffix: SUFFIX [STRINGS]
```


## EXAMPLES




Related: [contains](contains.html), [count](count.html), [englishtitle](englishtitle.html), [hasprefix](hasprefix.html), [hassuffix](hassuffix.html), [join](join.html), [length](length.html), [pad](pad.html), [padleft](padleft.html), [position](position.html), [replace](replace.html), [replacen](replacen.html), [split](split.html), [splitn](splitn.html), [tolower](tolower.html), [totitle](totitle.html), [toupper](toupper.html), [trim](trim.html), [trimleft](trimleft.html), [trimprefix](trimprefix.html), [trimright](trimright.html), [trimsuffix](trimsuffix.html)

string v0.0.19-dev
