
# csv2mdtable

## USAGE

    csv2mdtable [OPTIONS]

## SYNOPSIS

csv2mdtable reads CSV from stdin and writes a Github Flavored Markdown
table to stdout. 

## OPTIONS

```
    -d          set delimiter character
    -delimiter  set delimiter character
	-h	display help
	-help	display help
	-i	set input file
	-l	display license
	-license	display license
	-o	set output file
	-v	display version
	-version	display version
```

## EXAMPLES

Convert data1.csv to data1.md using Unix pipes.

```
    cat data1.csv | csv2mdtable > data1.md
```

Convert data1.csv to data1.md using options.

```
    csv2mdtable -i data1.csv -o data1.md
```

