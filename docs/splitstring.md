
# USAGE

## splitstring [OPTIONS] [STRING_TO_SPLIT]

splitstring splits a string based on a delimiting string provided. The default
delimiter is a space. You can specify a delimiting string via
the -d or --delimiter option.  splitstring will split the string provided
as a command line argument but can read split string(s) recieved on
stdin in with the -i or --input option. By default the split
strings are render as a JSON array but with the option -nl or
--newline you can render each split string one per line.

## OPTIONS

```
	-d	set the delimiting string value
	-delimiter	set the delimiting string value
	-example	display example(s)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-newline	output as one substring per line rather than JSON
	-nl	output as one substring per line rather than JSON
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLES

Splitting a string that is double pipe delimited rendering
one sub string per line.

```shell
    splitstring -nl -d '||' "one||two||three"
```

This should yield

```
    one
	two
	three
```

Splitting a string that is double pipe delimited rendering JSON

```shell
    splitstring -d '||' "one||two||three"
```

This should yield

```json
   ["one","two","three"]
```

splitstring v0.0.18
