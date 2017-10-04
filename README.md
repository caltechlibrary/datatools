
# datatools

## For data

Command line utilities for simplifying work with CSV, JSON, Excel Workbooks and plain text files or content and
general purpose shell scripting. 

+ [csv2json](docs/csv2json.html) - a tool to take a CSV file and convert it into a JSON blob array or a list of JSON blobs one per line
+ [csv2mdtable](docs/csv2mdtable.html) - a tool to render CSV as a Github Flavored Markdown table
+ [csv2xlsx](docs/csv2xlsx.html) - a tool to take a CSV file and add it as a sheet to a Excel Workbook file.
+ [csvcols](docs/csvcols.html) - a tool for formatting command line arguments into CSV row of columns or filtering CSV rows for specific columns
+ [csvfind](docs/csvfind.html) - a tool for filtering a CSV file by column's value 
+ [csvjoin](docs/csvjoin.html) - a tool to join to CSV files on common values in designated columns, writes combined CSV rows to stdout
+ [csvrows](docs/csvrows.html) - a tool for formatting command line arguments into CSV columns of rows or filtering CSV columns for specific rows
+ [jsoncols](docs/jsoncols.html) - a tool for exploring and extracting JSON values into columns
+ [jsonjoin](docs/jsonjoin.html) - a tool for joining JSON object documents
+ [jsonmunge](docs/jsonmunge.html) - a tool to transform JSON documents into something else
+ [jsonrange](docs/jsonrange.html) - a tool for iterating for JSON maps and arrays
+ [vcard2json](docs/vcard2json.html) - an experimental tool to convert vCards to JSON
+ [xlsx2csv](docs/xlsx2csv.html) - a tool for converting Excel Workbooks sheets to a CSV file(s)
+ [xlsx2json](docs/xlsx2json.html) - a tool for converting Excel Workbooks to JSON files


Compiled versions are provided for Linux (amd64), Mac OS X (amd64),
Windows 10 (amd64) and Raspbian (ARM7). See https://github.com/caltechlibrary/datatools/releases.

Use the utilities try "-help" option for a full list of options.


## For scripting

Various utilities for simplifying work on the command line. 

+ [findfile](docs/findfile.html) - find files based on prefix, suffix or contained string
+ [finddir](docs/finddir.html) - find directories based on prefix, suffix or contained string
+ [mergepath](docs/mergepath.html) - prefix, append, clip path variables
+ [range](docs/range.html) - emit a range of integers (useful for numbered loops in Bash)
+ [reldate](docs/reldate.html) - display a relative date in YYYY-MM-DD format
+ [timefmt](docs/timefmt.html) - format a time value based on Golang's time format language
+ [urlparse](docs/urlparse.html) - split a URL into parts

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


