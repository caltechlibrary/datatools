
# USAGE

    fields [OPTIONS]

## SYNOPSIS

fields reads a line of text from stdin and writing fields as JSON array, CSV row or delimited text to stdout. 
Additional options include ignoring punctation, changing case or allowing special characters. The standard
delimited output is a space.

## OPTIONS

```
	-allow-characters	also allow these characters
	-allow-punctuation	allow punctuation (i.e. allows letters, numbers and punctuation)
	-csv	output as a CSV row
	-delimiter	use this delimiter for output and stop words (default is space)
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-json	output as a JSON array
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-stop-words	a colon delimited list of stop words to ignore (case insensitive)
	-to-lower	lower case the input string
	-to-upper	upper case the input string
	-v	display version
	-version	display version
```

## EXAMPLES

Convert sentence into a JSON array of words.

```shell
    echo "The cat jumpted over the shifty fox." | fields -json
```

Convert each word into a column in a CSV row.

```shell
    echo "The cat jumpted over the shifty fox." | fields -csv
```

