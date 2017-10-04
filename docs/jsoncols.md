
# USAGE

## jsoncols [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]

## SYSNOPSIS

jsoncols provides scripting flexibility for data extraction from JSON data 
returning the results in columns.  This is helpful in flattening content 
extracted from JSON blobs.  The default delimiter for each value 
extracted is a comma. This can be overridden with an option.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to 
  explicity read from stdin
	+ if not provided then jsoncols reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to 
  explicity write to stdout
	+ if not provided then jsoncols write to stdout

## OPTIONS

	-csv	output as CSV or other flat delimiter row
	-d	set the delimiter for multi-field csv output
	-dimiter	set the delimiter for multi-field csv output
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-m	display output in monochrome
	-o	output filename
	-output	output filename
	-permissive	suppress error messages
	-quiet	suppress error messages
	-quote	if dilimiter is found in column value add quotes for non-CSV output
	-r	run interactively
	-repl	run interactively
	-v	display version
	-version	display version

## EXAMPLES

If myblob.json contained

```json
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
```


Getting just the name could be done with

```shell
    jsoncols -i myblob.json .name
```

This would yeild

```json
    "Doe, Jane"
```

Flipping .name and .age into pipe delimited columns is as 
easy as listing each field in the expression inside a 
space delimited string.

```shell
    jsoncols -i myblob.json -d\|  .name .age 
```

This would yeild

```
    "Doe, Jane"|42
```

You can also pipe JSON data in.

```shell
    cat myblob.json | jsoncols .name .email .age
```

Would yield

```csv
   "Doe, Jane",jane.doe@xample.org,42
```

<<<<<<< HEAD
=======

jsoncols v0.0.14
>>>>>>> 25fa0d856c527e91cd7efb24a6331b26291d07a7
