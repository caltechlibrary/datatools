
# USAGE

	string [OPTIONS] [VERB] [VERB PARAMETERS...]

## DESCRIPTION


string is a command line tool for transforming strings in common ways.

+ string length
+ changing cases
+ checking for prefixes, suffixes 
+ trimming prefixes, suffixes and cutsets (i.e. list of characters to cut)
+ position, counting and replacing substrings
+ splitting a string into a JSON array of strings, joining JSON a string arrays into a string


## OPTIONS

Below are a set of options available.

```
    -d, -delimiter           set the delimiter
    -do, -output-delimiter   set the output delimiter
    -e, -examples            display examples
    -generate-manpage        generate man page
    -generate-markdown       generate Markdown documentation
    -h, -help                display help
    -i, -input               input file name
    -l, -license             display license
    -nl, -newline            if true add a trailing newline
    -o, -output              output file name
    -quiet                   suppress error messages
    -v, -version             display version
```


## EXAMPLES


Convert text to upper case

	string toupper "one"

Convert text to lower case

	string tolower "ONE"

Captialize an English phrase

	string englishtitle "one more thing to know"

Split a space newline delimited list of words into a JSON array

	string -i wordlist.txt split "\n"

Join a JSON array of strings into a newline delimited list

	string join '\n' '["one","two","three","four","five"]'



string v0.0.25
