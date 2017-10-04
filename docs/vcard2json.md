
# USAGE

## vcard2json [OPTIONS]

## SYNOPSIS

vcard2json converts a VCard to JSON. The vcard can be read from stdin or form a file
with the usual options. The JSON version will be written to stdout.

## OPTIONS

	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version

## EXAMPLES

Simple usage of building a CSV file one rows at a time.

```shell
    cat my.cvf | vcard2json > myVCard.json
```

Or reading, writing to specific file

```shell
    vcard2json -i mv.cvf -o myVCard.json
```


vcard2json v0.0.15
