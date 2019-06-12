
# USAGE

	csv2xlsx [OPTIONS] WORKBOOK_NAME SHEET_NAME

## DESCRIPTION


csv2xlsx will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created.


## OPTIONS

Below are a set of options available.

```
    -d, -delimiter       set delimiter character (input)
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate markdown documentation
    -h, -help            display help
    -i, -input           input filename (CSV content)
    -l, -license         display license
    -o, -output          output filename
    -quiet               suppress error messages
    -sheet               Sheet name to create/replace
    -trim-leading-space  trim leading space in field(s) for CSV input
    -use-lazy-quotes     use lazy quotes for CSV input
    -v, -version         display version
    -workbook            Workbook name
```


## EXAMPLES


Converting a csv to a workbook.

	csv2xlsx -i data.csv MyWorkbook.xlsx 'My worksheet 1'

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

	cat data.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet 2'

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.


csv2xlsx v0.0.25
