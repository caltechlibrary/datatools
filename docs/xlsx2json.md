
# USAGE

## xlsx2json [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

## SYNOPSIS

xlsx2json is a tool that converts individual Excel Workbook Sheets into
JSON output.

## OPTIONS	

```
    -c	display number of sheets in Excel Workbook
	-example	display example(s)
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-n	display sheet names in Excel Workbook
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version
```

## EXAMPLE

This would get the sheet "Sheet 1" from "my-workbook.xlsx" and save as sheet1.json

```shell
    xlsx2json my-workbook.xlsx "Sheet 1" > sheet1.json
```

This would get the number of sheets in a workbook

```shell
    xlsx2json -c my-workbook.xlsx
```

This will output the title of the sheets in the Workbook.

```shell
    xlsx2json -n my-workbook.xlsx
```

Putting it all together in a shell script and converting
all sheets to JSON documents.

```shell
    xlsx2json -n my-workbook.xlsx | while read SHEET_NAME; do
       xlsx2json my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.json"
    done
```

xlsx2json v0.0.18
