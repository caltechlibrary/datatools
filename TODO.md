
Action Items
------------

Bug
---

+ [ ] findfile v0.0.23-pre option -f, -full-path doesn't return full paths
+ [ ] [Issue #12](https://github.com/caltechlibrary/datatools/issues/12)
    - do we support non-string representation of context?
    - do we support string representation in person/organisation?
    - is this added complexity worth it?

Next
----

+ [ ] Create man pages for all cli, adopt man page structure for usage.
+ [ ] Possible needed tooling
    - [x] A codemeta generator (e.g. read a CITATION.cff and write a codemeta file)
+ [ ] Drop cli package, update Makefile, switch from mk_website.py to website.mak, add Man pages
+ [ ] Review https://csvkit.readthedocs.io/en/latest/ and implement the features in datatools that make sense, e.g. csvsql, a csv2sql, sql2csv, etc.
+ [ ] upgrade to use the new cli v0.0.5-dev
+ [ ] csvrows would output a range of rows (e.g. [2:] would be all rows but the first row)
+ [ ] csv utilities should support integer ranges notation for columns and rows references, E.g. "1,3:4,7,10:" or all

Someday, Maybe
--------------

+ [ ] finddir should have an option to exclude directories (e.g. exclude .git directories from a listing)
+ [ ] textscraper - a tool for select out text and storing it as a JSON field value, sort grep plus sed cleanup and semi-structured text (e.g. webpage)
    + look at how cut, sed, grep are commonly used in my scripts and merge that functionality into a single tool
+ [ ] csvcols, csvrows should have a length option to give you a number of columns or rows respectively
+ [ ] csvcols, csvrows should have a filter option to filter to support filting output conditionally
+ [ ] csvsort should allow a multi-column sort respecting column headings
    + plus column number would be ascending by that column
    + minos column number would be descending by that column
    + sort would be read from left to right
    + it would be good to include support for column names and not just column numbers to describe the sort
+ [ ] jsonmodify takes a JSON document, a dotpath and value then creates/updates the dotpath in the JSON document with the new value
    + "(delete DOTPATH)" would remove the property described by the dotpath
    + "(update DOTPATH NEW_VALUE)" would replace the property described by the dotpath with a new value (value can be a string, number, or JSON)
    + "(create" DOTPATH NEW_VALUE)" would add a new property at the described dotpath with a new value (value can be a string, number, or JSON)
    + "(join DOTH_PATH SEP)" combines JSON array elements into a string version using separator
    + "(concat DOTPATH1 DOTPATH2... SEP)" combines values into a concatenated string, it takes one or more dotpath values (must be string or number) and return them as a concatenated value (concat .last_name .first_name ", ") would return a last name comma first name string.
    + "(split DOTH_PATH SEP)" turns a string into an array of strings using separator
+ [ ] csvcols, csvrows should have a filter mechanism should provide a mechanism to filter by column or row
    + using a prefix notation (e.g. '(and (eq (join (cols (colNo "Last Name") (colNo "First Name")) ", ") "Doiel, R. S.") (gt (cols 4) "2017-06-12"))')
+ [ ] csvfind, csvjoin should have an inverted match operation
+ [ ] a range should accept the word "all" as well as comma delimited list of rows and ranges
+ [ ] Add -uuid and -skip-header-row options constistantly to all csv tools
    + [ ] csvcols
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
+ [ ] smartcat would function like cat but with support for ranges of lines (e.g. show me last 20 lines: smartcat -start=0 -end="-20" file.txt; cat starting with 10th line: smartcat -start=10 file.txt)
    + [ ] allow prefix line number with a specific delimiter (E.g. comma would let you cat a CSV file adding row numbers as first column)
    + [ ] show lines with prefix, suffix, containing or regxp
    + [ ] show lines without prefix, suffix, containing or regexp

Completed
---------

+ [x] consolidate string utilities (e.g. toupper, tolower, totitle) into string cli
+ [x] csvcols -col option should not be a boolean, it should take a range like other csv cli
+ [x] utilities should use starting index of 1 instead of zero as humans refer to column 1 when intending to work on the first column
+ [x] for all cli the -delimiter option should support special characters like \t, \n, \r
+ [x] csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -levenshtein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -inverted -levenstein -stop-words="the:a:of" -col=1 "This Red Book of West March"`
    + E.g. `cat file1.csv | csvfind -contains -col=1 "Red Book"`
+ [x] csvjoin should have option for fuzzy match on columns (e.g. comparing titles)
