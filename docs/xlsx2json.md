
# xlsx2json

## USAGE

    xlsx2json [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

## SYNOPSIS

xlsx2json is a tool that converts individual Excel Workbook Sheets into
JSON output.

## OPTIONS

```
	-c	display number of sheets in Excel Workbook
	-h	display help
	-l	display license
	-n	display sheet names in Excel Workbook
	-v	display version
```

## EXAMPLE

```
    xlsx2json my-workbook.xlsx "Sheet 1" > sheet1.json
```

This would get the first sheet from the workbook and save it as a JSON file.

```
    xlsx2json -c my-workbook.xlsx
```

This will output the number of sheels in the Workbook.

```
    xlsx2json -n my-workbook.xlsx
```

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

```
    for SHEET_NAME in $(xlsx2json -n my-workbook.xlsx); do
       xlsx2json my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.json"
    done
```

