---
title: "Datatools User Manual"
author: "R. S. Doiel"
pubDate: 2023-01-12
---

# Datatools User Manual

Below is a collection man manual pages for each of the command line tools
that comes with the datatools project.

## Data transformation and conversion

- [codemeta2cff](codemeta2cff.1.html), generate a CITATION.cff file from a codemeta.json file
- [csv2json](csv2json.1.html), convert CSV into a JSON
- [csv2jsonl](csv2jsonl.1.html), convert CSV into a [JSON lines](https://jsonlines.org) stream.
- [csv2mdtable](csv2mdtable.1.html), convert CSV into a Markdown table (for use with Pandoc)
- [csv2tab](csv2tab.1.html), convert CSV to a tab delimited file
- [csv2xlsx](csv2xlsx.1.html), convert CSV to Excel XML formatted file
- [csvcleaner](csvcleaner.1.html), cleanup a CSV file and normalize it
- [csvcols](csvcols.1.html), extract columns of values from a CSV file
- [csvfind](csvfind.1.html), find content in a CSV file
- [csvjoin](csvjoin.1.html), join two CSV files into one
- [csvrows](csvrows.1.html), extract rows of values from a CSV file
- [finddir](finddir.1.html), find a directory 
- [findfile](findfile.1.html), find a file (e.g. list for a files recursively by file extension)
- [json2toml](json2toml.1.html), convert JSON to TOML
- [json2yaml](json2yaml.1.html), convert JSON to YAML
- [jsoncols](jsoncols.1.html), extract columns from JSON
- [jsonjoin](jsonjoin.1.html), join JSON documents
- [jsonmunge](jsonmunge.1.html), process JSON through a go template
- [jsonrange](jsonrange.1.html), iterate a JSON expression of a list
- [jsonobjects2csv](jsonobjects2csv.1.html), render a JSON list of objects to CSV file, flattens cells as YAML if needed
- [json2jsonl](json2jsonl.1.html), render a JSON array document as JSON lines
- [sql2csv](sql2csv.1.html), convert a SQL query into a CSV output
- [tab2csv](tab2csv.1.html), tab delimited file to CSV
- [toml2json](toml2json.1.html), TAML to JSON
- [xlsx2csv](xlsx2csv.1.html), convert an Excel XML file's "sheet" to csv
- [xlsx2json](xlsx2json.1.html), convert an Excel XML file's "sheet" into JSON
- [yaml2json](yaml2json.1.html), convert YAML into JSON

## Shell helpers

- [mergepath](mergepath.1.html), manage the PATH environment variable (e.g. remove duplication paths, append, insert and cut from PATH list)
- [range](range.1.html), emit a range of numbers (can be ascending, descending, odd/even, etc)
- [reldate](reldate.1.html), compute a relative date in YYYY-MM-DD format
- [reltime](reltime.1.html), compute a relative time in HH:MM:SS format
- [timefmt](timefmt.1.html), format a time string
- [urlparse](urlparse.1.html), parse a URL into its components (e.g. protocol, hostname, path)


## String manipution

- [string](string.1.html), various string manipulation actions, append, cut, pad, trim, etc.


