
# Using csvjoin

Simple usage of building a merged CSV file from data1-2.csv
and data2-2.csv where column 1 in data1.csv matches the value in
column 3 of data2-2.csv with the results being written to 
merged-data-2.csv.

```shell
    csvjoin -csv1=data1-2.csv -col1=2 \
       -csv2=data2-2.csv -col2=4 \
       -output=merged-data-2.csv
```

## example files

- [data1-2.csv](data1-2.csv)
- [data2-2.csv](data2-2.csv)
- [merged-data-2.csv](merged-data-2.csv)
- [csvjoin-demo.bash](csvjoin-demo.bash)

