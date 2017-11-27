
# USAGE

## hasprefix [OPTIONS] PREFIX_STRING [STRINGS_TO_CHECK] 

## SYNOPSIS

hasprefix returns a string "true" or "false" based on whether or not the provided
string(s) contains the prefix. Also hasprefix will exit with a status code 0 prefix
is found in all string or status code 1 if one string does not have prefix.
Can read string to check from standard in with an option
of "-i -" or a file with "-i FILENAME".

## OPTIONS

```
	-example	display example(s)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-new-line	include a trailing newline in output
	-nl	include a trailing newline in output
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLES

See if the words "unknown" and "unusual" start with the prefix "un"

```shell
    hasprefix "un" "unknown" "unusual"
```

Should yield

```
    true
    true
```

and exit with a status code of zero.

See if the words "unknown" and "common" start with the prefix "un"

```shell
    hasprefix "un" "unknown" "common"
```

Should yield

```
    true
    false
```

and exist with a status code of one.

hasprefix v0.0.18
