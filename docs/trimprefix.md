
# USAGE

## trimprefix [OPTIONS] PREFIX_STRING [STRINGS_TO_CHECK] 

## SYNOPSIS

trimprefix returns a version of the string(s) without the designated prefix string.

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

Trim the prefix "un" from the words "unknowning" and "unusual"

```
    trimprefix "un" "unknowning" "unusual"
```

Should yield

```
    knowning
    usual
```

Trim the prefix "un" from the words "unknown" and "common"

```
    trimprefix "un" "unknowning" "common"
```

Should yield

```
    knowning
    common
```

trimprefix v0.0.18
