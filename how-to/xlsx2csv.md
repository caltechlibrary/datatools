
# Using xlsx2csv

Extract a workbook sheet as a CSV file

```shell
    xlsx2csv my-workbook.xlsx "Sheet 1" > sheet1.csv
```

This would get the first sheet from the workbook and save it as a CSV file.

```shell
    xlsx2csv -count my-workbook.xlsx
```

This will output the number of sheets in the Workbook.

```shell
    xlsx2csv -sheets my-workbook.xlsx
```

This will display a list of sheet names, one per line.
Putting it all together in a shell script.

```shell
    for SHEET_NAME in $(xlsx2csv -n my-workbook.xlsx); do
       xlsx2csv my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.csv"
    done
```

## example files

- [sheet1.csv](sheet1.csv)
- [xlsx2csv-demo.bash](xlsx2csv-demo.bash)

