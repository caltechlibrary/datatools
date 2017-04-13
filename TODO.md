
# Action Items

## Next

+ [ ] csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -levenshtein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -inverted -levenstein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -contains -col=1 "Red Book"`
+ [ ] csvjoin should have option for fuzzy match on columns (e.g. comparing titles)

## Someday, Maybe

+ [ ] utilities should use starting index of 1 instead of zero as humans refer to column 1 as first column
+ [ ] csv2json should have an option that will include a row number in JSON blob output
+ [ ] text2terms should have a mininum term length options (e.g. ignore words that are shorter than three letters)
+ [ ] text2terms should have a max term count (e.g. only return up to 12 words)
+ [ ] csv2json should have the options to normalize property names in JSON objects
    + camel case
    + snake case
    + lower case/upper case
    + space to underscores 
    + strip punctuation
    + rename keys
+ [ ] json2csv would convert a 2d JSON array to CSV output, it would comvert a JSON object/map to a column of keys next to a column of values
    + E.g. `cat data.json | json2csv`

## Completed

+ [x] for all cli the -delimiter option should support special characters like \t, \n, \r
