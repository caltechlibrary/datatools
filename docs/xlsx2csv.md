
# USAGE

## xlsx2csv [OPTIONS] EXCEL_WORKBOOK_NAME [SHEET_NAME]

## SYNOPSIS

xlsx2csv is a tool that converts individual Excel Sheets to CSV output.

## OPTIONS

	-c	display number of sheets in Excel Workbook
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-n	display sheet names in Excel W9rkbook
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version

## EXAMPLE

```shell
    xlsx2csv my-workbook.xlsx "Sheet 1" > sheet1.csv
```

This would get the first sheet from the workbook and save it as a CSV file.

```shell
    xlsx2csv -c my-workbook.xlsx
```

This will output the number of sheets in the Workbook.

```shell
    xlsx2csv -n my-workbook.xlsx
```

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

```shell
    for SHEET_NAME in $(xlsx2csv -n my-workbook.xlsx); do
       xlsx2csv my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.csv"
    done
```


xlsx2csv v0.0.16
