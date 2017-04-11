
# Action Items

## Next

+ fields should have a mininum term length options (e.g. ignore words that are shorter than three letters)
+ fields should have a max term count (e.g. only return up to 12 words)
+ csv2json should have an option that will include a row number in JSON blob output


## Someday, Maybe

+ csv2json should have the options to normalize property names in JSON objects
    + camel case
    + snake case
    + lower case/upper case
    + space to underscores 
    + strip punctuation
    + rename keys
+ csvjoin should have option for fuzzy match on columns (e.g. comparing titles)
+ csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -column=3 "Book"`
+ json2csv would convert a 2d JSON array to CSV output, it would comvert a JSON object/map to a column of keys next to a column of values
    + E.g. `cat data.json | json2csv`

