
# Using csvjoin

Simple usage of building a merged CSV file from data1.csv
and data2.csv where column 1 in data1.csv matches the value in
column 3 of data2.csv with the results being written to 
merged-data.csv.

```shell
    csvjoin -csv1=data1.csv -col1=2 \
       -csv2=data2.csv -col2=4 \
       -output=merged-data.csv
```

