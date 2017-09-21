
# csv2json

## USAGE

    csv2json [OPTIONS]

## SYNOPSIS

csv2json reads CSV from stdin and writes a JSON to stdout. JSON output
can be either an array of JSON blobs or one JSON blob (row as object)
per line.

## OPTIONS

```
	-as-blobs	output as one JSON blob per line
    -d          set delimiter character
    -delimiter  set delimiter character
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-use-header	treat the first row as field names
	-v	display version
	-version	display version
```

## EXAMPLES

Convert data1.csv to data1.json using Unix pipes.

```shell
    cat data1.csv | csv2json > data1.json
```

Convert data1.csv to JSON blobs, one line per blob

```shell
    csv2json -as-blobs -i data1.csv
```

