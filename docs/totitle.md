
# USAGE

## totitle [OPTIONS] [STRINGS]

## SYNOPSIS

totitle turns a UTF-8 string(s) into title case (in English the same as upper case).
If the option -c or -capitalize is used it'll use a naive approach to
captialization rather than title case. If the option -ce or -capitalize-english
added rules English specific capitalization rules will be used.

## OPTIONS

```
	-c	capitalize words
	-capitalize	capitalize words
	-capitalize-english	english capitalization of words
	-ce	english capitalization of words
	-example	display example(s)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-newline	output a newline
	-nl	output a newline
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLE

Title case the string "the cat in the hat"

```
    totitle "the cat in the hat"
```

This should yield

```
    THE CAT IN THE HAT
```

Usage -c or -capitalize option "the cat in the hat"

```
    totitle -c "the cat in the hat"
```

should yeild

```
    "The Cat In The Hat"
```

Using -ce or -capitalize-english option "the cat in the hat"

```
    totitle -ce "the cat in the hat"
```

should yeild

```
    "The Cat in the Hat"
```

totitle v0.0.18
