
# Using csv2json

Convert data1.csv to data1.json using command line options and Unix pipes. The result will be an array of JSON objects.

```shell
    csv2json -i data1.csv -o data1.json
    cat data1.csv | csv2json > data1.json
```

Convert data1.csv to JSON blobs, one line per blob.

```shell
    csv2json -as-blobs -i data1.csv
    cat data1.csv | csv2json -as-blobs
```

## example files

- [data1.csv](data1.csv)
- [blobs.txt](blobs.txt)
- [csv2json-demo.bash](csv2json-demo.bash)


