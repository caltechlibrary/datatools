
# csv2xlsx

## USAGE

    csv2xlsx [OPTIONS] WORKBOOK_NAME SHEET_NAME

## SYNOPSIS

csv2xlsx will take CSV input and create a new sheet in an Excel Workbook.
If the Workbook does not exist then it is created. 

## OPTIONS

```
	-h	display help
	-help	display help
	-i	input filename (CSV content)
	-input	input filename (CSV content)
	-l	display license
	-license	display license
	-sheet	Sheet name to create/replace
	-v	display version
	-version	display version
	-workbook	Workbook name
```

## EXAMPLE

```
	csv2xlsx -i data.csv MyWorkbook.xlsx 'My worksheet'
```

This creates a new 'My worksheet' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

```
	cat data.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet'
```

This does the same but the contents of data.csv are piped into
the workbook's sheet.

