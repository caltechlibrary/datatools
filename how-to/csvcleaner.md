
# Using csvcleaner

Normalizing the a spread sheet's column count to 2 per row.

```shell
    cat mysheet.csv | csvcleaner -field-per-row=2 > 2cols.csv
```

Normalizing a spread sheet's column count to 3 (add a padding column as needed per row).

```shell
    cat mysheet.csv | csvcleaner -field-per-row=3 > 3cols.csv
```

Trim leading spaces.

```shell
    cat mysheet.csv | csvcleaner -left-trim-spaces > ltrim.csv
```

Trim trailing spaces.

```shell
    cat mysheet.csv | csvcleaner -right-trim-spaces > rtrim.csv
```

Trim leading and trailing spaces

```shell
    cat mysheet.csv | csvcleaner -trim-spaces > trim.csv
```

## example files

- [mysheet.csv](mysheet.csv)
- [2cols.csv](2cols.csv)
- [3cols.csv](3cols.csv)
- [ltrim.csv](ltrim.csv)
- [rtrim.csv](rtrim.csv)
- [trim.csv](trim.csv)
- [csvcleaner-demo.bash](csvcleaner-demo.bash)


