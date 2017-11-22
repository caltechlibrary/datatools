
# USAGE

## csv2mdtable [OPTIONS]

## SYNOPSIS

csv2mdtable reads CSV from stdin and writes a Github Flavored Markdown
table to stdout. 

## OPTIONS

	-d	set delimiter character
	-delimiter	set delimiter character
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

Convert data1.csv to data1.md using Unix pipes.

```shell
    cat data1.csv | csv2mdtable > data1.md
```

Convert data1.csv to data1.md using options.

```shell
    csv2mdtable -i data1.csv -o data1.md
```


csv2mdtable v0.0.16
