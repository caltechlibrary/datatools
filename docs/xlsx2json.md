
# USAGE

## xlsx2json [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

## SYNOPSIS

xlsx2json is a tool that converts individual Excel Workbook Sheets into
JSON output.

## OPTIONS

	-c	display number of sheets in Excel Workbook
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-n	display sheet names in Excel Workbook
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version

## EXAMPLE

```shell
    xlsx2json my-workbook.xlsx "Sheet 1" > sheet1.json
```

This would get the first sheet from the workbook and save it as a JSON file.

```shell
    xlsx2json -c my-workbook.xlsx
```

This will output the number of sheels in the Workbook.

```shell
    xlsx2json -n my-workbook.xlsx
```

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

```shell
    for SHEET_NAME in $(xlsx2json -n my-workbook.xlsx); do
       xlsx2json my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.json"
    done
```


xlsx2json v0.0.14
