
# datatools

Command line utilities for simplifying work with CSV, JSON, Excel Workbooks and plain text files or content.

+ [csvcols](docs/csvcols.html) - a tool for formatting command line arguments into CSV row of columns or filtering CSV rows for specific columns
+ [csvrows](docs/csvrows.html) - a tool for formatting command line arguments into CSV columns of rows or filtering CSV columns for specific rows
+ [csvfind](docs/csvfind.html) - a tool for filtering a CSV file by column's value 
+ [csvjoin](docs/csvjoin.html) - a tool to join to CSV files on common values in designated columns, writes combined CSV rows to stdout
+ [csv2json](docs/csv2json.html) - a tool to take a CSV file and convert it into a JSON blob array or a list of JSON blobs one per line
+ [csv2mdtable](docs/csv2mdtable.html) - a tool to render CSV as a Github Flavored Markdown table
+ [csv2xlsx](docs/csv2xlsx.html) - a tool to take a CSV file and add it as a sheet to a Excel Workbook file.
+ [jsoncols](docs/jsoncols.html) - a tool for exploring and extracting JSON values into columns
+ [jsonrange](docs/jsonrange.html) - a tool for iterating for JSON maps and arrays
+ [xlsx2json](docs/xlsx2json.html) - a tool for converting Excel Workbooks to JSON files
+ [xlsx2csv](docs/xlsx2csv.html) - a tool for converting Excel Workbooks sheets to a CSV file(s)


Compiled versions are provided for Linux (amd64), Mac OS X (amd64),
Windows 10 (amd64) and Raspbian (ARM7). See https://github.com/caltechlibrary/datatools/releases.

Use the utilities try "-help" option for a full list of options.


## Installation

_datatools_ is go get-able.

```
    go get github.com/caltechlibrary/datatools/...
```

Or see [INSTALL.md](install.html) for details for installing 
compiled versions of the programs.


