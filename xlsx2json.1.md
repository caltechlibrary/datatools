---
title: "{app_name} (1) user manual"
author: "R. S. Doiel"
pubDate: 2023-01-09
---

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

# DESCRIPTION

{app_name} is a tool that converts individual Excel Workbook Sheets into
JSON output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-N, -sheets
: display sheet names in Excel Workbook

-c, -count
: display number of sheets in Excel Workbook

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

This would get the sheet named "Sheet 1" from "MyWorkbook.xlsx" and save as sheet1.json

~~~
    {app_name} MyWorkbook.xlsx "My worksheet 1" > sheet1.json
~~~

This would get the number of sheets in the workbook

~~~
    {app_name} -count MyWorkbook.xlsx
~~~

This will output the title of the sheets in the workbook

~~~
    {app_name} -sheets MyWorkbook.xlsx
~~~

Putting it all together in a shell script and convert all sheets to
into JSON documents..

~~~
	{app_name} -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	JSON_NAME="${SHEET_NAME// /-}.json"
    	{app_name} -o "${JSON_NAME}" MyWorkbook.xlsx "$SHEET_NAME"
	done
~~~

{app_name} {version}

