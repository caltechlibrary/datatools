
# jsoncols

## USAGE

    jsoncols [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]

## SYSNOPSIS

jsoncols provides for both interactive exploration of JSON structures like jid 
and command line scripting flexibility for data extraction into delimited
columns. This is helpful in flattening content extracted from JSON blobs.
The default delimiter for each value extracted is a comma. This can be
overridden with an option.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to 
  explicity read from stdin
	+ if not provided then jsoncols reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to 
  explicity write to stdout
	+ if not provided then jsoncols write to stdout

## OPTIONS

```
	-d	set the delimiter for multi-field output
	-h	display help
	-i	read JSON from a file
	-input	read JSON from a file
	-l	display license
	-m	display output in monochrome
	-r	run interactively
	-repl	run interactively
	-v	display version
```

## EXAMPLES

If myblob.json contained

```
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
```

Getting just the name could be done with

```
    jsoncols .name myblob.json
```

This would yeild

```
    "Doe, Jane"
```

Flipping .name and .age into pipe delimited columns is as 
easy as listing each field in the expression inside a 
space delimited string.

```
    jsoncols -d\|  ".name .age" myblob.json
```

This would yeild

```
    "Doe, Jane"|42
```

