
# Action Items

## Next

+ [ ] utilities should use starting index of 1 instead of zero as humans refer to column 1 when intending to work on the first column
+ [ ] csvfind, csvjoin should have an inverted match operation
+ [ ] csv utilities to support integer ranges notation for columns and rows references, E.g. "1,3-4,7" or all

## Someday, Maybe

+ [ ] unify the options vocabulary to work the same between each cli
    + Need a common approach to column ranges in csvcols, csvfind, csvjoin
    + csv2json, csv2mdtable, csv2xlsx should accept a column and row range option for output
+ [ ] csvfind add filter by row number (helpful when combined with csvcols for snapshotting the middle of a table)
+ [ ] csv2json should have an option that will include a row number in JSON blob output
+ [ ] csv2json should have the options to normalize property names in JSON objects
    + camel case
    + snake case
    + lower case/upper case
    + space to underscores 
    + strip punctuation
    + rename keys
+ [ ] csvrotate would take a CSV file as import and output columns as rows
+ [ ] json2csv would convert a 2d JSON array to CSV output, it would comvert a JSON object/map to a column of keys next to a column of values
    + E.g. `cat data.json | json2csv`

## Completed

+ [x] for all cli the -delimiter option should support special characters like \t, \n, \r
+ [x] csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -levenshtein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -inverted -levenstein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -contains -col=1 "Red Book"`
+ [x] csvjoin should have option for fuzzy match on columns (e.g. comparing titles)
