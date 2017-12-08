
# Using csvrows

Simple usage of building a CSV file one rows at a time.

```shell
    csvrows -o 4rows.csv "First,Second,Third" "one,two,three"
    csvrows "ein,zwei,drei" "1,2,3" >> 4rows.csv
    cat 4row.csv
```

Example parsing a pipe delimited string into a CSV line

```shell
    csvrows -d "|" "First,Second,Third|one,two,three" > 4rows.csv
    csvrows -delimiter "|" "ein,zwei,drei|1,2,3" >> 4rows.csv
    cat 4rows.csv
```

Filter rows 1 and 3 from CSV file

```shell
    cat 4row.csv | csvrows -row 1,3 > result1.csv
```

Filter rows 1 and 3 from CSV file from file named "4row.csv"

```shell
    csvrows -i 4row.csv -row 1,3 > result2.csv
```

