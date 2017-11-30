
# USAGE

## string [OPTIONS] [ACTION] [ACTION PARAMETERS...]

## OPTIONS

```
    -e, -examples  display examples
    -h, -help      display help
    -i, -input     input file name
    -l, -license   display license
    -nl, -newline  output a trailing newline
    -o, -output    output file name
    -quiet         suppress error messages
    -t, -text      handle arrays as plain text
    -v, -version   display version
```

## ACTIONS

```
    hasprefix   output true if string(s) have prefix otherwise false, first parameter is prefix
    hassuffix   output true if string(s) have suffix otherwise false, first parameter is suffix
    join        join JSON array(s) of strings or join delimited input, first parameter is delimiter
    lefttrim    left trim the cutset from a string(s), first parameter is cutset
    lower       converts a string(s) to lower case
    split       splits a string into a JSON array or delimiter output, first parameter is delimiter
    title       converts a string(s) to title case
    trim        trims the cutset from beginning and end of string(s), first parameter is cutset
    trimprefix  trims the prefix from a string(s), first parameter is prefix
    trimsuffix  trims the suffix from a string(s), first parameter is suffix
    upper       converts a string(s) to upper case
``

string v0.0.19-dev
