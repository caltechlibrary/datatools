
# datatools

_datatools_ provides a variety of command line programs for working with 
data in different formats as well as to ease Posix shell scripting 
(e.g. writing scripts that run under Bash). The tools are group as data, 
strings and scripting.

## For data

Command line utilities for simplifying work with CSV, JSON, TOML, YAML, 
Excel Workbooks and plain text files or content.

+ [csv2json](docs/csv2json/) - a tool to take a CSV file and convert it into a JSON array or a list of JSON blobs one per line
+ [csv2mdtable](docs/csv2mdtable/) - a tool to render CSV as a Github Flavored Markdown table
+ [csv2xlsx](docs/csv2xlsx/) - a tool to take a CSV file and add it as a sheet to a Excel Workbook
+ [csvcleaner](docs/csvcleaner/) - normalize a CSV file by column and row including trimming spaces and removing comments
+ [csvcols](docs/csvcols/) - a tool for formatting command line arguments into CSV row of columns or filtering CSV rows for specific columns
+ [csvfind](docs/csvfind/) - a tool for filtering a CSV file rows by column
+ [csvjoin](docs/csvjoin/) - a tool to join two CSV files on common values in designated columns, writes combined CSV rows
+ [csvrows](docs/csvrows/) - a tool for formatting command line arguments into CSV columns of rows or filtering CSV for specific rows
+ [json2toml](docs/json2toml/) - a tool for converting JSON to TOML
+ [json2yaml](docs/json2yaml/) - a tool for converting JSON to YAML
+ [jsoncols](docs/jsoncols/) - a tool for exploring and extracting JSON values into columns
+ [jsonjoin](docs/jsonjoin/) - a tool for joining JSON object documents
+ [jsonmunge](docs/jsonmunge/) - a tool to transform JSON documents into something else
+ [jsonrange](docs/jsonrange/) - a tool for iterating over JSON objects and arrays (return keys or values)
+ [toml2json](docs/toml2json/) - a tool for converting TOML to JSON
+ [xlsx2csv](docs/xlsx2csv/) - a tool for converting Excel Workbooks sheets to CSV files
+ [xlsx2json](docs/xlsx2json/) - a tool for converting Excel Workbooks to JSON files
+ [yaml2json](docs/yaml2json/) - a tool for converting YAML files to JSON


Compiled versions are provided for Linux (amd64), Mac OS X (amd64),
Windows 10 (amd64) and Raspbian (ARM7). See https://github.com/caltechlibrary/datatools/releases.

Use "-help" option for a full list of options for each utility (e.g. `csv2json -help`).

## For strings

_datatools_ provides the [string](docs/string/) command for working with 
text strings (limited to memory available).  This is commonly needed when 
cleanup data for analysis. The _string_ command was created for when the 
old Unix standbys- grep, awk, sed, tr are unwieldly or inconvient. 
_string_ provides operations are common in most language like, trimming, 
spliting, and transforming letter case.  The _string_ command also makes 
it easy to join JSON string arrays into single a string using a delimiter 
or split a string into a JSON array based on a delimiter. The form of the 
command is `string [OPTIONS] [ACTION] [ARCTION_PARAMETERS...]`

```shell
    string toupper "one two three"
```

Would yield "ONE TWO THREE".

Some of the features included

+ change case (upper, lower, title, English title)
+ length, position and count of substrings
+ has prefix, suffix or contains
+ trim prefix, suffix and cutsets
+ split and join to/from JSON string arrays

See [string](docs/string/) for full details

## For scripting

Various utilities for simplifying work on the command line. 

+ [findfile](docs/findfile/) - find files based on prefix, suffix or contained string
+ [finddir](docs/finddir/) - find directories based on prefix, suffix or contained string
+ [mergepath](docs/mergepath/) - prefix, append, clip path variables
+ [range](docs/range/) - emit a range of integers (useful for numbered loops in Bash)
+ [reldate](docs/reldate/) - display a relative date in YYYY-MM-DD format
+ [timefmt](docs/timefmt/) - format a time value based on Golang's time format language
+ [urlparse](docs/urlparse/) - split a URL into parts

Compiled versions are provided for Linux (amd64), Mac OS X (amd64),
Windows 10 (amd64) and Raspbian (ARM7). See https://github.com/caltechlibrary/datatools/releases.

Use the utilities try "-help" option for a full list of options.


## Installation

See [INSTALL.md](install.html) for details for installing pre-compiled 
versions of the programs.

