
# Using csv2xlsx

Converting a csv to a workbook.

```shell
	csv2xlsx -i data.csv MyWorkbook.xlsx 'My worksheet 1'
```

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

```shell
	cat data.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet 2'
```

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.

