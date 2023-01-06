
USAGE: xlsx2json [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

DESCRIPTION

xlsx2json is a tool that converts individual Excel Workbook Sheets into
JSON output.

OPTIONS

    -N, -sheets          display sheet names in Excel Workbook
    -c, -count           display number of sheets in Excel Workbook
    -examples            display example(s)
    -generate-manpage    generate man page
    -generate-markdown   generate markdown documentation
    -h, -help            display help
    -l, -license         display license
    -nl, -newline        if true add a trailing newline
    -o, -output          output filename
    -quiet               suppress error messages
    -v, -version         display version


EXAMPLES

This would get the sheet named "Sheet 1" from "MyWorkbook.xlsx" and save as sheet1.json

    xlsx2json MyWorkbook.xlsx "My worksheet 1" > sheet1.json

This would get the number of sheets in the workbook

    xlsx2json -count MyWorkbook.xlsx

This will output the title of the sheets in the workbook

    xlsx2json -sheets MyWorkbook.xlsx

Putting it all together in a shell script and convert all sheets to
into JSON documents..

	xlsx2json -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	JSON_NAME="${SHEET_NAME// /-}.json"
    	xlsx2json -o "${JSON_NAME}" MyWorkbook.xlsx "$SHEET_NAME"
	done

xlsx2json 1.2.1
