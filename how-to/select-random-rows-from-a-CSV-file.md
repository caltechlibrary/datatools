
# How to select random rows from a spreadsheet


Given the a CSV file called _data.csv_ you select a sampling 
of rows with the command *csvrows* using a few option. For
the purposes this example we will assume _data.csv_ has a
header row we want to preserve and that our resulting sample
will be called _sample.csv_. 

```shell
    csvrows -i data.csv -header=true -random=20
```

NOTE: If _data.csv_ has less than 20 rows then the result
will include all the rows of _data.csv_ in a shuffled
order.

## How -random works

_csvrows_ reads in the entire csv file into memory, shuffles
the row using Go's rand package to calculate the rows to swap
and then write out the number of rows request in the shuffled
order.  The randomness is limitted by the shuffle and the 
resulting writing of the first N shuffled rows.

