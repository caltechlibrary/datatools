
# Using csv2xlsx

Converting a csv to a workbook.

```shell
	csv2xlsx -i data1.csv MyWorkbook.xlsx 'My worksheet 1'
```

This creates a new 'My worksheet 1' in the Excel Workbook
called 'MyWorkbook.xlsx' with the contents of data.csv.

```shell
	cat data1.csv | csv2xlsx MyWorkbook.xlsx 'My worksheet 2'
```

This does the same but the contents of data.csv are piped into
the workbook's 'My worksheet 2' sheet.

## example files

- [data1.csv](data1.csv)
- [MyWorkbook.xlsx](MyWorkbook.xlsx)
- [csv2xlsx-demo.bash](csv2xlsx-demo.bash)

