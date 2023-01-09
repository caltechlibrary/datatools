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

{app_name} is a tool that converts individual Excel Sheets to CSV output.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-N, -sheets
: display the Workbook sheet names

-c, -count
: display number of Workbook sheets

-nl, -newline
: if true add a trailing newline

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

Extract a workbook sheet as a CSV file

~~~
    {app_name} MyWorkbook.xlsx "My worksheet 1" > sheet1.csv
~~~

This would get the first sheet from the workbook and save it as a CSV file.

~~~
    {app_name} -count MyWorkbook.xlsx
~~~


This will output the number of sheets in the Workbook.

~~~
    {app_name} -sheets MyWorkbook.xlsx
~~~

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

~~~
	{app_name} -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	CSV_NAME="${SHEET_NAME// /-}.csv"
    	{app_name} -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
	done
~~~

{app_name} {version}

