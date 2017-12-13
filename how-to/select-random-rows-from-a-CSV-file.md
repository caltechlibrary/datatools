
# How to select random rows from a CSV file


Given the a CSV file called _data.csv_ you select a sampling 
of rows with the command *csvrows* using a few options. In
this example we will assume _data.csv_ has a
header row we want to preserve and that our resulting sample
will be called _sample.csv_.  The options we use are

+ `-i` selecting _data.csv_ as the input source
+ `-o` sends the resulting CSV to the file named _sample.csv_
+ `-header=true` indicates the header should be preserved and not be counted as part of the sample
+ `-random` sets the number or rows to return in the sample, in this case twenty

Putting it all together--

```shell
    csvrows -i data.csv -o sample.csv -header=true -random=20
```

NOTE: If _data.csv_ has less than 20 rows then _sample.csv_
will include all the rows of _data.csv_ in a shuffled
order.

## How -random works

_csvrows_ reads in the entire csv file into memory, shuffles
the row using Go's rand package to calculate the rows to swap
and then write out the number of rows request in the shuffled
order.  The randomness is limitted by the shuffle and the 
write of the first N shuffled rows.

