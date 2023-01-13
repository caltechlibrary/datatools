
# Using csvcols

Simple usage of building a CSV file one row at a time.

```shell
    csvcols one two three > 3col-2.csv
    csvcols 1 2 3 >> 3col-2.csv
    cat 3col-2.csv
```

Example parsing a pipe delimited string into a CSV line.

```shell
    csvcols -d "|" "one|two|three" > 3col-2.csv
    csvcols -delimiter "|" "1|2|3" >> 3col-2.csv
    cat 3col-2.csv
```

Using a pipe filter a 3 column CSV for columns 1 and 3 into 2col-2.csv.

```shell
    cat 3col-2.csv | csvcols -col 1,3 > 2col-2.csv
```

Using options filter a 3 column CSV file for columns 1,3 into 2col-2.csv.

```shell
    csvcols -i 3col-2.csv -col 1,3 -o 2col-2.csv
```

## Quoting issues

Sometimes CSV files have sloppy or problematic quoting. Two options
are provided (-use-lazy-quotes, -trim-leading-space). In the example
below the CSV has trailing white space in the second row. By using
`-use-lazy-quotes` option we can get the first two columns without
running into CSV parsing errors.

```shell
    csvcols -i quoting-example.csv -use-lazy-quotes -col 1,2
```

Input (with trailing whitespace in row 2)

```csv
    "A","B","C@caltech.edu","2017-03-27 14:38:57","Yes","CCE","refund of my deposit"
    "C","D","E@caltech.edu","2017-04-05 10:50:42","Yes","EAS","receive a refund of my deposit"  
```

Output

```
    A,B
    C,D
```

## example files

- [3col-2.csv](3col-2.csv)
- [2col-2.csv](2col-2.csv)
- [quoting-example.csv](quoting-example.csv)
- [quoting-expected.csv](quoting-expected.csv)
- [csvcols-demo.bash](csvcols-demo.bash)
