%xlsx2csv(1) user manual | version 1.3.5 10335ae
% R. S. Doiel
% 2026-02-12

# NAME

xlsx2csv

# SYNOPSIS

xlsx2csv [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

# DESCRIPTION

xlsx2csv is a tool that converts individual Excel Sheets to CSV output.

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
: add a trailing newline to the end of file (EOF)

-crlf
: use CRLF for end of line (EOL). On Windows this option is the default.
Set to false to use a LF on Windows. On other OS LF is the default.

-o, -output
: output filename

-quiet
: suppress error messages


# EXAMPLES

Extract a workbook sheet as a CSV file

~~~
    xlsx2csv MyWorkbook.xlsx "My worksheet 1" > sheet1.csv
~~~

This would get the first sheet from the workbook and save it as a CSV file.

~~~
    xlsx2csv -count MyWorkbook.xlsx
~~~


This will output the number of sheets in the Workbook.

~~~
    xlsx2csv -sheets MyWorkbook.xlsx
~~~

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

~~~
	xlsx2csv -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	CSV_NAME="${SHEET_NAME// /-}.csv"
    	xlsx2csv -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
	done
~~~

xlsx2csv 1.3.5


