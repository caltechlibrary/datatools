
# Someday, Maybe

+ csv2json would convert CSV content to a 2d-JSON array
    + E.g. `cat file1.csv | csv2json`
+ csv2xlsx would take CSV content piped from stdin and write it do a sheet in an Excel file (creating the Excel workbook if needed)
    + E.g. `cat file1.csv | csv2xlsx MyWorkbook.xlsx "Sheet from file1.csv"`
+ csvfind would accept CSV input from stdin and output rows with matching column values
    + E.g. `cat file1.csv | csvfind -column=3 "Book"`
+ json2csv would convert a 2d JSON array to CSV output, it would comvert a JSON object/map to a column of keys next to a column of values
    + E.g. `cat data.json | json2csv`

