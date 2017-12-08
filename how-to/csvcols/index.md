
# Using csvcols

Simple usage of building a CSV file one row at a time.

```shell
    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv
```

Example parsing a pipe delimited string into a CSV line.

```shell
    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
```

Using a pipe filter a 3 column CSV for columns 1 and 3 into 2col.csv.

```shell
    cat 3col.csv | csvcols -col 1,3 > 2col.csv
```

Using options filter a 3 column CSV file for columns 1,3 into 2col.csv.

```shell
    csvcols -i 3col.csv -col 1,3 -o 2col.csv
```

