
# Action Items

## Next

+ csv2json should have an option that will include a row number in JSON blob output

## Someday, Maybe

+ csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -column=3 "Book"`
+ json2csv would convert a 2d JSON array to CSV output, it would comvert a JSON object/map to a column of keys next to a column of values
    + E.g. `cat data.json | json2csv`

