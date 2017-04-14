
# How to find duplicates in a column

Searching for duplicate values in a column can be done using _cat_, _csvcols_, _sort_ 
and _csvfind_. Here's the basic algorithm from the command line or Bash script.

+ for each line of your CSV file
    + extract the value in the colum
    + sort for unique values
    + for each unique value use _csvfind_ to output matching rows

Here's an example Bash script looking for duplicates in *dups.csv*
in column 2, second column (columns are counted from 1 rather than zero)

```shell
    CSV_FILE="dups.csv"
    CSV_COL_NO="2"

    csvcols -i "$CSV_FILE" -col "$CSV_COL_NO" | sort -u | while read CELL; do
	    if [ "$CELL" != "" ]; then
		    csvfind -i "$CSV_FILE" -trim-spaces -col "$CSV_COL_NO"  "${CELL}"
	    fi
    done
```

This would result a new CSV file with duplicates grouped together.
