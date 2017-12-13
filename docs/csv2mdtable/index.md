
# USAGE

	csv2mdtable [OPTIONS]

## SYNOPSIS


csv2mdtable reads CSV from stdin and writes a Github Flavored Markdown
table to stdout.


## OPTIONS

```
    -d, -delimiter            set delimiter character
    -examples                 display example(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -i, -input                input filename
    -l, -license              display license
    -nl, -newline             if true include leading/trailing newline
    -o, -output               output filename
    -quiet                    suppress error message
    -v, -version              display version
```


## EXAMPLES


Convert data1.csv to data1.md using Unix pipes.

    cat data1.csv | csv2mdtable > data1.md

Convert data1.csv to data1.md using options.

    csv2mdtable -i data1.csv -o data1.md


csv2mdtable v0.0.22-pre
