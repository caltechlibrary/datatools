
USAGE: xlsx2csv [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

DESCRIPTION

xlsx2csv is a tool that converts individual Excel Sheets to CSV output.

OPTIONS

    -N, -sheets         display the Workbook sheet names
    -c, -count          display number of Workbook sheets
    -examples           display example(s)
    -generate-manpage   generate man page
    -generate-markdown  generate markdown documentation
    -h, -help           display help
    -l, -license        display license
    -nl, -newline       if true add a trailing newline
    -o, -output         output filename
    -quiet              suppress error messages
    -v, -version        display version


EXAMPLES

Extract a workbook sheet as a CSV file

    xlsx2csv MyWorkbook.xlsx "My worksheet 1" > sheet1.csv

This would get the first sheet from the workbook and save it as a CSV file.

    xlsx2csv -count MyWorkbook.xlsx

This will output the number of sheets in the Workbook.

    xlsx2csv -sheets MyWorkbook.xlsx

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

	xlsx2csv -N MyWorkbook.xlsx | while read SHEET_NAME; do
    	CSV_NAME="${SHEET_NAME// /-}.csv"
    	xlsx2csv -o "${CSV_NAME}" MyWorkbook.xlsx "${SHEET_NAME}" 
	done

xlsx2csv 1.2.1
