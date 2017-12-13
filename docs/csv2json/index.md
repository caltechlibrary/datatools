
# USAGE

	csv2json [OPTIONS]

## SYNOPSIS


csv2json reads CSV from stdin and writes a JSON to stdout. JSON output
can be either an array of JSON blobs or one JSON blob (row as object)
per line.


## OPTIONS

```
    -as-blobs                 output as one JSON blob per line
    -d, -delimiter            set the delimter character
    -examples                 display example(s)
    -generate-markdown-docs   generation markdown documentation
    -h, -help                 display help
    -i, -input                input filename
    -l, -license              display license
    -nl, -newline             include trailing newline in output
    -o, -output               output filename
    -quiet                    suppress error output
    -trim-leading-space       trim leading space in fields for CSV input
    -use-header               treat the first row as field names
    -use-lazy-quotes          use lazy quotes for for CSV input
    -v, -version              display version
```


## EXAMPLES


Convert data1.csv to data1.json using Unix pipes.

    cat data1.csv | csv2json > data1.json

Convert data1.csv to JSON blobs, one line per blob

    csv2json -as-blobs -i data1.csv


csv2json v0.0.23-pre
