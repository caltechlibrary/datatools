
# jsoncols

## USAGE: 

    jsoncols [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]

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

```
	-d	        set the delimiter for multi-field output
	-h	        display help
	-i	        input filename
	-input	    input filename
	-l	        display license
	-m	        display output in monochrome
	-o	        output filename
	-output	    output filename
	-permissive	suppress error messages
	-r	        run interactively
	-repl	    run interactively
	-v	        display version
```


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

```
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

```
   "Doe, Jane",jane.doe@xample.org,42
```

