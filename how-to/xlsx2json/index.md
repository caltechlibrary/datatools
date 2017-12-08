
# Using xlsx2json

This would get the sheet named "Sheet 1" from "my-workbook.xlsx" and save as sheet1.json

```shell
    xlsx2json my-workbook.xlsx "Sheet 1" > sheet1.json
```

This would get the number of sheets in the workbook

```shell
    xlsx2json -count my-workbook.xlsx
```

This will output the title of the sheets in the workbook

```shell
    xlsx2json -sheets my-workbook.xlsx
```

Putting it all together in a shell script and convert all sheets to
into JSON documents..

```shell
	 xlsx2json -n my-workbook.xlsx | while read SHEET_NAME; do
       xlsx2json my-workbook.xlsx "$SHEET_NAME" > \
	       "${SHEET_NAME// /-}.json"
    done
```

