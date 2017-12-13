
# USAGE

	csv2xlsx [OPTIONS] WORKBOOK_NAME SHEET_NAME

## SYNOPSIS


csv2xlsx will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created.


## OPTIONS

```
    -d, -delimiter            set delimiter character (input)
    -examples                 display example(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -i, -input                input filename (CSV content)
    -l, -license              display license
    -o, -output               output filename
    -quiet                    suppress error messages
    -sheet                    Sheet name to create/replace
    -v, -version              display version
    -workbook                 Workbook name
```


## EXAMPLES


Converting a csv to a workbook.

	csv2xlsx -i data.csv MyWorkbook.xlsx 'My worksheet 1'

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

	cat data.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet 2'

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.
%!(EXTRA string=csv2xlsx, string=csv2xlsx)

csv2xlsx v0.0.22-pre
