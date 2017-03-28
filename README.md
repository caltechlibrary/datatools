
# datatools

Various utilities for simplifying JSON, Excel Workbook and CSV data work on the command line. 

+ [csvcols](csvcols.html) - a tool for formatting command line arguments intoa CSV row or filtering CSV rows for specific columns
+ [csvjoin](csvjoin.html) - a tool to join to CSV files on common values in designated columns, writes combined CSV rows to stdout
+ [csv2xlsx](csv2xlsx.html) - a tool to take a CSV file and add it as a sheet to a Excel Workbook file.
+ [csv2mdtable](csv2mdtable.html) - a tool to render CSV as a Github Flavored Markdown table
+ [jsoncols](jsoncols.html) - a tool for exploring and extracting JSON values into columns
+ [jsonrange](jsonrange.html) - a tool for iterating for JSON maps and arrays
+ [xlsx2json](xlsx2json.html) - a tool for converting Excel Workbooks to JSON files
+ [xlsx2csv](xlsx2csv.html) - a tool for converting Excel Workbooks sheets to a CSV file(s)

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


